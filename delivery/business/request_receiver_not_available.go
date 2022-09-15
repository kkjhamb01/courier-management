package business

import (
	"context"
	"errors"
	"fmt"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/delivery/db"
	"github.com/kkjhamb01/courier-management/delivery/model"
	"github.com/kkjhamb01/courier-management/delivery/status"
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
	uaa "github.com/kkjhamb01/courier-management/uaa/proto"
)

var (
	ReceiverNotAvailableInvalidOrderErr = errors.New("receiver not available destination order is invalid")
)

func RequestReceiverNotAvailable(ctx context.Context, requestId string, destinationOrder int, description string) (err error) {

	logger.Infof("RequestReceiverNotAvailable requestId = %v, destinationOrder = %v, description = %v", requestId, destinationOrder, description)

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
	if err != nil {
		logger.Error("failed to check last status details", err)
		return uaa.InvalidDestinationOrder.Error(err)
	}
	if lastStatusLocation.Order != destinationOrder {
		logger.Errorf("destination order is invalid. expected: %v, actual: %v",
			ReceiverNotAvailableInvalidOrderErr, lastStatusLocation.Order, destinationOrder)
		return uaa.InvalidDestinationOrder.Error(ReceiverNotAvailableInvalidOrderErr)
	}

	isValidTransition := status.RequestTransition(requestModel.Status, deliveryPb.RequestStatus_RECEIVER_NOT_AVAILABLE.String())
	if !isValidTransition {
		err := fmt.Errorf("transition to status RequestStatus_RECEIVER_NOT_AVAILABLE for request (id: %v) is not possible. current status: %v", requestId, requestModel.Status)
		logger.Error("failed to set RequestStatus_RECEIVER_NOT_AVAILABLE for request", err)
		return uaa.DeliveryInvalidStateTransition.Error(err)
	}

	requestModel.Status = deliveryPb.RequestStatus_RECEIVER_NOT_AVAILABLE.String()
	requestModel.LastProcessedDestination = destinationOrder
	err = tx.Omit("Locations").
		Save(&requestModel).Error
	if err != nil {
		logger.Error("failed to update request status to receiver not available", err)
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

	err = publishRequestReceiverNotAvailable(ctx, deliveryPb.RequestReceiverNotAvailableEvent{
		RequestId:        requestModel.ID,
		DestinationOrder: int32(destinationOrder),
		Desc:             description,
	}, requestModel.CustomerId, destinationLocation.ToProto(), requestModel.HumanReadableId)
	if err != nil {
		logger.Error("failed to publish request receiver unavailable event", err)
		return uaa.Internal.Error(err)
	}

	return nil
}
