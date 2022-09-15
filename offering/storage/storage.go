package storage

import (
	"context"
)

type storageImpl struct {
	redisImpl
	tile38Impl
}

func (s storageImpl) Rollback() error {
	err := s.redisImpl.rollback()
	if err != nil {
		return err
	}

	err = s.tile38Impl.rollback()
	if err != nil {
		return err
	}

	return nil
}

func (s storageImpl) Commit(ctx context.Context) error {
	err := s.redisImpl.commit(ctx)
	if err != nil {
		return err
	}

	err = s.tile38Impl.commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func CreateTx() storageImpl {
	return storageImpl{
		redisImpl:  createRedisTx(),
		tile38Impl: createTile38Tx(),
	}
}
