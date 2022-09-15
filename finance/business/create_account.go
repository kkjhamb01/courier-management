package business

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/finance/db"
	"github.com/kkjhamb01/courier-management/finance/model"
)

func CreateAccount(ctx context.Context, userId string, accessToken string) error {

	logger.Infof("CreateAccount userId = %v", userId)

	isCustomer := true
	if _, err := getCustomerInfo(ctx, accessToken); err != nil {
		isCustomer = false
	}
	isCourier := true
	if _, err := getCourierInfo(ctx, accessToken); err != nil {
		isCourier = false
	}
	if isCourier == isCustomer {
		err := fmt.Errorf("failed to detect user: is courier: %v, is customer: %v", isCourier, isCustomer)
		logger.Error("failed to determine if the user is either courier or customer", err, tag.Str("userId", userId))
		return err
	}

	var accountType model.AccountType
	if isCourier {
		accountType = model.AccountTypeCourier
	} else {
		accountType = model.AccountTypeCustomer
	}

	account := model.Account{
		Roles: []model.AccountRole{
			{
				UserId:   userId,
				FromDate: time.Now(),
				ToDate:   nil,
				Status:   model.AccountRoleStatusActive,
				Type:     model.AccountRoleTypeOwner,
			},
		},
		Balance: 0,
		Status:  model.AccountStatusOpen,
		Type:    accountType,
	}
	result := db.MariaDbClient().Create(&account)
	if result.Error != nil {
		logger.Error("failed to create account", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		err := errors.New("no row is affected")
		logger.Error("failed to create Account", err)
		return err
	}

	return nil
}
