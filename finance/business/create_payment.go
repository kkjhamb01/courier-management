package business

import (
	"context"
	"fmt"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/finance/db"
	"github.com/kkjhamb01/courier-management/finance/model"
	"github.com/kkjhamb01/courier-management/finance/storage"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/v72/paymentmethod"
	"gorm.io/gorm"
)

func CreatePayment(ctx context.Context, amountInCent int64, currency string, customerId string, courierId string, paymentMethodId string, requestId string) (clientSecret string, err error) {

	logger.Infof("CreatePayment courierId = %v, customerId = %v, amountInCent = %v, currency = %v, paymentMethodId = %v, requestId = %v", courierId, customerId, amountInCent, currency, paymentMethodId, requestId)

	customerStripeId, err := storage.GetStripeCustomerId(ctx, customerId)
	if err != nil {
		logger.Error("failed to get customer stripe ID", err)
		return "", err
	}

	paymentMethod, err := paymentmethod.Get(
		paymentMethodId,
		nil,
	)
	if err != nil {
		logger.Error("failed to get payment method", err, tag.Str("payment method id", paymentMethodId))
		return "", err
	}

	// Create a PaymentIntent:
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amountInCent),
		Currency: stripe.String(currency),
		PaymentMethodTypes: stripe.StringSlice([]string{
			string(paymentMethod.Type),
		}),
		PaymentMethod: stripe.String(paymentMethodId),
		Customer:      stripe.String(customerStripeId),
		Confirm:       stripe.Bool(true),
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		logger.Error("failed to new payment intent", err)
		return "", err
	}
	print(pi)

	courierShareInCent := amountInCent * int64(config.Finance().PaymentCourierSharePercent) / 100
	revenueShareInCent := amountInCent * int64(config.Finance().PaymentRevenuePercent) / 100
	taxShareInCent := amountInCent * int64(config.Finance().PaymentTaxSharePercent) / 100

	var customerAccountModel model.Account
	err = db.MariaDbClient().Joins("left join account_roles on account_roles.account_id = accounts.id").
		First(&customerAccountModel, "account_roles.user_id = ?", customerId).Error
	if err != nil {
		logger.Error("failed to fetch customer's account", err)
		return "", err
	}

	var courierAccountModel model.Account
	err = db.MariaDbClient().Joins("left join account_roles on account_roles.account_id = accounts.id").
		First(&courierAccountModel, "account_roles.user_id = ?", courierId).Error
	if err != nil {
		logger.Error("failed to fetch courier's account", err)
		return "", err
	}

	courierTransactionModel := model.Transaction{
		Amount:          courierShareInCent,
		Currency:        currency,
		Status:          model.TransactionStatusScheduled,
		PaymentMethod:   model.TransactionPaymentMethod(paymentMethod.Type),
		SourceAccountId: customerAccountModel.ID,
		TargetAccountId: courierAccountModel.ID,
		TargetBalance:   courierAccountModel.Balance + courierShareInCent,
		RequestId:       requestId,
		Description:     fmt.Sprintf("paying courier, from customer (id: %v) to courier (id: %v)", customerId, courierId),
	}

	var revenueAccount model.Account
	err = db.MariaDbClient().Model(&revenueAccount).
		Where("type = ?", model.AccountTypeRevenue).
		First(&revenueAccount).Error
	if err != nil {
		logger.Error("failed to get revenue account", err)
		return "", err
	}
	revenueTransactionModel := model.Transaction{
		Amount:          revenueShareInCent,
		Currency:        currency,
		Status:          model.TransactionStatusScheduled,
		PaymentMethod:   model.TransactionPaymentMethod(paymentMethod.Type),
		SourceAccountId: customerAccountModel.ID,
		TargetAccountId: revenueAccount.ID,
		TargetBalance:   revenueAccount.Balance + revenueShareInCent,
		RequestId:       requestId,
		Description:     fmt.Sprintf("paying revenue, from customer (id: %v) to revenue", customerId),
	}

	var taxAccount model.Account
	err = db.MariaDbClient().Model(&taxAccount).
		Where("type = ?", model.AccountTypeTax).
		First(&taxAccount).Error
	if err != nil {
		logger.Error("failed to get tax account", err)
		return "", err
	}
	taxTransactionModel := model.Transaction{
		Amount:          taxShareInCent,
		Currency:        currency,
		Status:          model.TransactionStatusScheduled,
		PaymentMethod:   model.TransactionPaymentMethod(paymentMethod.Type),
		SourceAccountId: customerAccountModel.ID,
		TargetAccountId: taxAccount.ID,
		TargetBalance:   taxAccount.Balance + taxShareInCent,
		RequestId:       requestId,
		Description:     fmt.Sprintf("paying tax, from customer (id: %v) to tax", customerId),
	}

	err = db.MariaDbClient().Transaction(func(tx *gorm.DB) error {
		err = tx.Create(&courierTransactionModel).Error
		if err != nil {
			logger.Error("failed to persist create courier transaction", err)
			return err
		}

		err = tx.Create(&revenueTransactionModel).Error
		if err != nil {
			logger.Error("failed to persist create revenue transaction", err)
			return err
		}

		err = tx.Create(&taxTransactionModel).Error
		if err != nil {
			logger.Error("failed to persist create tax transaction", err)
			return err
		}

		tx.Model(&taxAccount).Update("balance", taxAccount.Balance+taxShareInCent)
		tx.Model(&revenueAccount).Update("balance", revenueAccount.Balance+revenueShareInCent)
		tx.Model(&courierAccountModel).Update("balance", courierAccountModel.Balance+courierShareInCent)

		return nil
	})

	if err != nil {
		logger.Error("failed to persist transaction logs", err)
		return "", err
	}

	return pi.ClientSecret, nil
}
