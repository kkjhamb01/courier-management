package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/delivery/db"
	uaa "gitlab.artin.ai/backend/courier-management/uaa/proto"
	"time"
)

func GetCourierRequestsDuration(ctx context.Context, userId string, from time.Time, to time.Time) (time.Duration, error) {

	logger.Infof("GetCourierRequestsDuration userId = %v, from = %v, to = %v", userId, from, to)

	var totalTimeMinutes int64
	err := db.MariaDbClient().Raw("SELECT TIMESTAMPDIFF(MINUTE , start_log.created_at, end_log.created_at) AS TimeInMin "+
		"FROM ride_statuses AS start_log "+
		"INNER JOIN ride_statuses AS end_log "+
		"ON (start_log.request_id = end_log.request_id AND start_log.status = 'ACCEPTED' AND end_log.status = 'COMPLETED') "+
		"INNER JOIN requests ON requests.id = start_log.request_id "+
		"WHERE start_log.created_at >= ? AND end_log.created_at <= ? AND "+
		"requests.courier_id = ?", from, to, userId).Scan(&totalTimeMinutes).
		Error
	if err != nil {
		logger.Error("failed to fetch requests duration", err)
		return 0, uaa.Internal.Error(err)
	}

	totalTime := time.Minute * time.Duration(totalTimeMinutes)
	return totalTime, nil
}

func GetCustomerRequestsDuration(ctx context.Context, userId string, from time.Time, to time.Time) (time.Duration, error) {

	logger.Infof("GetCustomerRequestsDuration userId = %v, from = %v, to = %v", userId, from, to)

	var totalTimeMinutes int64
	err := db.MariaDbClient().Raw("SELECT TIMESTAMPDIFF(MINUTE , start_log.created_at, end_log.created_at) AS TimeInMin "+
		"FROM ride_statuses AS start_log "+
		"INNER JOIN ride_statuses AS end_log "+
		"ON (start_log.request_id = end_log.request_id AND start_log.status = 'ACCEPTED' AND end_log.status = 'COMPLETED') "+
		"INNER JOIN requests ON requests.id = start_log.request_id "+
		"WHERE start_log.created_at >= ? AND end_log.created_at <= ? AND "+
		"requests.customer_id = ?", from, to, userId).Scan(&totalTimeMinutes).
		Error
	if err != nil {
		logger.Error("failed to fetch requests duration", err)
		return 0, uaa.Internal.Error(err)
	}

	totalTime := time.Minute * time.Duration(totalTimeMinutes)
	return totalTime, nil
}
