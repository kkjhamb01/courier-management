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
	"gorm.io/gorm"
)

var SenderNotAvailableHasPickedUpErr = errors.New("the package has been picked up")

func RequestSenderNotAvailable(ctx context.Context, requestId string, description string) (err error) {

	logger.Infof("RequestReceiverNotAvailable requestId = %v, description = %v", requestId, description)

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
	var destinationStatusLocations []model.RequestLocation
	err = tx.Where("ride_statuses.request_id = ? AND status = ? AND is_origin = 1",
		requestModel.ID, deliveryPb.RequestStatus_PICKED_UP.String()).
		Joins("join ride_statuses on ride_statuses.ride_location_id = request_locations.id").
		First(&destinationStatusLocations).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error("failed to check last status details", err)
		return uaa.NotFound.Error(err)
	}
	if len(destinationStatusLocations) > 0 {
		logger.Error("failed to change status to sender not available", SenderNotAvailableHasPickedUpErr)
		return uaa.Internal.Error(SenderNotAvailableHasPickedUpErr)
	}

	isValidTransition := status.RequestTransition(requestModel.Status, deliveryPb.RequestStatus_SENDER_NOT_AVAILABLE.String())
	if !isValidTransition {
		err := fmt.Errorf("transition to status RequestStatus_SENDER_NOT_AVAILABLE for request (id: %v) is not possible. current status: %v", requestId, requestModel.Status)
		logger.Error("failed to set RequestStatus_SENDER_NOT_AVAILABLE for request", err)
		return uaa.DeliveryInvalidStateTransition.Error(err)
	}

	requestModel.Status = deliveryPb.RequestStatus_SENDER_NOT_AVAILABLE.String()
	err = tx.Omit("Locations").
		Save(&requestModel).Error
	if err != nil {
		logger.Error("failed to update request status to sender not available", err)
		return uaa.Internal.Error(err)
	}

	var originLocation model.RequestLocation
	for _, location := range requestModel.Locations {
		if location.IsOrigin {
			originLocation = location
			break
		}
	}

	rideStatus := model.RideStatus{
		RideLocationId:   originLocation.ID,
		RequestId:        requestModel.ID,
		Status:           requestModel.Status,
		CancellationNote: "",
	}
	err = tx.Create(&rideStatus).Error
	if err != nil {
		logger.Error("failed to create ride status", err)
		return uaa.Internal.Error(err)
	}

	err = publishRequestSenderNotAvailable(ctx, deliveryPb.RequestSenderNotAvailableEvent{
		RequestId: requestModel.ID,
		Desc:      description,
	}, originLocation.ToProto(), requestModel.CustomerId, requestModel.HumanReadableId)
	if err != nil {
		logger.Error("failed to publish request sender unavailable event", err)
		return uaa.Internal.Error(err)
	}

	return nil
}
