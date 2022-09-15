package storage

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/delivery/db"
)

const (
	mutexPrefix = "deliveryMutex:"
)

var (
	lockEnabled = true
)

func startLock(key string) (*redsync.Mutex, error) {

	//logger.Infof("startLock key = %v", key)

	mutex := db.RedisMutex(mutexPrefix + key)

	if !lockEnabled {
		return mutex, nil
	}

	if err := mutex.Lock(); err != nil {
		logger.Error("failed to lock the mutex", err)
		return mutex, err
	}

	return mutex, nil
}

func startLocks(keys ...string) ([]*redsync.Mutex, error) {
	mutexes := make([]*redsync.Mutex, len(keys))
	for i, key := range keys {
		mutexes[i] = db.RedisMutex(mutexPrefix + key)

		if !lockEnabled {
			return mutexes, nil
		}

		if err := mutexes[i].Lock(); err != nil {
			logger.Error("failed to lock the mutex", err)
			//rollback:
			releaseLocks(mutexes[:i])
			return nil, err
		}
	}

	return mutexes, nil
}

func releaseLock(mutex *redsync.Mutex) error {
	logger.Infof("releaseLock")

	if !lockEnabled {
		return nil
	}

	if ok, err := mutex.Unlock(); !ok || err != nil {
		logger.Error("failed to unlock the mutex", err)
		return err
	}

	return nil
}

func releaseLocks(mutexes []*redsync.Mutex) error {
	logger.Infof("releaseLocks")
	if !lockEnabled {
		return nil
	}

	for _, mutex := range mutexes {
		if mutex == nil {
			continue
		}
		if ok, err := mutex.Unlock(); !ok || err != nil {
			logger.Error("failed to unlock the mutex", err)
			return err
		}
	}

	return nil
}
