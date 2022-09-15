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
)

var (
	NavigatingInvalidOrderErr         = errors.New("navigating destination order is invalid")
	NavigatingNoFurtherDestinationErr = errors.New("there is no further destination to navigate to")
)

func RequestNavigatingToReceiver(ctx context.Context, requestId string, destinationOrder int) (err error) {

	logger.Infof("RequestNavigatingToReceiver requestId = %v, destinationOrder = %v", requestId, destinationOrder)

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
	var lastNavigatingLocation model.RequestLocation
	err = tx.Where("ride_statuses.request_id = ? AND ride_statuses.status = ?",
		requestModel.ID, deliveryPb.RequestStatus_NAVIGATING.String()).
		Joins("join ride_statuses on ride_statuses.ride_location_id = request_locations.id").
		Order("ride_statuses.created_at DESC").
		First(&lastNavigatingLocation).Error
	if err != nil {
		logger.Error("failed to check last navigation details", err)
		return uaa.Internal.Error(err)
	}
	if lastNavigatingLocation.Order != destinationOrder-1 {
		logger.Errorf("destination order is invalid. expected: %v, actual: %v",
			NavigatingInvalidOrderErr, lastNavigatingLocation.Order+1, destinationOrder)
		return uaa.InvalidDestinationOrder.Error(NavigatingInvalidOrderErr)
	}
	if destinationOrder > len(requestModel.Locations)-1 {
		logger.Error("destination order is invalid", NavigatingNoFurtherDestinationErr)
		return uaa.InvalidDestinationOrder.Error(NavigatingNoFurtherDestinationErr)
	}

	isValidTransition := status.RequestTransition(requestModel.Status, deliveryPb.RequestStatus_NAVIGATING.String())
	if !isValidTransition {
		err = fmt.Errorf("transition to status navigating for request (id: %v) is not possible. current status: %v", requestId, requestModel.Status)
		logger.Error("failed to process request navigating", err)
		return uaa.DeliveryInvalidStateTransition.Error(err)
	}

	requestModel.Status = deliveryPb.RequestStatus_NAVIGATING.String()
	requestModel.LastProcessedDestination = destinationOrder
	err = tx.Omit("Locations").
		Save(&requestModel).Error
	if err != nil {
		logger.Error("failed to update request status to navigating", err)
		return uaa.Internal.Error(err)
	}

	var destinationId string
	for _, location := range requestModel.Locations {
		if location.Order == destinationOrder && !location.IsOrigin {
			destinationId = location.ID
			break
		}
	}
	rideStatus := model.RideStatus{
		RideLocationId:   destinationId,
		RequestId:        requestModel.ID,
		Status:           requestModel.Status,
		CancellationNote: "",
	}
	err = tx.Create(&rideStatus).Error
	if err != nil {
		logger.Error("failed to create ride status", err)
		return uaa.Internal.Error(err)
	}

	err = publishRequestNavigatingToReceiver(deliveryPb.RequestNavigatingToReceiverEvent{
		RequestId:              requestModel.ID,
		TargetDestinationOrder: int32(destinationOrder),
	})
	if err != nil {
		logger.Error("failed to publish request navigating event", err)
		return uaa.Internal.Error(err)
	}

	return nil
}

func RequestNavigatingToSender(ctx context.Context, requestId string) (err error) {
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

	// state transitioning rules of NAVIGATING is different for navigating to sender and navigating to receiver
	// it is recommended to include it into status module of delivery
	// I will the state transition here for now:
	if requestModel.Status != deliveryPb.RequestStatus_ACCEPTED.String() {
		err = fmt.Errorf("transition to status navigating for request (id: %v) is not possible. current status: %v", requestId, requestModel.Status)
		logger.Error("failed to process request navigating", err)
		return uaa.DeliveryInvalidStateTransition.Error(err)
	}

	requestModel.Status = deliveryPb.RequestStatus_NAVIGATING.String()
	err = tx.Omit("Locations").
		Save(&requestModel).Error
	if err != nil {
		logger.Error("failed to update request status to navigating", err)
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

	err = publishRequestNavigatingToSender(deliveryPb.RequestNavigatingToSenderEvent{
		RequestId: requestModel.ID,
	})
	if err != nil {
		logger.Error("failed to publish request navigating to sender event", err)
		return uaa.Internal.Error(err)
	}

	return nil
}
