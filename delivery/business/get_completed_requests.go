package business

import (
	"context"
	"time"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/delivery/db"
	"github.com/kkjhamb01/courier-management/delivery/model"
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
	uaa "github.com/kkjhamb01/courier-management/uaa/proto"
)

func GetCourierCompletedRequests(ctx context.Context, userId string, from time.Time, to time.Time) (int, error) {

	logger.Infof("GetCourierCompletedRequests userId = %v, from = %v, to = %v", userId, from, to)

	var completedRequests int64
	err := db.MariaDbClient().Model(&model.RideStatus{}).
		Joins("join requests on requests.id = ride_statuses.request_id").
		Where("ride_statuses.status = ? AND "+
			"ride_statuses.created_at >= ? AND "+
			"ride_statuses.created_at <= ? AND "+
			"requests.courier_id = ?",
			deliveryPb.RequestStatus_COMPLETED.String(),
			from,
			to,
			userId).
		Count(&completedRequests).
		Distinct("ride_statuses.request_id").
		Error
	if err != nil {
		logger.Error("failed to fetch number of completed statuses", err)
		return 0, uaa.Internal.Error(err)
	}

	return int(completedRequests), nil
}

func GetCustomerCompletedRequests(ctx context.Context, userId string, from time.Time, to time.Time) (int, error) {
	var completedRequests int64
	err := db.MariaDbClient().Model(&model.RideStatus{}).
		Joins("join requests on requests.id = ride_statuses.request_id").
		Where("ride_statuses.status = ? AND "+
			"ride_statuses.created_at >= ? AND "+
			"ride_statuses.created_at <= ? AND "+
			"requests.customer_id = ?",
			deliveryPb.RequestStatus_COMPLETED.String(),
			from,
			to,
			userId).
		Count(&completedRequests).
		Distinct("ride_statuses.request_id").
		Error
	if err != nil {
		logger.Error("failed to fetch number of completed statuses", err)
		return 0, uaa.Internal.Error(err)
	}

	return int(completedRequests), nil
}
