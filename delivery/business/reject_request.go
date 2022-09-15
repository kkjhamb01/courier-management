package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/delivery/db"
	"gitlab.artin.ai/backend/courier-management/delivery/model"
	deliveryPb "gitlab.artin.ai/backend/courier-management/grpc/delivery/go"
	uaa "gitlab.artin.ai/backend/courier-management/uaa/proto"
)

func RejectRequest(ctx context.Context, requestId string, courierId string, desc string) error {
	logger.Infof("RejectRequest requestId = %v, courierId = %v, desc = %v", requestId, courierId, desc)
	var requestModel model.Request
	result := db.MariaDbClient().Where("ID = ?", requestId).First(&requestModel)
	if result.Error != nil {
		logger.Error("failed to check current request status", result.Error)
		return uaa.Internal.Error(result.Error)
	}

	var origin model.RequestLocation
	for _, location := range requestModel.Locations {
		if location.IsOrigin {
			origin = location
			break
		}
	}

	err := publishRequestRejected(ctx, deliveryPb.RequestRejectedEvent{
		RequestId:  requestModel.ID,
		CustomerId: requestModel.CustomerId,
		CourierId:  courierId,
		Desc:       desc,
	}, origin.PhoneNumber)
	if err != nil {
		logger.Error("failed to publish offer rejected event", err)
		return uaa.Internal.Error(err)
	}

	return nil
}
