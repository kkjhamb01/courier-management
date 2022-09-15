package storage

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/finance/db"
)

const (
	accPrefix          = "acc_"
	customerPrefix     = "customer_"
	customerUserPrefix = "customer_user_"
	accUserPrefix      = "acc_user_"
)

func GetStripeAccId(ctx context.Context, userId string) (string, error) {
	accId, err := db.FinanceRedisClient().Get(ctx, accPrefix+userId).Result()
	if err != nil {
		logger.Error("failed to get from redis", err)
		return "", err
	}

	return accId, nil
}

func SetStripeAccId(ctx context.Context, userId, accId string) error {
	err := db.FinanceRedisClient().Set(ctx, accPrefix+userId, accId, -1).Err()
	if err != nil {
		logger.Error("failed to set into redis", err)
		return err
	}

	return nil
}

func GetStripeCustomerId(ctx context.Context, userId string) (string, error) {
	stripeId, err := db.FinanceRedisClient().Get(ctx, customerPrefix+userId).Result()
	if err != nil {
		logger.Error("failed to get from redis", err)
		return "", err
	}

	return stripeId, nil
}

func SetStripeCustomerId(ctx context.Context, userId, stripeId string) error {
	err := db.FinanceRedisClient().Set(ctx, customerPrefix+userId, stripeId, -1).Err()
	if err != nil {
		logger.Error("failed to set into redis", err)
		return err
	}

	return nil
}

func GetCustomerUserId(ctx context.Context, stripeId string) (string, error) {
	stripeId, err := db.FinanceRedisClient().Get(ctx, customerUserPrefix+stripeId).Result()
	if err != nil {
		logger.Error("failed to get from redis", err)
		return "", err
	}

	return stripeId, nil
}

func SetCustomerUserId(ctx context.Context, stripeId, userId string) error {
	err := db.FinanceRedisClient().Set(ctx, customerUserPrefix+stripeId, userId, -1).Err()
	if err != nil {
		logger.Error("failed to set into redis", err)
		return err
	}

	return nil
}

func GetAccUserId(ctx context.Context, accId string) (string, error) {
	accId, err := db.FinanceRedisClient().Get(ctx, accUserPrefix+accId).Result()
	if err != nil {
		logger.Error("failed to get from redis", err)
		return "", err
	}

	return accId, nil
}

func SetAccUserId(ctx context.Context, accId, userId string) error {
	err := db.FinanceRedisClient().Set(ctx, accUserPrefix+accId, userId, -1).Err()
	if err != nil {
		logger.Error("failed to set into redis", err)
		return err
	}

	return nil
}
