package model

import (
	"errors"
	"time"

	"github.com/kkjhamb01/courier-management/common/logger"
	financePb "github.com/kkjhamb01/courier-management/grpc/finance/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountRoleStatus string
type AccountRoleType string

const (
	AccountRoleStatusActive  AccountRoleStatus = "ACTIVE"
	AccountRoleStatusDeleted AccountRoleStatus = "DELETED"
	AccountRoleStatusBlocked AccountRoleStatus = "BLOCKED"

	AccountRoleTypeOwner  AccountRoleType = "OWNER"
	AccountRoleTypeClient AccountRoleType = "CLIENT"
	AccountRoleTypeAdmin  AccountRoleType = "ADMIN"
)

type AccountRole struct {
	UserId    string `gorm:"type:BINARY(36)"`
	StripeId  string
	FromDate  time.Time
	ToDate    *time.Time
	AccountId string
	Status    AccountRoleStatus `json:"status" sql:"type:ENUM('ACTIVE', 'DELETED', 'BLOCKED')"`
	Type      AccountRoleType   `json:"type" sql:"type:ENUM('OWNER', 'CLIENT', 'ADMIN')"`
	Base
}

func (a AccountRole) ToProto() (financePb.AccountRole, error) {
	roleStatus, ok := financePb.AccountRole_Status_value[string(a.Status)]
	if !ok {
		err := errors.New("failed to match account role status to AccountRole_Status_value map")
		logger.Error("the account role status is not valid", err)
		return financePb.AccountRole{}, err
	}

	roleType, ok := financePb.AccountRole_Type_value[string(a.Type)]
	if !ok {
		err := errors.New("failed to match account role type to AccountRole_Type_value map")
		logger.Error("the account role type is not valid", err)
		return financePb.AccountRole{}, err
	}

	var roleToDate *timestamppb.Timestamp
	if a.ToDate != nil {
		roleToDate = timestamppb.New(*a.ToDate)
	}

	return financePb.AccountRole{
		Id:        a.ID,
		UserId:    a.UserId,
		FromDate:  timestamppb.New(a.FromDate),
		ToDate:    roleToDate,
		Status:    financePb.AccountRole_Status(roleStatus),
		Type:      financePb.AccountRole_Type(roleType),
		CreatedAt: timestamppb.New(a.CreatedAt),
		UpdatedAt: timestamppb.New(a.UpdatedAt),
	}, nil
}

func (a AccountRole) ToProtoP() (*financePb.AccountRole, error) {
	proto, err := a.ToProto()
	return &proto, err
}
