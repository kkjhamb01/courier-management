package business

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/delivery/db"
	"github.com/kkjhamb01/courier-management/delivery/model"
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
	uaa "github.com/kkjhamb01/courier-management/uaa/proto"
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
