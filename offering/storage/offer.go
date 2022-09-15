package storage

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/offering/db"
	"strconv"
)

const (
	doneOffersPfx           = "doneOffers_"
	offerPendingCouriersPfx = "offerPendingCouriers_"
	offerSentCouriersPfx    = "offerSentCouriers_"
	offerMaxRetriesPfx      = "offerMaxRetries_"
)

func (s storageImpl) AddPendingCouriersToOffer(ctx context.Context, offerId string, courierIds []string) error {
	key := offerPendingCouriersPfx + offerId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return err
	}
	defer s.releaseLock(mutex)

	courierIdsInterface := make([]interface{}, len(courierIds))
	for i, d := range courierIds {
		courierIdsInterface[i] = d
	}

	added, err := db.OfferRedisClient().SAdd(ctx, key, courierIdsInterface...).Result()
	if err != nil {
		logger.Error("failed to add courier ids to offer id", err, tag.Obj("courierIds", courierIds), tag.Str("offerId", offerId))
		return err
	}

	if added != int64(len(courierIds)) {
		logger.Error("not all courier ids were added to offer id", err,
			tag.Int("expected adds", len(courierIds)),
			tag.Int64("actual added", added),
			tag.Obj("courierIds", courierIds))

		if rollbackErr := db.OfferRedisClient().Del(ctx, offerPendingCouriersPfx+offerId).Err(); rollbackErr != nil {
			logger.Error("failed to rollback added courierIds to offerId", rollbackErr)
		}
		return err
	}

	return nil
}

func (s storageImpl) RemoveOfferAndAllPendingCouriers(ctx context.Context, offerId string) error {
	key := offerPendingCouriersPfx + offerId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return err
	}
	defer s.releaseLock(mutex)

	courierIds, err := db.OfferRedisClient().SMembers(ctx, key).Result()
	if err != nil {
		logger.Error("failed to fetch the offer courier Ids", err, tag.Str("offerId", offerId))
		return err
	}

	for _, courierId := range courierIds {
		if err := db.OfferRedisClient().SRem(ctx, key, courierId).Err(); err != nil {
			logger.Error("failed to courier id from offer", err, tag.Str("offerId", offerId), tag.Str("courierId", courierId))
		}
	}

	if err := db.OfferRedisClient().Del(ctx, key).Err(); err != nil {
		logger.Error("failed to delete offer key", err, tag.Str("offerId", offerId))
		return err
	}

	return nil
}

func (s storageImpl) RemovePendingCourierFromOffer(ctx context.Context, offerId string, courierId string) error {
	key := offerPendingCouriersPfx + offerId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return err
	}
	defer s.releaseLock(mutex)

	affected, err := db.OfferRedisClient().SRem(ctx, key, courierId).Result()
	if err != nil {
		logger.Error("failed to remove courier from offer", err, tag.Str("offerId", offerId), tag.Str("courierId", courierId))
		return err
	}

	if affected == 0 {
		err = errors.New("SRem did not affect any data")
		logger.Error("failed to remove courier from offer", err, tag.Str("offerId", offerId), tag.Str("courierId", courierId))
		return err
	}

	return nil
}

func (s storageImpl) GetOfferPendingCouriers(ctx context.Context, offerId string) ([]string, error) {
	key := offerPendingCouriersPfx + offerId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return nil, err
	}
	defer s.releaseLock(mutex)

	courierIds, err := db.OfferRedisClient().SMembers(ctx, key).Result()
	if err != nil {
		logger.Error("failed to get the offer courier Ids", err, tag.Str("offerId", offerId))
		return nil, err
	}

	return courierIds, nil
}

func (s storageImpl) IsOfferCreatedEver(ctx context.Context, offerId string) (bool, error) {
	key := offerPendingCouriersPfx + offerId

	keys, err := db.OfferRedisClient().Exists(ctx, key).Result()
	if err != nil {
		logger.Error("failed to check whether the key exists", err, tag.Str("key", key))
		return false, err
	}

	return keys != 0, nil
}

func (s storageImpl) AddOfferSentCourier(ctx context.Context, offerId string, courierIds []string) error {
	key := offerSentCouriersPfx + offerId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return err
	}
	defer s.releaseLock(mutex)

	courierIdsInterface := make([]interface{}, len(courierIds))
	for i, d := range courierIds {
		courierIdsInterface[i] = d
	}

	_, err = db.OfferRedisClient().SAdd(ctx, key, courierIdsInterface...).Result()
	if err != nil {
		logger.Error("failed to add sent courier ids to offer id", err, tag.Obj("courierIds", courierIds), tag.Str("offerId", offerId))
		return err
	}

	return nil
}

func (s storageImpl) GetOfferSentCouriers(ctx context.Context, offerId string) ([]string, error) {
	key := offerSentCouriersPfx + offerId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return nil, err
	}
	defer s.releaseLock(mutex)

	courierIds, err := db.OfferRedisClient().SMembers(ctx, key).Result()
	if err != nil {
		logger.Error("failed to get the offer sent courier Ids", err, tag.Str("offerId", offerId))
		return nil, err
	}

	return courierIds, nil
}

func (s storageImpl) IsCourierPendingOffer(ctx context.Context, offerId string, courierId string) (bool, error) {
	courierIds, err := s.GetOfferPendingCouriers(ctx, offerId)
	if err != nil {
		logger.Error("failed to get offer pending couriers", err)
		return false, err
	}

	for _, pendingCourierId := range courierIds {
		if pendingCourierId == courierId {
			return true, nil
		}
	}

	return false, nil
}

func (s storageImpl) GetOfferRetries(ctx context.Context, offerId string) (int, error) {
	key := offerMaxRetriesPfx + offerId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return 0, err
	}
	defer s.releaseLock(mutex)

	retriesS, err := db.OfferRedisClient().Get(ctx, key).Result()
	if err == redis.Nil {
		retriesS = "0"
	} else if err != nil {
		logger.Error("failed to get offer retries", err, tag.Str("offerId", offerId))
		return 0, err
	}

	retries, err := strconv.Atoi(retriesS)
	if err != nil {
		logger.Error("wrong va", err, tag.Str("offerId", offerId))
		return 0, err
	}

	return retries, nil
}

func (s storageImpl) IncreaseOfferRetries(ctx context.Context, offerId string) (int, error) {
	key := offerMaxRetriesPfx + offerId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return 0, err
	}
	defer s.releaseLock(mutex)

	retries, err := db.OfferRedisClient().Incr(ctx, key).Result()
	if err != nil {
		logger.Error("failed to increase offer retries", err, tag.Str("offerId", offerId))
		return 0, err
	}

	return int(retries), nil
}

func (s storageImpl) ResetOfferRetries(ctx context.Context, offerId string) error {
	key := offerMaxRetriesPfx + offerId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return err
	}
	defer s.releaseLock(mutex)

	_, err = db.OfferRedisClient().Set(ctx, key, 0, 0).Result()
	if err != nil {
		logger.Error("failed to reset offer retries", err, tag.Str("offerId", offerId))
		return err
	}

	return nil
}

func (s storageImpl) NumberOfPendingOffers(ctx context.Context) (int, error) {
	var pendingOffersCounter int

	var keys []string
	var courser uint64 // initially zero
	var err error

	continueScan := true
	for continueScan {
		keys, courser, err = db.OfferRedisClient().Scan(ctx, courser, offerPendingCouriersPfx+"*", 100).Result()
		if err != nil {
			logger.Error("failed to scan number of pending offers", err)
			return 0, err
		}

		pendingOffersCounter += len(keys)
		continueScan = courser != 0
	}
	return pendingOffersCounter, nil
}

func (s storageImpl) CloseOffer(ctx context.Context, offerId string) error {
	key := doneOffersPfx + offerId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return err
	}
	defer s.releaseLock(mutex)

	_, err = db.OfferRedisClient().Set(ctx, key, "is done", -1).Result()
	if err != nil {
		logger.Error("failed to set done offer", err, tag.Str("offerId", offerId))
		return err
	}

	return nil
}

func (s storageImpl) IsOfferClosed(ctx context.Context, offerId string) (bool, error) {
	key := doneOffersPfx + offerId

	mutex, err := s.startLock(key)
	if err != nil {
		logger.Error("failed to start lock", err)
		return false, err
	}
	defer s.releaseLock(mutex)

	err = db.OfferRedisClient().Get(ctx, key).Err()
	switch err {
	case redis.Nil:
		return false, nil
	case nil:
		return true, nil
	default:
		logger.Error("failed to get closed offer", err, tag.Str("offerId", offerId))
		return false, err
	}
}
