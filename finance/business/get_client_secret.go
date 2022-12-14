package business

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/finance/storage"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/setupintent"
)

func GetClientSecret(ctx context.Context, userId string) (string, error) {
	customerId, err := storage.GetStripeCustomerId(ctx, userId)

	logger.Infof("GetClientSecret userId = %v, customerId = %v", userId, customerId)

	if err != nil {
		logger.Errorf("failed to load customer id for user", err, userId)
		return "", err
	}

	params := &stripe.SetupIntentParams{
		Customer: stripe.String(customerId),
	}
	si, err := setupintent.New(params)
	if err != nil {
		logger.Error("failed to create a setup intent", err)
		return "", err
	}

	return si.ClientSecret, nil
}
