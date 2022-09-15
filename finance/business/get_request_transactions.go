package business

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/finance/db"
	"github.com/kkjhamb01/courier-management/finance/model"
	financePb "github.com/kkjhamb01/courier-management/grpc/finance/go"
)

func GetRequestTransactions(ctx context.Context, requestId string) ([]*financePb.Transaction, error) {

	logger.Infof("GetRequestTransactions requestId = %v", requestId)

	var transactions []model.Transaction
	err := db.MariaDbClient().
		Preload("TargetAccount").
		Preload("SourceAccount").
		Where("transactions.request_id = ?", requestId).
		Find(&transactions).
		Error
	if err != nil {
		logger.Error("failed to get request transactions from DB", err)
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
