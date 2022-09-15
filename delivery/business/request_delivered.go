package business

import (
	"context"
	"errors"
	"fmt"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/delivery/db"
	"gitlab.artin.ai/backend/courier-management/delivery/model"
	"gitlab.artin.ai/backend/courier-management/delivery/status"
	deliveryPb "gitlab.artin.ai/backend/courier-management/grpc/delivery/go"
	uaa "gitlab.artin.ai/backend/courier-management/uaa/proto"
	"gorm.io/gorm"
)

var (
	DeliveredInvalidOrderErr = errors.New("delivered destination order is invalid")
)

func RequestDelivered(ctx context.Context, requestId string, destinationOrder int, description string, name string, signature string) (err error) {

	logger.Infof("RequestDelivered requestId = %v, destinationOrder = %v, description = %v, name = %v, signature = %v", requestId, destinationOrder, description, name, signature)

	tx := db.MariaDbClient().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	var requestModel model.Request
	result := tx.Where("ID = ?", requestId).
		Preload("Locations").
		First(&requestModel)
	if result.Error != nil {
		logger.Error("failed to check current request status", result.Error)
		return uaa.Internal.Error(result.Error)
	}

	// check destination order validity
	var lastStatusLocation model.RequestLocation
	err = tx.Where("ride_statuses.request_id = ? AND is_origin = 0",
		requestModel.ID).
		Joins("join ride_statuses on ride_statuses.ride_location_id = request_locations.id").
		Order("ride_statuses.created_at DESC").
		First(&lastStatusLocation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("there is no previous destination location in ride status history")
		logger.Error("the request has not reached to any destination yet", err)
		return uaa.NotFound.Error(err)
	} else if err != nil {
		logger.Error("failed to check last status details", err)
		return uaa.Internal.Error(err)
	}
	if lastStatusLocation.Order != destinationOrder {
		logger.Errorf("destination order is invalid. expected: %v, actual: %v",
			DeliveredInvalidOrderErr, lastStatusLocation.Order, destinationOrder)
		return uaa.InvalidDestinationOrder.Error(DeliveredInvalidOrderErr)
	}

	isValidTransition := status.RequestTransition(requestModel.Status, deliveryPb.RequestStatus_DELIVERED.String())
	if !isValidTransition {
		err = fmt.Errorf("transition to status delivered for request (id: %v) is not possible. current status: %v", requestId, requestModel.Status)
		logger.Error("failed to process request delivered", err)
		return uaa.DeliveryInvalidStateTransition.Error(err)
	}

	requestModel.Status = deliveryPb.RequestStatus_DELIVERED.String()
	requestModel.LastProcessedDestination = destinationOrder
	err = tx.Omit("Locations").
		Save(&requestModel).Error
	if err != nil {
		logger.Error("failed to update request status to delivered", err)
		return uaa.Internal.Error(err)
	}

	var destinationLocation model.RequestLocation
	for _, location := range requestModel.Locations {
		if location.Order == destinationOrder && !location.IsOrigin {
			destinationLocation = location
			break
		}
	}
	rideStatus := model.RideStatus{
		RideLocationId:   destinationLocation.ID,
		RequestId:        requestModel.ID,
		Status:           requestModel.Status,
		CancellationNote: "",
	}
	err = tx.Create(&rideStatus).Error
	if err != nil {
		logger.Error("failed to create ride status", err)
		return uaa.Internal.Error(err)
	}

	rideConfirmation := model.RideConfirmation{
		RideLocationId: destinationLocation.ID,
		RequestId:      requestModel.ID,
		Name:           name,
		Signature:      signature,
	}
	err = tx.Create(&rideConfirmation).Error
	if err != nil {
		logger.Error("failed to create ride confirmation", err)
		return uaa.Internal.Error(err)
	}

	err = publishRequestDelivered(ctx, deliveryPb.RequestDeliveredEvent{
		RequestId:        requestModel.ID,
		DestinationOrder: int32(destinationOrder),
		Desc:             description,
	}, requestModel.CustomerId, destinationLocation.ToProto(), requestModel.HumanReadableId)
	if err != nil {
		logger.Error("failed to publish request delivered event", err)
		return uaa.Internal.Error(err)
	}

	return nil
}
