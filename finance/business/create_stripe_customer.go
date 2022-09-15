package business

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/finance/storage"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
)

func CreateStripeCustomer(ctx context.Context, userId, email, name, phone string) error {

	logger.Infof("CreateStripeCustomer userId = %v, email = %v, name = %v, phone = %v", userId, email, name, phone)

	params := &stripe.CustomerParams{
		Email: stripe.String(email),
		Name:  stripe.String(name),
		Phone: stripe.String(phone),
	}
	c, err := customer.New(params)
	if err != nil {
		logger.Error("failed to create a new customer", err)
		return err
	}

	err = storage.SetStripeCustomerId(ctx, userId, c.ID)
	if err != nil {
		logger.Error("failed to save customer ID", err)
		return err
	}

	err = storage.SetCustomerUserId(ctx, c.ID, userId)
	if err != nil {
		logger.Error("failed to save stripe ID", err)
		return err
	}

	return nil
}
