package business

import (
	"context"
	"fmt"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentmethod"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/finance/storage"
	financePb "gitlab.artin.ai/backend/courier-management/grpc/finance/go"
)

func GetCustomerPaymentMethods(ctx context.Context, userId string) ([]*financePb.PaymentMethod, error) {
	customerId, err := storage.GetStripeCustomerId(ctx, userId)

	logger.Infof("GetCustomerPaymentMethods userId = %v, customerId = %v", userId, customerId)

	if err != nil {
		logger.Errorf("failed to load customer id for user", err, userId)
		return nil, err
	}

	cards := make([]*financePb.PaymentMethod, 0, 5)

	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(customerId),
		Type:     stripe.String(string(stripe.PaymentMethodTypeCard)),
	}

	i := paymentmethod.List(params)
	err = i.Err()
	if err != nil {
		logger.Error("failed to get list of payment methods", err)
		return nil, err
	}

	for i.Next() {
		pm := i.PaymentMethod()

		// TODO return valuable attributes from PaymentMethod object
		cards = append(cards, &financePb.PaymentMethod{
			Id:   pm.ID,
			Type: string(pm.Type),
			Card: &financePb.PaymentMethod_Card{
				Brand:       string(pm.Card.Brand),
				Country:     pm.Card.Country,
				ExpMonth:    fmt.Sprintf("%v", pm.Card.ExpMonth),
				ExpYear:     fmt.Sprintf("%v", pm.Card.ExpYear),
				Fingerprint: pm.Card.Fingerprint,
				Funding:     string(pm.Card.Funding),
				Last4:       pm.Card.Last4,
			},
		})
	}

	return cards, nil
}
