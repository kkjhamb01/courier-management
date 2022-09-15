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
	ArrivedInvalidOrderErr         = errors.New("arrived order is invalid")
	ArrivedNoFurtherDestinationErr = errors.New("there is no further destination to arrive to")
)

func RequestArrivedOrigin(ctx context.Context, requestId string, description string) (err error) {

	logger.Infof("RequestArrivedOrigin requestId = %v, description = %v", requestId, description)

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

	isValidTransition := status.RequestTransition(requestModel.Status, deliveryPb.RequestStatus_ARRIVED.String())
	if !isValidTransition {
		err := fmt.Errorf("transition to status arrived for request (id: %v) is not possible. current status: %v", requestId, requestModel.Status)
		logger.Error("failed to process request arrived", err)
		return uaa.DeliveryInvalidStateTransition.Error(err)
	}

	requestModel.Status = deliveryPb.RequestStatus_ARRIVED.String()
	err = tx.Omit("Locations").
		Save(&requestModel).Error
	if err != nil {
		logger.Error("failed to update request status to arrived", err)
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

	err = publishRequestArrivedOrigin(ctx, deliveryPb.RequestArrivedOriginEvent{
		RequestId: requestModel.ID,
		Desc:      description,
	}, requestModel.CustomerId, originLocation.ToProto(), requestModel.HumanReadableId)
	if err != nil {
		logger.Error("failed to publish request arrived event", err)
		return uaa.Internal.Error(err)
	}

	return nil
}

func RequestArrivedDestination(ctx context.Context, requestId string, destinationOrder int, description string) (err error) {

	logger.Infof("RequestArrivedDestination requestId = %v, destinationOrder = %v, description = %v", requestId, destinationOrder, description)

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
	err = tx.Where("ride_statuses.request_id = ? AND ride_statuses.status = ? AND is_origin = false",
		requestModel.ID, deliveryPb.RequestStatus_NAVIGATING.String()).
		Joins("join ride_statuses on ride_statuses.ride_location_id = request_locations.id").
		Order("ride_statuses.created_at DESC").
		First(&lastNavigatingLocation).Error
	if err != nil {
		logger.Error("failed to check last navigating details", err)
		return uaa.Internal.Error(err)
	}
	if lastNavigatingLocation.Order != destinationOrder {
		logger.Errorf("destination order is invalid. expected: %v, actual: %v",
			ArrivedInvalidOrderErr, lastNavigatingLocation.Order, destinationOrder)
		return uaa.InvalidDestinationOrder.Error(ArrivedInvalidOrderErr)
	}
	if destinationOrder > len(requestModel.Locations)-1 {
		logger.Error("destination order is invalid", ArrivedNoFurtherDestinationErr)
		return uaa.InvalidDestinationOrder.Error(ArrivedNoFurtherDestinationErr)
	}

	isValidTransition := status.RequestTransition(requestModel.Status, deliveryPb.RequestStatus_ARRIVED.String())
	if !isValidTransition {
		err := fmt.Errorf("transition to status arrived for request (id: %v) is not possible. current status: %v", requestId, requestModel.Status)
		logger.Error("failed to process request arrived", err)
		return uaa.DeliveryInvalidStateTransition.Error(err)
	}

	requestModel.Status = deliveryPb.RequestStatus_ARRIVED.String()
	requestModel.LastProcessedDestination = destinationOrder
	err = tx.Omit("Locations").
		Save(&requestModel).Error
	if err != nil {
		logger.Error("failed to update request status to arrived", err)
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

	var originLocation model.RequestLocation
	for _, location := range requestModel.Locations {
		if location.IsOrigin {
			originLocation = location
		}
	}

	err = publishRequestArrivedDestination(ctx, deliveryPb.RequestArrivedDestinationEvent{
		RequestId:        requestModel.ID,
		DestinationOrder: int32(destinationOrder),
		Desc:             description,
	}, requestModel.CustomerId, originLocation.PhoneNumber, destinationLocation.ToProto(), requestModel.HumanReadableId)
	if err != nil {
		logger.Error("failed to publish request arrived event", err)
		return uaa.Internal.Error(err)
	}

	return nil
}
