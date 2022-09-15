package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	commonPb "github.com/kkjhamb01/courier-management/grpc/common/go"
	"github.com/kkjhamb01/courier-management/offering/db"
)

const (
	fieldCourierType = "courier_type"
)

func (s storageImpl) SetCourierLocation(ctx context.Context, location commonPb.Location, courierId string) error {
	return s.SetCourierLocationByCourierType(ctx, location, commonPb.VehicleType_ANY, courierId)
}

func (s storageImpl) SetCourierLocationByCourierType(ctx context.Context, location commonPb.Location, courierType commonPb.VehicleType, courierId string) error {
	key := courierId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return err
	}
	defer s.releaseLock(mutex)

	err = db.Tile38Client().Keys.Set(config.Tile38().CourierLocationCollection, key).
		Point(location.Lat, location.Lon).
		Field(fieldCourierType, float64(courierType)).
		Do()
	if err != nil {
		logger.Error("failed to add courier location into tile38", err)
		return err
	}

	if err := s.redisImpl.pipeliner.Set(ctx, fmt.Sprintf("alive-%s", courierId), 1, time.Duration(30*time.Minute)).Err(); err != nil {
		logger.Error("failed to set courier alive location in redis", err)
		return err
	}

	return nil
}

func (s storageImpl) GetCourierLocation(ctx context.Context, courierId string) (commonPb.Location, error) {
	mutex, err := s.startLock(courierId)
	if err != nil {
		logger.Error("failed to start lock", err)
		return commonPb.Location{}, err
	}
	defer s.releaseLock(mutex)

	res, err := db.Tile38Client().Keys.Get(config.Tile38().CourierLocationCollection, courierId, true)
	if err != nil {
		logger.Error("failed to get tile38 location", err)
		return commonPb.Location{}, err
	}

	if res.Object.Geometry.Point == nil {
		err = errors.New("no point was returned")
		logger.Error("failed to get location", err)
		return commonPb.Location{}, err
	}

	return commonPb.Location{
		Lat: res.Object.Geometry.Point[1],
		Lon: res.Object.Geometry.Point[0],
	}, err
}

func (s storageImpl) GetNearbyCouriersByCourierType(ctx context.Context, location commonPb.Location, radiusMeter int, courierType commonPb.VehicleType) ([]commonPb.CourierLocation, error) {
	if courierType == commonPb.VehicleType_ANY {
		return s.GetNearbyCouriers(ctx, location, radiusMeter)
	}

	res, err := db.Tile38Client().Search.Nearby(config.Tile38().CourierLocationCollection, location.Lat, location.Lon, float64(radiusMeter)).
		Wherein(fieldCourierType, float64(courierType)).
		Do()
	if err != nil {
		logger.Error("failed to search nearBy", err)
		return nil, err
	}

	couriers := make([]commonPb.CourierLocation, res.Count)
	for i := 0; i < res.Count; i++ {
		couriers[i] = commonPb.CourierLocation{
			Courier: &commonPb.Courier{
				Id: res.Objects[i].ID,
			},
			Location: &commonPb.Location{
				Lat: res.Objects[i].Object.Geometry.Point[1],
				Lon: res.Objects[i].Object.Geometry.Point[0],
			},
		}
		logger.Infof("storage.GetNearbyCouriersByCourierType, i: %v, courier:%v", i, couriers[i])
		courierTypeIndex := findFieldIndex(res.Fields, fieldCourierType)
		if courierTypeIndex >= 0 {
			couriers[i].Courier.VehicleType = commonPb.VehicleType(res.Objects[i].Fields[courierTypeIndex])
		}
	}
	logger.Infof("storage.GetNearbyCouriersByCourierType DONE, couriers: %v", couriers)

	return couriers, nil
}

func (s storageImpl) GetNearbyCouriersByCourierTypeAndMaxResult(ctx context.Context, location commonPb.Location, radiusMeter int, courierType commonPb.VehicleType, maxResult int) ([]commonPb.CourierLocation, error) {
	maxResultSpecified := maxResult > 0
	courierTypeSpecified := courierType != commonPb.VehicleType_ANY

	switch {
	case !maxResultSpecified && !courierTypeSpecified:
		return s.GetNearbyCouriers(ctx, location, radiusMeter)
	case !maxResultSpecified && courierTypeSpecified:
		return s.GetNearbyCouriersByCourierType(ctx, location, radiusMeter, courierType)
	case maxResultSpecified && !courierTypeSpecified:
		return s.GetNearbyCouriersByMaxResult(ctx, location, radiusMeter, maxResult)
	}

	query := db.Tile38Client().Search.Nearby(config.Tile38().CourierLocationCollection, location.Lat, location.Lon, float64(radiusMeter)).
		Wherein(fieldCourierType, float64(courierType)).
		Limit(maxResult)

	//if len(excludedCourierIds) > 0 {
	//	excludedCouriers := strings.Join(excludedCourierIds, "|")
	//	query = query.Match("!(" + excludedCouriers + ")")
	//}

	res, err := query.Do()
	if err != nil {
		logger.Error("failed to search nearBy", err)
		return nil, err
	}

	couriers := make([]commonPb.CourierLocation, res.Count)
	for i := 0; i < res.Count; i++ {
		couriers[i] = commonPb.CourierLocation{
			Courier: &commonPb.Courier{
				Id: res.Objects[i].ID,
			},
			Location: &commonPb.Location{
				Lat: res.Objects[i].Object.Geometry.Point[1],
				Lon: res.Objects[i].Object.Geometry.Point[0],
			},
		}
		courierTypeIndex := findFieldIndex(res.Fields, fieldCourierType)
		if courierTypeIndex >= 0 {
			couriers[i].Courier.VehicleType = commonPb.VehicleType(res.Objects[i].Fields[courierTypeIndex])
		}
	}

	return couriers, nil
}

func (s storageImpl) GetNearbyCouriersByMaxResult(ctx context.Context, location commonPb.Location, radiusMeter int, maxResult int) ([]commonPb.CourierLocation, error) {
	if maxResult < 1 {
		return s.GetNearbyCouriers(ctx, location, radiusMeter)
	}

	res, err := db.Tile38Client().Search.Nearby(config.Tile38().CourierLocationCollection, location.Lat, location.Lon, float64(radiusMeter)).
		Limit(maxResult).
		Do()
	if err != nil {
		logger.Error("failed to search nearBy", err)
		return nil, err
	}

	couriers := make([]commonPb.CourierLocation, res.Count)
	for i := 0; i < res.Count; i++ {
		couriers[i] = commonPb.CourierLocation{
			Courier: &commonPb.Courier{
				Id: res.Objects[i].ID,
			},
			Location: &commonPb.Location{
				Lat: res.Objects[i].Object.Geometry.Point[1],
				Lon: res.Objects[i].Object.Geometry.Point[0],
			},
		}
		courierTypeIndex := findFieldIndex(res.Fields, fieldCourierType)
		if courierTypeIndex >= 0 {
			couriers[i].Courier.VehicleType = commonPb.VehicleType(res.Objects[i].Fields[courierTypeIndex])
		}
	}

	return couriers, nil
}

func (s storageImpl) GetNearbyCouriers(ctx context.Context, location commonPb.Location, radiusMeter int) ([]commonPb.CourierLocation, error) {
	res, err := db.Tile38Client().Search.Nearby(config.Tile38().CourierLocationCollection, location.Lat, location.Lon, float64(radiusMeter)).
		Do()
	if err != nil {
		logger.Error("failed to search nearBy", err)
		return nil, err
	}

	couriers := make([]commonPb.CourierLocation, res.Count)
	for i := 0; i < res.Count; i++ {
		couriers[i] = commonPb.CourierLocation{
			Courier: &commonPb.Courier{
				Id: res.Objects[i].ID,
			},
			Location: &commonPb.Location{
				Lat: res.Objects[i].Object.Geometry.Point[1],
				Lon: res.Objects[i].Object.Geometry.Point[0],
			},
		}
		courierTypeIndex := findFieldIndex(res.Fields, fieldCourierType)
		if courierTypeIndex >= 0 {
			couriers[i].Courier.VehicleType = commonPb.VehicleType(res.Objects[i].Fields[courierTypeIndex])
		}
	}

	return couriers, nil
}

func findFieldIndex(fields []string, key string) int {
	for i, field := range fields {
		if field == key {
			return i
		}
	}

	return -1
}
