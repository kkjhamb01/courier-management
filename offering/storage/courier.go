package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	commonPb "github.com/kkjhamb01/courier-management/grpc/common/go"
	"github.com/kkjhamb01/courier-management/offering/db"
)

const (
	courierPendingOffers = "courierPendingOffers_"
	courierStatus        = "courierStatus_"
	courierTypePrefix    = "courierType_"
)

func (s storageImpl) GetCouriersStatusByCourierETAs(ctx context.Context, courierETAs []commonPb.CourierETA) ([]commonPb.CourierStatus, error) {
	couriers := make([]string, len(courierETAs))
	for i, courierETA := range courierETAs {
		couriers[i] = courierETA.Courier.Id
	}

	return s.GetCouriersStatusByIds(ctx, couriers)
}

func (s storageImpl) GetCouriersStatus(ctx context.Context, couriers []commonPb.Courier) ([]commonPb.CourierStatus, error) {
	courierIds := make([]string, len(couriers), len(couriers))
	for i, courier := range couriers {
		courierIds[i] = courier.Id
	}

	return s.GetCouriersStatusByIds(ctx, courierIds)
}

func (s storageImpl) GetCouriersStatusByIds(ctx context.Context, courierIds []string) ([]commonPb.CourierStatus, error) {
	for _, courier := range courierIds {
		courier = courierStatus + courier
	}

	mutexes, err := s.startLocks(courierIds...)
	if err != nil {
		logger.Error("failed to start locks", err)
		return nil, err
	}
	defer s.releaseLocks(mutexes)

	statusNames, err := db.OfferRedisClient().MGet(ctx, courierIds...).Result()
	if err != nil {
		logger.Error("failed to get courier status", err, tag.Obj("courierIds", courierIds))
		return nil, err
	}

	statusValues := make([]commonPb.CourierStatus, len(courierIds), len(courierIds))

	for i, nullableStatusName := range statusNames {
		// check if any value (status) is set for the key (courier id)
		if nullableStatusName == nil {
			// assumption: if no status is set on the couriers db, the driver is available
			statusValues[i] = commonPb.CourierStatus_AVAILABLE
			nullableStatusName = commonPb.CourierStatus_AVAILABLE.String()
		}
		statusName := fmt.Sprintf("%v", nullableStatusName)

		statusValue, ok := commonPb.CourierStatus_value[statusName]
		if !ok {
			err = errors.New("fetched courier status name is invalid")
			logger.Error("fetched courier status name is invalid", err, tag.Str("StatusName", statusName))
			return nil, err
		}

		statusValues[i] = commonPb.CourierStatus(statusValue)
	}

	return statusValues, nil
}

func (s storageImpl) ChangeCourierStatus(ctx context.Context, courierId string, status commonPb.CourierStatus) error {
	key := courierStatus + courierId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return err
	}
	defer s.releaseLock(mutex)

	if err = db.OfferRedisClient().Set(ctx, key, status.String(), 0).Err(); err != nil {
		logger.Error("failed to set courier status", err, tag.Obj("status", status), tag.Str("courierId", courierId))
		return err
	}

	return nil
}

func (s storageImpl) AddPendingOfferToCourier(ctx context.Context, offerId string, courierId string) error {
	key := courierPendingOffers + courierId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return err
	}
	defer s.releaseLock(mutex)

	_, err = db.OfferRedisClient().SAdd(ctx, key, offerId).Result()
	if err != nil {
		logger.Error("failed to add offer id to the courier", err, tag.Str("courierId", courierId), tag.Str("offerId", offerId))
		return err
	}

	return nil
}

func (s storageImpl) GetCourierPendingOffers(ctx context.Context, courierId string) ([]string, error) {
	key := courierPendingOffers + courierId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return nil, err
	}
	defer s.releaseLock(mutex)

	offerIds, err := db.OfferRedisClient().SMembers(ctx, key).Result()
	if err != nil {
		logger.Error("failed to get offer id couriers", err, tag.Str("courierId", courierId))
		return nil, err
	}

	return offerIds, nil
}

func (s storageImpl) RemovePendingOfferFromCourier(ctx context.Context, offerId string, courierId string) error {
	key := courierPendingOffers + courierId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return err
	}
	defer s.releaseLock(mutex)

	affected, err := db.OfferRedisClient().SRem(ctx, key, offerId).Result()
	if err != nil {
		logger.Error("failed to remove offer id from the courier", err, tag.Str("courierId", courierId), tag.Str("offerId", offerId))
		return err
	}

	if affected == 0 {
		err = errors.New("SRem did not affect any data")
		logger.Error("failed to remove offer id from the courier", err, tag.Str("courierId", courierId), tag.Str("offerId", offerId))
		return err
	}

	return nil
}

func (s storageImpl) GetCourierType(ctx context.Context, courierId string) (commonPb.VehicleType, error) {
	key := courierTypePrefix + courierId

	courierTypeStr, err := db.OfferRedisClient().Get(ctx, key).Result()
	if err != nil {
		logger.Error("failed to get courier type", err, tag.Str("courierId", courierId))
		return commonPb.VehicleType_ANY, err
	}

	courierType, ok := commonPb.VehicleType_value[courierTypeStr]
	if !ok {
		logger.Error("failed to convert courierType string value to courier type", err)
		return commonPb.VehicleType_ANY, err
	}

	return commonPb.VehicleType(courierType), nil
}

func (s storageImpl) SetCourierType(ctx context.Context, courierId string, courierType commonPb.VehicleType) error {
	key := courierTypePrefix + courierId

	err := db.OfferRedisClient().Set(ctx, key, courierType.String(), -1).Err()
	if err != nil {
		logger.Error("failed to set courier type", err, tag.Str("courierId", courierId))
		return err
	}

	return nil
}
