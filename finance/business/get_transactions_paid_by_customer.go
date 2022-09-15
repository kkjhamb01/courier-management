package business

import (
	"context"
	"time"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/finance/db"
	"github.com/kkjhamb01/courier-management/finance/model"
	financePb "github.com/kkjhamb01/courier-management/grpc/finance/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetTransactionsPaidByCustomer(ctx context.Context, customerId string, from time.Time, to time.Time) ([]*financePb.GetTransactionsPaidByCustomerResponse_Item, error) {

	logger.Infof("GetTransactionsPaidByCustomer customerId = %v, from = %v, to = %v", customerId, from, to)

	var transactions []model.Transaction
	err := db.MariaDbClient().
		Joins("join accounts on accounts.id = source_account_id").
		Joins("join account_roles on accounts.id = account_roles.account_id").
		Where("transactions.status = ? OR transactions.status = ?", model.TransactionStatusScheduled, model.TransactionStatusDone).
		Where("accounts.type = ?", model.AccountTypeCustomer).
		Where("account_roles.user_id = ?", customerId).
		Where("transactions.created_at >= ?", from).
		Where("transactions.created_at <= ?", to).
		Order("transactions.created_at DESC").
		Find(&transactions).
		Error
	if err != nil {
		logger.Error("failed to fetch transactions from DB", err)
		return nil, err
	}

	requestIdMap := make(map[string]*financePb.GetTransactionsPaidByCustomerResponse_Item, 0)
	response := make([]*financePb.GetTransactionsPaidByCustomerResponse_Item, 0)

	for _, transaction := range transactions {
		otherTransactions, ok := requestIdMap[transaction.RequestId]
		if !ok {
			requestIdMap[transaction.RequestId] = &financePb.GetTransactionsPaidByCustomerResponse_Item{
				RequestId: transaction.RequestId,
				Amount:    float64(transaction.Amount),
				Currency:  transaction.Currency,
				CreatedAt: timestamppb.New(transaction.CreatedAt),
			}
			response = append(response, requestIdMap[transaction.RequestId])
		} else {
			otherTransactions.Amount += float64(transaction.Amount)
		}
	}

	return response, nil
}
