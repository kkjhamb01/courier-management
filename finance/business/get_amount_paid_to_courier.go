package business

import (
	"context"
	"database/sql"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/finance/db"
	"gitlab.artin.ai/backend/courier-management/finance/model"
	"time"
)

func GetAmountPaidToCourier(ctx context.Context, courierId string, from time.Time, to time.Time) (int64, string, error) {

	logger.Infof("GetAmountPaidToCourier courierId = %v, from = %v, to = %v", courierId, from, to)

	var totalAmount sql.NullInt64
	err := db.MariaDbClient().Model(&model.Transaction{}).
		Select("SUM(transactions.amount)").
		Joins("join accounts on accounts.id = target_account_id").
		Joins("join account_roles on accounts.id = account_roles.account_id").
		Where("transactions.status = ?", model.TransactionStatusDone).
		Where("accounts.type = ?", model.AccountTypeCourier).
		Where("account_roles.user_id = ?", courierId).
		Where("transactions.created_at >= ?", from).
		Where("transactions.created_at <= ?", to).
		Scan(&totalAmount).
		Error
	if err != nil {
		logger.Error("failed to fetch transactions amount from DB", err)
		return 0, "", err
	}

	// TODO all amounts (including their currency) needs to be fetched and converted to the target currency
	// it means the client should request a currency to calculate the total amountAndCurrency based on
	var transaction model.Transaction
	err = db.MariaDbClient().Model(&model.Transaction{}).
		Select("transactions.currency").
		Joins("join accounts on accounts.id = target_account_id").
		Joins("join account_roles on accounts.id = account_roles.account_id").
		Where("transactions.status = ?", model.TransactionStatusDone).
		Where("accounts.type = ?", model.AccountTypeCourier).
		Where("account_roles.user_id = ?", courierId).
		Where("transactions.created_at >= ?", from).
		Where("transactions.created_at <= ?", to).
		Scan(&transaction).Error
	if err != nil {
		logger.Error("failed to fetch transaction from DB", err)
		return 0, "", err
	}

	var totalAmountInt64 int64
	if totalAmount.Valid {
		totalAmountInt64 = totalAmount.Int64
	}

	return totalAmountInt64, transaction.Currency, nil
}
