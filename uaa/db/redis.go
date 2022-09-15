package db

import (
	"gitlab.artin.ai/backend/courier-management/common/config"
	"github.com/go-redis/redis/v8"
)

func NewRedisConnection(config config.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.Db,
	})
}
