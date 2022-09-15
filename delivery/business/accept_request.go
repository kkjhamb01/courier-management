package business

import (
	"context"
	"fmt"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/delivery/db"
	"github.com/kkjhamb01/courier-management/delivery/model"
	"github.com/kkjhamb01/courier-management/delivery/status"
	storage "github.com/kkjhamb01/courier-management/delivery/storage/lock"
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
	"github.com/kkjhamb01/courier-management/uaa/proto"
)

func AcceptRequest(ctx context.Context, requestId string, courierId string, desc string) (err error) {

	logger.Infof("AcceptRequest requestId = %v, courierId = %v, desc = %v, ", requestId, courierId, desc)

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
		return proto.Internal.Error(result.Error)
	}

	isValidTransition := status.RequestTransition(requestModel.Status, deliveryPb.RequestStatus_ACCEPTED.String())
	if !isValidTransition {
		err := fmt.Errorf("transition to status accepted for request (id: %v) is not possible. current status: %v", requestId, requestModel.Status)
		logger.Error("failed to accept request", err)
		return proto.DeliveryInvalidStateTransition.Error(err)
	}

	lock, err := storage.LockRequest(ctx, requestId)
	if err != nil {
		logger.Error("the request is not available to accept", err)
		return proto.RequestIsNotAvailableToAccept.Error(err)
	}

	requestModel.Status = deliveryPb.RequestStatus_ACCEPTED.String()
	requestModel.CourierId = courierId
	err = tx.Omit("Locations").
		Save(&requestModel).Error
	if err != nil {
		logger.Error("failed to update request status to accepted", err)
		return proto.Internal.Error(err)
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
		return proto.Internal.Error(err)
	}

	var origin model.RequestLocation
	for _, location := range requestModel.Locations {
		if location.IsOrigin {
			origin = location
			break
		}
	}

	err = publishRequestAccepted(ctx, deliveryPb.RequestAcceptedEvent{
		RequestId:  requestModel.ID,
		CustomerId: requestModel.CustomerId,
		CourierId:  courierId,
		Desc:       desc,
	}, origin.PhoneNumber)
	if err != nil {
		logger.Error("failed to publish offer accepted event", err)
		return proto.Internal.Error(err)
	}

	err = storage.UnlockRequest(ctx, requestId, lock)
	if err != nil {
		logger.Error("failed to unlock request", err)
		return proto.Internal.Error(err)
	}

	return nil
}
