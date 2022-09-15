package business

import (
	"context"
	"fmt"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/delivery/db"
	"github.com/kkjhamb01/courier-management/delivery/model"
	"github.com/kkjhamb01/courier-management/delivery/status"
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
	uaa "github.com/kkjhamb01/courier-management/uaa/proto"
)

func RequestPickedUp(ctx context.Context, requestId string, name string, signature string, description string) (err error) {

	logger.Infof("RequestPickedUp requestId = %v, description = %v, name = %v, signature = %v", requestId, description, name, signature)

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

	isValidTransition := status.RequestTransition(requestModel.Status, deliveryPb.RequestStatus_PICKED_UP.String())
	if !isValidTransition {
		err := fmt.Errorf("transition to status picked up for request (id: %v) is not possible. current status: %v", requestId, requestModel.Status)
		logger.Error("failed to process request picked up", err)
		return uaa.DeliveryInvalidStateTransition.Error(err)
	}

	requestModel.Status = deliveryPb.RequestStatus_PICKED_UP.String()
	err = tx.Omit("Locations").
		Save(&requestModel).Error
	if err != nil {
		logger.Error("failed to update request status to picked up", err)
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

	rideConfirmation := model.RideConfirmation{
		RideLocationId: originLocation.ID,
		RequestId:      requestModel.ID,
		Name:           name,
		Signature:      signature,
	}
	err = tx.Create(&rideConfirmation).Error
	if err != nil {
		logger.Error("failed to create ride confirmation", err)
		return uaa.Internal.Error(err)
	}

	err = publishRequestPickedUp(ctx, deliveryPb.RequestPickedUpEvent{
		RequestId: requestModel.ID,
		Desc:      description,
	}, requestModel.CustomerId, originLocation.ToProto(), requestModel.HumanReadableId)
	if err != nil {
		logger.Error("failed to publish request picked up event", err)
		return uaa.Internal.Error(err)
	}

	return nil
}
