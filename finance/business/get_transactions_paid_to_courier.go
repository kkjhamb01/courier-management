package business

import (
	"context"
	"time"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/finance/db"
	"github.com/kkjhamb01/courier-management/finance/model"
	financePb "github.com/kkjhamb01/courier-management/grpc/finance/go"
)

func GetTransactionsPaidToCourier(ctx context.Context, courierId string, from time.Time, to time.Time, pageNumber int, pageSize int) ([]*financePb.Transaction, error) {

	logger.Infof("GetTransactionsPaidToCourier courierId = %v, from = %v, to = %v, pageNumber = %v, pageSize = %v", courierId, from, to, pageNumber, pageSize)

	offset := pageSize * (pageNumber - 1)

	var transactions []model.Transaction
	err := db.MariaDbClient().
		Preload("SourceAccount").
		Preload("TargetAccount").
		Joins("join accounts on accounts.id = target_account_id").
		Joins("join account_roles on accounts.id = account_roles.account_id").
		Where("transactions.status = ? OR transactions.status = ? OR transactions.status = ?", model.TransactionStatusDone, model.TransactionStatusScheduled, model.TransactionStatusSettled).
		Where("accounts.type = ?", model.AccountTypeCourier).
		Where("account_roles.user_id = ?", courierId).
		Where("transactions.updated_at >= ?", from).
		Where("transactions.updated_at <= ?", to).
		Order("transactions.updated_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&transactions).
		Error
	if err != nil {
		logger.Error("failed to fetch transactions from DB", err)
		return nil, err
	}

	transactionsProto := make([]*financePb.Transaction, len(transactions))
	for i, transaction := range transactions {
		transactionsProto[i], err = transaction.ToProtoP()
		if err != nil {
			logger.Error("failed to convert model transaction to proto transaction", err)
			return nil, err
		}
	}

	return transactionsProto, nil
}
