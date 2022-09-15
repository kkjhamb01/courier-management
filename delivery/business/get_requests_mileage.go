package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/delivery/db"
	uaa "gitlab.artin.ai/backend/courier-management/uaa/proto"
	"time"
)

func GetCourierRequestsMileage(ctx context.Context, userId string, from time.Time, to time.Time) (int, error) {

	logger.Infof("GetCourierRequestsMileage userId = %v, from = %v, to = %v", userId, from, to)

	var distanceInKm float64
	queryText := `
	SELECT SUM(
	       111.111 *
	       DEGREES(ACOS(LEAST(1.0, COS(RADIANS(origin.Lat)) * COS(RADIANS(dest.Lat)) * COS(RADIANS(origin.Lon - dest.Lon))
	     + SIN(RADIANS(origin.Lat)) * SIN(RADIANS(dest.Lat)))))) AS distance_in_km
	from ride_statuses
	join requests req on ride_statuses.request_id = req.id
	join request_locations origin on
	    (ride_statuses.request_id = origin.request_id and origin.is_origin = 1)
	join request_locations dest on
	    (ride_statuses.request_id = dest.request_id and dest.is_origin = 0 and
	     -- to get the last destination
	     dest.order = (select count(id) from request_locations where request_id = req.id and is_origin = 0))
	where ride_statuses.status = 'COMPLETED'
	  and ride_statuses.created_at >= ? and ride_statuses.created_at < ? and req.courier_id = ? 
`

	err := db.MariaDbClient().Raw(queryText, from, to, userId).Scan(&distanceInKm).Error
	if err != nil {
		logger.Error("failed to fetch requests mileage", err)
		return 0, uaa.Internal.Error(err)
	}

	return int(distanceInKm), nil
}

func GetCustomerRequestsMileage(ctx context.Context, userId string, from time.Time, to time.Time) (int, error) {

	logger.Infof("GetCustomerRequestsMileage userId = %v, from = %v, to = %v", userId, from, to)

	var distanceInKm float64
	queryText := `
	SELECT SUM(
	       111.111 *
	       DEGREES(ACOS(LEAST(1.0, COS(RADIANS(origin.Lat)) * COS(RADIANS(dest.Lat)) * COS(RADIANS(origin.Lon - dest.Lon))
	     + SIN(RADIANS(origin.Lat)) * SIN(RADIANS(dest.Lat)))))) AS distance_in_km
	from ride_statuses
	join requests req on ride_statuses.request_id = req.id
	join request_locations origin on
	    (ride_statuses.request_id = origin.request_id and origin.is_origin = 1)
	join request_locations dest on
	    (ride_statuses.request_id = dest.request_id and dest.is_origin = 0 and
	     -- to get the last destination
	     dest.order = (select count(id) from request_locations where request_id = req.id and is_origin = 0))
	where ride_statuses.status = 'COMPLETED'
	  and ride_statuses.created_at >= ? and ride_statuses.created_at < ? and req.customer_id = ? 
`

	err := db.MariaDbClient().Raw(queryText, from, to, userId).Scan(&distanceInKm).Error
	if err != nil {
		logger.Error("failed to fetch requests mileage", err)
		return 0, uaa.Internal.Error(err)
	}

	return int(distanceInKm), nil
}
