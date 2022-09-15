package db

import (
	"fmt"
	goredisV7lib "github.com/go-redis/redis/v7"
	//"context"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"sync"
)

// to make sure redis db clients would be set up only once
var redisDbSetupOnce sync.Once

var redisClient *goredislib.Client
var redisV7Client *goredisV7lib.Client

func SetupRedisClient() {
	redisDbSetupOnce.Do(func() {
		redisClient = goredislib.NewClient(&goredislib.Options{
			Addr:     config.OfferRedis().Address,
			Password: config.OfferRedis().Password,
			DB:       config.OfferRedis().DB,
		})

		// TODO PING DB
		//	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		//	panic("redis is not accessible")
		//	}

		option, err := goredisV7lib.ParseURL("redis://" + config.OfferRedis().Address + "/" + fmt.Sprintf("%v", config.OfferRedis().DB))
		if err != nil {
			logger.Fatalf("failed to parse redis url, %v", err)
			panic(err)
		}
		redisV7Client = goredisV7lib.NewClient(option)

		logger.Info("created connection to redis (Not guaranteed as pinging is disabled)")
	})
}

func FinanceRedisClient() *goredislib.Client {
	return redisClient
}

func RedisV7Client() *goredisV7lib.Client {
	return redisV7Client
}

func RedisMutex(mutexName string) *redsync.Mutex {
	pool := goredis.NewPool(FinanceRedisClient())
	rs := redsync.New(pool)
	return rs.NewMutex(mutexName)
}
