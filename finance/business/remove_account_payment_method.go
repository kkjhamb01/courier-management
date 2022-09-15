package business

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/finance/db"
	"github.com/kkjhamb01/courier-management/finance/model"
)

func RemovePaymentMethodFromAccount(ctx context.Context, paymentMethodId string) error {

	logger.Infof("RemovePaymentMethodFromAccount paymentMethodId = %v", paymentMethodId)

	pm := model.PaymentMethod{}
	pm.ID = paymentMethodId

	err := db.MariaDbClient().
		Delete(&pm).Error
	if err != nil {
		logger.Error("failed to delete payment method", err)
		return err
	}

	return err
}
