package business

import (
	"context"
	"errors"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/finance/db"
	"gitlab.artin.ai/backend/courier-management/finance/model"
)

func SetDefaultPaymentMethod(ctx context.Context, userId string, paymentMethodId string) error {

	logger.Infof("SetDefaultPaymentMethod userId = %v, paymentMethodId = %v", userId, paymentMethodId)

	paymentMethods, err := GetCustomerPaymentMethods(ctx, userId)
	if err != nil {
		logger.Error("failed to find customer payment methods", err)
		return err
	}

	found := false
	for _, pm := range paymentMethods {
		if pm.Id == paymentMethodId {
			found = true
			break
		}
	}

	if !found {
		err = errors.New("the payment method does not belong to the user")
		logger.Error("failed to find payment method", err)
	}

	var account model.Account
	err = db.MariaDbClient().Joins("left join account_roles on account_roles.account_id = accounts.id").
		First(&account, "account_roles.user_id = ?", userId).Error
	if err != nil {
		logger.Error("failed to fetch user's account id", err)
		return err
	}

	account.DefaultPaymentMethodId = paymentMethodId
	db.MariaDbClient().Save(&account)

	return nil
}
