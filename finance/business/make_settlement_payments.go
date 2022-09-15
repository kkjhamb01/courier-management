package business

import (
	"context"
	"errors"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/transfer"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/finance/db"
	"gitlab.artin.ai/backend/courier-management/finance/model"
	"gitlab.artin.ai/backend/courier-management/finance/storage"
	"time"
)

func MakeSettlementPayments() {

	//logger.Infof("MakeSettlementPayments")

	var err error

	// fetch transactions need to be actually paid
	var transactions []model.Transaction
	err = db.MariaDbClient().Where("status = ?", model.TransactionStatusScheduled).
		Preload("SourceAccount").
		Preload("TargetAccount").
		Preload("SourceAccount.Roles").
		Preload("TargetAccount.Roles").
		Find(&transactions).
		Error
	if err != nil {
		logger.Error("failed to fetch scheduled transactions", err)
		return
	}

	type amountAndCurrency struct {
		amount   int64
		currency string
	}

	targetStripeAccountToAmountMap := make(map[string]*amountAndCurrency, 0)
	targetStripeAccountToTransactionIdMap := make(map[string][]string, 0)
	targetStripeAccountToDBAccountId := make(map[string]string, 0)

	// create a maps
	for _, transaction := range transactions {
		var destinationStripeAccountId string

		switch transaction.TargetAccount.Type {
		case model.AccountTypeTax:
			destinationStripeAccountId = config.Finance().TaxStripeId
		case model.AccountTypeRevenue:
			destinationStripeAccountId = config.Finance().RevenueStripeId
		case model.AccountTypeCourier:
			if len(transaction.TargetAccount.Roles) == 0 {
				err = errors.New("target account roles is empty")
				return
			}
			destinationAccountUserId := transaction.TargetAccount.Roles[0].UserId
			destinationStripeAccountId, err = storage.GetStripeAccId(context.Background(), destinationAccountUserId)
			if err != nil {
				return
			}
		}

		otherAmounts, ok := targetStripeAccountToAmountMap[destinationStripeAccountId]
		if !ok {
			targetStripeAccountToAmountMap[destinationStripeAccountId] = &amountAndCurrency{
				amount:   transaction.Amount,
				currency: transaction.Currency,
			}
		} else {
			targetStripeAccountToAmountMap[destinationStripeAccountId].amount += otherAmounts.amount
		}

		_, ok = targetStripeAccountToTransactionIdMap[transaction.ID]
		if !ok {
			targetStripeAccountToTransactionIdMap[destinationStripeAccountId] = []string{
				transaction.ID,
			}
		} else {
			targetStripeAccountToTransactionIdMap[destinationStripeAccountId] = append(targetStripeAccountToTransactionIdMap[destinationStripeAccountId], transaction.ID)
		}

		targetStripeAccountToDBAccountId[destinationStripeAccountId] = transaction.TargetAccount.ID
	}

	artinSourceAccount := model.Account{}
	err = db.MariaDbClient().
		Where("type = ?", model.AccountTypeArtin).
		First(&artinSourceAccount).
		Error
	if err != nil {
		logger.Error("failed to find artin source account", err)
		return
	}

	// transfer the money and change the status of transactions related to the target to DONE
	for targetStripeAccountId, targetAmountAndCurrency := range targetStripeAccountToAmountMap {
		transferParams := &stripe.TransferParams{
			Amount:      stripe.Int64(targetAmountAndCurrency.amount),
			Currency:    stripe.String(targetAmountAndCurrency.currency),
			Destination: stripe.String(targetStripeAccountId),
		}
		_, err = transfer.New(transferParams)
		if err != nil {
			logger.Error("failed to create transfer to the target account", err, tag.Str("target account", targetStripeAccountId))
			return
		}

		for _, transactionId := range targetStripeAccountToTransactionIdMap[targetStripeAccountId] {
			err = db.MariaDbClient().Table("transactions").
				Where("ID = ?", transactionId).
				Update("status", model.TransactionStatusDone).Error
			if err != nil {
				return
			}
		}

		settlementTransaction := model.Transaction{
			Amount:          targetAmountAndCurrency.amount * -1,
			Currency:        targetAmountAndCurrency.currency,
			Status:          model.TransactionStatusSettled,
			PaymentMethod:   "other",
			SourceAccountId: artinSourceAccount.ID,
			TargetAccountId: targetStripeAccountToDBAccountId[targetStripeAccountId],
			TargetBalance:   0,
			Description:     "Settlement run on " + time.Now().Format("2006-01-02 15:04:05"),
		}

		err = db.MariaDbClient().
			Create(&settlementTransaction).
			Error
		if err != nil {
			return
		}
	}
}
