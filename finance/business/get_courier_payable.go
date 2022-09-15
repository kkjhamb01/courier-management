package business

import (
	"context"
	"database/sql"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/finance/db"
	"github.com/kkjhamb01/courier-management/finance/model"
)

func GetCourierPayable(ctx context.Context, courierId string) (float64, string, error) {
	var courierPayable sql.NullInt64

	logger.Infof("GetCourierPayable courierId = %v", courierId)

	err := db.MariaDbClient().Model(&model.Transaction{}).
		Select("sum(transactions.amount)").
		Joins("join accounts on accounts.id = target_account_id").
		Joins("join account_roles on accounts.id = account_roles.account_id").
		Where("transactions.status = ?", model.TransactionStatusScheduled).
		Where("accounts.type = ?", model.AccountTypeCourier).
		Where("account_roles.user_id = ?", courierId).
		Scan(&courierPayable).
		Error
	if err != nil {
		logger.Error("failed to fetch sum of courier paybale from DB", err)
		return 0, "", err
	}

	var courierPayableFloat64 float64
	if courierPayable.Valid {
		courierPayableFloat64 = float64(courierPayable.Int64)
	}

	return courierPayableFloat64, "GBP", nil
}
