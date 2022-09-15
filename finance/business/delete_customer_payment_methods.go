package business

import (
	"context"
	"fmt"
	"github.com/stripe/stripe-go/v72/paymentmethod"
	"gitlab.artin.ai/backend/courier-management/common/logger"
)

func DeleteCustomerPaymentMethod(ctx context.Context, userId string, paymentMethodId string) error {
	// TODO do we need to add it to our history ?

	logger.Infof("DeleteCustomerPaymentMethod userId = %v, paymentMethodId = %v", userId, paymentMethodId)

	paymentMethods, err := GetCustomerPaymentMethods(ctx, userId)
	if err != nil {
		logger.Error("failed to check the user's payment methods", err)
		return err
	}

	paymentMethodBelongsToUser := false
	for _, paymentMethod := range paymentMethods {
		if paymentMethod.Id == paymentMethodId {
			paymentMethodBelongsToUser = true
			break
		}
	}
	if !paymentMethodBelongsToUser {
		err := fmt.Errorf("the payment method %v is not belonged to the user", paymentMethodId)
		logger.Error("failed to delete the payment method", err)
		return err
	}

	_, err = paymentmethod.Detach(paymentMethodId, nil)
	if err != nil {
		logger.Error("failed to delete the payment method", err)
		return err
	}

	return nil
}
