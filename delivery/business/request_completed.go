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
	CompletedLastDestinationNotDeliveredErr = errors.New("the last location has not been delivered")
)

func RequestCompleted(ctx context.Context, requestId string, description string) (err error) {

	logger.Infof("RequestCompleted requestId = %v, description = %v", requestId, description)

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
	err = tx.Where("ride_statuses.request_id = ? AND ride_statuses.status = ? AND is_origin = false",
		requestModel.ID, deliveryPb.RequestStatus_DELIVERED.String()).
		Joins("join ride_statuses on ride_statuses.ride_location_id = request_locations.id").
		Order("ride_statuses.created_at DESC").
		First(&lastStatusLocation).Error
	if err != nil {
		logger.Error("failed to check last status details", err)
		return uaa.Internal.Error(err)
	}
	if lastStatusLocation.Order != len(requestModel.Locations)-1 {
		logger.Errorf("the last destination has not been delivered", CompletedLastDestinationNotDeliveredErr)
		return uaa.LastDestinationIsNotDelivered.Error(CompletedLastDestinationNotDeliveredErr)
	}

	isValidTransition := status.RequestTransition(requestModel.Status, deliveryPb.RequestStatus_COMPLETED.String())
	if !isValidTransition {
		err = fmt.Errorf("transition to status completed for request (id: %v) is not possible. current status: %v", requestId, requestModel.Status)
		logger.Error("failed to process request delivered", err)
		return uaa.DeliveryInvalidStateTransition.Error(err)
	}

	requestModel.Status = deliveryPb.RequestStatus_COMPLETED.String()
	err = tx.Omit("Locations").
		Save(&requestModel).Error
	if err != nil {
		logger.Error("failed to update request status to completed", err)
		return uaa.Internal.Error(err)
	}

	var destinationLocation model.RequestLocation
	lastOrder := len(requestModel.Locations) - 1
	for _, location := range requestModel.Locations {
		if location.Order == lastOrder && !location.IsOrigin {
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

	err = publishRequestCompleted(ctx, deliveryPb.RequestCompletedEvent{
		RequestId: requestModel.ID,
		Desc:      description,
	}, requestModel.CustomerId, destinationLocation.ToProto(), requestModel.HumanReadableId)
	if err != nil {
		logger.Error("failed to publish request completed event", err)
		return uaa.Internal.Error(err)
	}

	return nil
}
