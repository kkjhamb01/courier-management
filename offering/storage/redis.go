package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/offering/db"
)

var (
	NoRecordFound = redis.Nil
)

type redisImpl struct {
	pipeliner redis.Pipeliner
}

func (r redisImpl) rollback() error {
	err := r.pipeliner.Discard()
	if err != nil {
		logger.Error("redis: failed to discard transaction (multi)", err)
	}

	return err
}

func (r redisImpl) commit(ctx context.Context) error {
	commands, err := r.pipeliner.Exec(ctx)
	if err != nil {
		logger.Error("redis: failed to exec transaction (multi)", err, tag.Obj("commands", commands))
	}

	return err
}

func createRedisTx() redisImpl {
	return redisImpl{
		pipeliner: db.OfferRedisClient().TxPipeline(),
	}
}
