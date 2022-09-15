package model

import (
	"errors"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	financePb "github.com/kkjhamb01/courier-management/grpc/finance/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TransactionStatus string
type TransactionPaymentMethod string

const (
	TransactionStatusSettled   TransactionStatus = "SETTLED"
	TransactionStatusScheduled TransactionStatus = "SCHEDULED"
	TransactionStatusDone      TransactionStatus = "DONE"
	TransactionStatusBlock     TransactionStatus = "BLOCK"
	TransactionStatusSysBlock  TransactionStatus = "SYS_BLOCK"
	TransactionStatusReject    TransactionStatus = "REJECT"

	TransactionPaymentMethodAcssDebit        TransactionPaymentMethod = "acss_debit"
	TransactionPaymentMethodAfterPayClearPay TransactionPaymentMethod = "afterpay_clearpay"
	TransactionPaymentMethodAlipay           TransactionPaymentMethod = "alipay"
	TransactionPaymentMethodAuBecsDebit      TransactionPaymentMethod = "au_becs_debit"
	TransactionPaymentMethodBacsDebit        TransactionPaymentMethod = "bacs_debit"
	TransactionPaymentMethodBanContact       TransactionPaymentMethod = "bancontact"
	TransactionPaymentMethodCard             TransactionPaymentMethod = "card"
	TransactionPaymentMethodCardPresent      TransactionPaymentMethod = "card_present"
	TransactionPaymentMethodEps              TransactionPaymentMethod = "eps"
	TransactionPaymentMethodFpx              TransactionPaymentMethod = "fpx"
	TransactionPaymentMethodGiropay          TransactionPaymentMethod = "giropay"
	TransactionPaymentMethodGrabpay          TransactionPaymentMethod = "grabpay"
	TransactionPaymentMethodIdeal            TransactionPaymentMethod = "ideal"
	TransactionPaymentMethodInteracPresent   TransactionPaymentMethod = "interac_present"
	TransactionPaymentMethodOxxo             TransactionPaymentMethod = "oxxo"
	TransactionPaymentMethodP24              TransactionPaymentMethod = "p24"
	TransactionPaymentMethodSepaDebit        TransactionPaymentMethod = "sepa_debit"
	TransactionPaymentMethodSofort           TransactionPaymentMethod = "sofort"
	TransactionPaymentMethodOther            TransactionPaymentMethod = "other"
)

type Transaction struct {
	Amount          int64 // in cent
	Currency        string
	Status          TransactionStatus        `json:"status" sql:"type:ENUM('DONE', 'BLOCK', 'SYS_BLOCK', 'REJECT')"`
	PaymentMethod   TransactionPaymentMethod `json:"payment_method" sql:"type:ENUM('acss_debit', 'afterpay_clearpay', 'alipay', 'au_becs_debit', 'bacs_debit', 'bancontact', 'card', 'card_present', 'eps', 'fpx', 'giropay', 'grabpay', 'ideal', 'interac_present', 'oxxo', 'p24', 'sepa_debit', 'sofort', 'other')"`
	SourceAccountId string
	SourceAccount   Account `gorm:"foreignKey:SourceAccountId"`
	TargetAccountId string
	TargetAccount   Account `gorm:"foreignKey:TargetAccountId"`
	TargetBalance   int64
	Description     string
	RequestId       string
	Base
}

func (t Transaction) ToProto() (financePb.Transaction, error) {
	txStatus, ok := financePb.Transaction_Status_value[string(t.Status)]
	if !ok {
		err := errors.New("failed to match transaction status to Transaction_Status_value map")
		logger.Error("the transaction status is not valid", err)
		return financePb.Transaction{}, err
	}

	txPaymentMethod, ok := financePb.Transaction_PaymentMethod_value[string(t.PaymentMethod)]
	if !ok {
		err := errors.New("failed to match transaction payment method to Transaction_PaymentMethod_value map")
		logger.Error("the transaction payment method value is not valid", err, tag.Obj("valid payment methods", financePb.Transaction_PaymentMethod_value))
		return financePb.Transaction{}, err
	}

	sourceAccount, err := t.SourceAccount.ToProtoP()
	if err != nil {
		logger.Error("failed to convert source account  to proto", err)
		return financePb.Transaction{}, err
	}

	targetAccount, err := t.TargetAccount.ToProtoP()
	if err != nil {
		logger.Error("failed to convert source account  to proto", err)
		return financePb.Transaction{}, err
	}

	return financePb.Transaction{
		Id:            t.ID,
		Amount:        t.Amount,
		Currency:      t.Currency,
		Status:        financePb.Transaction_Status(txStatus),
		PaymentMethod: financePb.Transaction_PaymentMethod(txPaymentMethod),
		SourceAccount: sourceAccount,
		TargetAccount: targetAccount,
		TargetBalance: t.TargetBalance,
		Description:   t.Description,
		RequestId:     t.RequestId,
		CreatedAt:     timestamppb.New(t.CreatedAt),
		UpdatedAt:     timestamppb.New(t.UpdatedAt),
	}, nil
}

func (t Transaction) ToProtoP() (*financePb.Transaction, error) {
	proto, err := t.ToProto()
	return &proto, err
}
