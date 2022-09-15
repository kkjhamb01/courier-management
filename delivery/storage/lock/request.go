package storage

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/delivery/db"
)

const (
	requestPrefix = "delivery:request"

	requestProcessing    = "processing"
	requestNotProcessing = ""
)

func LockRequest(ctx context.Context, requestId string) (*redsync.Mutex, error) {

	//logger.Infof("LockRequest requestId = %v", requestId)

	m, err := startLock(requestPrefix + requestId)
	if err != nil {
		//logger.Error("failed to get the lock", err)
		return nil, err
	}

	v, err := db.RedisClient().Get(ctx, requestPrefix+requestId).Result()
	if err != nil {
		//logger.Error("failed to get the value", err)
		return nil, err
	}

	if v != requestNotProcessing {
		//err = errors.New("request is processing now")
		return nil, err
	}

	err = db.RedisClient().Set(ctx, requestPrefix+requestId, requestProcessing, -1).Err()
	if err != nil {
		//logger.Error("failed to set the value", err)
		return nil, err
	}

	return m, nil
}

func UnlockRequest(ctx context.Context, requestId string, lock *redsync.Mutex) error {

	logger.Infof("UnlockRequest requestId = %v", requestId)

	defer func() {
		err := releaseLock(lock)
		if err != nil {
			logger.Error("failed to release the lock", err)
		}
	}()

	err := db.RedisClient().Set(ctx, requestPrefix+requestId, requestNotProcessing, -1).Err()
	if err != nil {
		logger.Error("failed to set the value", err)
		return err
	}

	return nil
}

func InitLockRequest(ctx context.Context, requestId string) error {

	logger.Infof("InitLockRequest requestId = %v", requestId)

	m, err := startLock(requestId)
	if err != nil {
		logger.Error("failed to get the lock", err)
		return err
	}

	defer func() {
		err = releaseLock(m)
		if err != nil {
			logger.Error("failed to release the lock", err)
		}
	}()

	err = db.RedisClient().Set(ctx, requestPrefix+requestId, requestNotProcessing, -1).Err()
	if err != nil {
		logger.Error("failed to set the value", err)
		return err
	}

	return nil
}
