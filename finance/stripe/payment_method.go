package stripe

import (
	"context"
	"encoding/json"
	"github.com/stripe/stripe-go/v72"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/finance/business"
	"gitlab.artin.ai/backend/courier-management/finance/storage"
	financePb "gitlab.artin.ai/backend/courier-management/grpc/finance/go"
)

func onPaymentMethodDetached(ctx context.Context, event stripe.Event) error {

	logger.Infof("onPaymentMethodDetached event = %+v", event)

	var paymentMethod stripe.PaymentMethod
	err := json.Unmarshal(event.Data.Raw, &paymentMethod)
	if err != nil {
		return err
	}

	err = business.RemovePaymentMethodFromAccount(ctx, paymentMethod.ID)
	if err != nil {
		logger.Error("failed to remove payment method", err)
		return err
	}

	return nil
}

func onPaymentMethodAttached(ctx context.Context, event stripe.Event) error {

	logger.Infof("onPaymentMethodAttached event = %+v", event)

	var paymentMethod stripe.PaymentMethod
	err := json.Unmarshal(event.Data.Raw, &paymentMethod)
	if err != nil {
		return err
	}

	userId, err := storage.GetCustomerUserId(ctx, paymentMethod.Customer.ID)
	if err != nil {
		logger.Error("failed to load user Id", err)
		return err
	}

	err = business.AddPaymentMethodToAccount(ctx, userId, financePb.PaymentMethod{
		Id:   paymentMethod.ID,
		Type: string(paymentMethod.Type),
		Card: &financePb.PaymentMethod_Card{
			Brand:       string(paymentMethod.Card.Brand),
			Country:     paymentMethod.Card.Country,
			ExpMonth:    string(paymentMethod.Card.ExpMonth),
			ExpYear:     string(paymentMethod.Card.ExpYear),
			Fingerprint: paymentMethod.Card.Fingerprint,
			Funding:     string(paymentMethod.Card.Funding),
			Last4:       paymentMethod.Card.Last4,
		},
	})
	if err != nil {
		logger.Error("failed to add payment method to account", err)
		return err
	}

	return nil
}
