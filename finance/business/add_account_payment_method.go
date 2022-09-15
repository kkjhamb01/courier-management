package business

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/finance/db"
	"github.com/kkjhamb01/courier-management/finance/model"
	financePb "github.com/kkjhamb01/courier-management/grpc/finance/go"
)

func AddPaymentMethodToAccount(ctx context.Context, userId string, paymentMethod financePb.PaymentMethod) error {

	logger.Infof("AddPaymentMethodToAccount userId = %v, paymentMethod = %v", userId, paymentMethod)

	cardModel := model.Card{
		Brand:             paymentMethod.Card.Brand,
		Checks:            paymentMethod.Card.Checks,
		Country:           paymentMethod.Card.Country,
		ExpMonth:          paymentMethod.Card.ExpMonth,
		ExpYear:           paymentMethod.Card.ExpYear,
		Fingerprint:       paymentMethod.Card.Fingerprint,
		Funding:           paymentMethod.Card.Funding,
		Last4:             paymentMethod.Card.Last4,
		Networks:          paymentMethod.Card.Networks,
		ThreeDSecureUsage: paymentMethod.Card.ThreeDSecureUsage,
		Wallet:            paymentMethod.Card.Wallet,
	}

	paymentMethodModel := model.PaymentMethod{
		Type: paymentMethod.Type,
		Card: cardModel,
	}

	err := db.MariaDbClient().
		Model(&model.Account{}).
		Where("id = ?", userId).
		Update("payment_methods", paymentMethodModel).Error
	if err != nil {
		logger.Error("failed to update account", err)
		return err
	}

	return err
}
