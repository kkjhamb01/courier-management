package business

import (
	"context"
	"errors"
	"fmt"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/delivery/db"
	"github.com/kkjhamb01/courier-management/delivery/model"
	"github.com/kkjhamb01/courier-management/delivery/status"
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
	"github.com/kkjhamb01/courier-management/uaa/proto"
)

var cancelAfterPickUpErr = errors.New("requests with any picked up in their history cannot be cancelled")

func CancelRequest(ctx context.Context, requestId string, reason deliveryPb.CancelReason,
	cancelledBy deliveryPb.CancelledBy, updatingUserId string, description string) (err error) {

	logger.Infof("CancelRequest requestId = %v, reason = %v, cancelledBy = %v, updatingUserId = %v, description = %v", requestId, reason, cancelledBy, updatingUserId, description)

	tx := db.MariaDbClient().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	// check whether the request has been in PICKED UP status before
	var ridePickedUpStatuses []model.RideStatus
	err = tx.Where("request_id = ? AND status = ?", requestId, deliveryPb.RequestStatus_PICKED_UP.String()).
		Find(&ridePickedUpStatuses).Error
	logger.Info("THIS IS RideStatuses on CANCEL", tag.Obj("obj", ridePickedUpStatuses))
	if err != nil {
		logger.Error("failed to fetch ride statuses", err)
		return proto.Internal.Error(err)
	}
	if len(ridePickedUpStatuses) > 0 {
		logger.Error("failed to cancel request", cancelAfterPickUpErr)
		return proto.RequestIsInPickedupState.Error(cancelAfterPickUpErr)
	}

	//TODO SUGGESTION: category for reasons
	//TODO SUGGESTION: validate category for user
	var requestModel model.Request
	result := tx.Where("ID = ?", requestId).First(&requestModel)
	if result.Error != nil {
		logger.Error("failed to check current request status", result.Error)
		return proto.Internal.Error(result.Error)
	}

	var nextStatus string
	if cancelledBy == deliveryPb.CancelledBy_COURIER {
		nextStatus = deliveryPb.RequestStatus_COURIER_CANCELLED.String()
	} else {
		nextStatus = deliveryPb.RequestStatus_CUSTOMER_CANCELLED.String()
	}
	isValidTransition := status.RequestTransition(requestModel.Status, nextStatus)
	if !isValidTransition {
		err := fmt.Errorf("transition to status cancelled for request (id: %v) is not possible. current status: %v", requestId, requestModel.Status)
		logger.Error("failed to cancel request", err)
		return proto.DeliveryInvalidStateTransition.Error(err)
	}

	err = tx.Model(&requestModel).
		Updates(map[string]interface{}{
			"cancelling_reason": reason.String(),
			"cancelled_by":      cancelledBy.String(),
			"updated_by":        updatingUserId,
			"status":            nextStatus,
		}).Error
	if err != nil {
		logger.Error("failed to update request cancelling columns", err)
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
		CancellationNote: description,
	}
	err = tx.Create(&rideStatus).Error
	if err != nil {
		logger.Error("failed to create ride status", err)
		return proto.Internal.Error(err)
	}

	err = publishRequestCancelled(ctx, deliveryPb.RequestCancelledEvent{
		RequestId:    requestModel.ID,
		CustomerId:   requestModel.CustomerId,
		CourierId:    requestModel.CourierId,
		CancelReason: reason,
		CancelledBy:  cancelledBy,
		UpdatedBy:    updatingUserId,
	}, description, originLocation.ToProto(), requestModel.HumanReadableId)
	if err != nil {
		logger.Error("failed to publish offer cancelled event", err)
		return proto.Internal.Error(err)
	}

	return nil
}
