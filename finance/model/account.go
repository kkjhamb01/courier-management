package model

import (
	"errors"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	financePb "gitlab.artin.ai/backend/courier-management/grpc/finance/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountStatus string
type AccountType string

const (
	AccountStatusOpen    AccountStatus = "OPEN"
	AccountStatusBlocked AccountStatus = "BLOCKED"
	AccountStatusClosed  AccountStatus = "CLOSED"

	AccountTypeArtin    AccountType = "ARTIN"
	AccountTypeTax      AccountType = "TAX"
	AccountTypeRevenue  AccountType = "REVENUE"
	AccountTypeWallet   AccountType = "WALLET"
	AccountTypeCustomer AccountType = "CUSTOMER"
	AccountTypeCourier  AccountType = "COURIER"
)

type Account struct {
	PaymentMethods         []PaymentMethod `gorm:"foreignKey:AccountId"`
	DefaultPaymentMethodId string
	Balance                int64
	Roles                  []AccountRole `gorm:"foreignKey:AccountId"`
	Status                 AccountStatus `json:"status" sql:"type:ENUM('OPEN', 'BLOCKED', 'CLOSED')"`
	Type                   AccountType   `json:"type" sql:"type:ENUM('TAX', 'REVENUE', 'WALLET', 'CUSTOMER', 'COURIER')"`
	Base
}

func (a Account) ToProto() (financePb.Account, error) {
	accountRoles := make([]*financePb.AccountRole, len(a.Roles))
	for i, role := range a.Roles {
		var err error
		accountRoles[i], err = role.ToProtoP()
		if err != nil {
			logger.Error("failed to convert account to account proto, due to an error in converting account role", err)
			return financePb.Account{}, err
		}
	}

	accountStatus, ok := financePb.Account_Status_value[string(a.Status)]
	if !ok {
		err := errors.New("failed to match account status to Account_Status_value map")
		logger.Error("the account status is not valid", err)
		return financePb.Account{}, err
	}

	accountType, ok := financePb.Account_Type_value[string(a.Type)]
	if !ok {
		err := errors.New("failed to match account type to Account_Type_value map")
		logger.Error("the account type is not valid", err)
		return financePb.Account{}, err
	}

	paymentMethods := make([]*financePb.PaymentMethod, len(a.PaymentMethods))
	for i, pm := range a.PaymentMethods {
		var err error
		paymentMethods[i], err = pm.ToProtoP()
		if err != nil {
			logger.Error("failed to convert payment method to proto", err)
			return financePb.Account{}, err
		}
	}

	return financePb.Account{
		Id:             a.ID,
		AccountRoles:   accountRoles,
		PaymentMethods: paymentMethods,
		Status:         financePb.Account_Status(accountStatus),
		Type:           financePb.Account_Type(accountType),
		Balance:        a.Balance,
		CreatedAt:      timestamppb.New(a.CreatedAt),
		UpdatedAt:      timestamppb.New(a.UpdatedAt),
	}, nil
}

func (a Account) ToProtoP() (*financePb.Account, error) {
	proto, err := a.ToProto()
	return &proto, err
}
