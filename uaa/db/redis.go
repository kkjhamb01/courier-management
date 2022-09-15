package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/kkjhamb01/courier-management/common/config"
)

func NewRedisConnection(config config.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.Db,
	})
}
