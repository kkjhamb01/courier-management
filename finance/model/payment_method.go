package model

import (
	"github.com/kkjhamb01/courier-management/common/logger"
	financePb "github.com/kkjhamb01/courier-management/grpc/finance/go"
)

type PaymentMethod struct {
	AccountId string
	Type      string
	Card      Card `gorm:"foreignKey:PaymentMethodId"`
	Base
}

type Card struct {
	PaymentMethodId   string
	Brand             string
	Checks            string
	Country           string
	ExpMonth          string
	ExpYear           string
	Fingerprint       string
	Funding           string
	Last4             string
	Networks          string
	ThreeDSecureUsage string
	Wallet            string
	Base
}

func (p PaymentMethod) ToProto() (financePb.PaymentMethod, error) {
	proto := financePb.PaymentMethod{
		Id:   p.ID,
		Type: p.Type,
	}

	cardP, err := p.Card.ToProtoP()
	if err != nil {
		logger.Error("failed to convert card to PaymentMethod", err)
		return financePb.PaymentMethod{}, err
	}
	proto.Card = cardP

	return proto, err
}

func (p PaymentMethod) ToProtoP() (*financePb.PaymentMethod, error) {
	proto, err := p.ToProto()
	return &proto, err
}

func (c Card) ToProto() (financePb.PaymentMethod_Card, error) {
	proto := financePb.PaymentMethod_Card{
		Brand:             c.Brand,
		Checks:            c.Checks,
		Country:           c.Country,
		ExpMonth:          c.ExpMonth,
		ExpYear:           c.ExpYear,
		Fingerprint:       c.Fingerprint,
		Funding:           c.Funding,
		Last4:             c.Last4,
		Networks:          c.Networks,
		ThreeDSecureUsage: c.ThreeDSecureUsage,
		Wallet:            c.Wallet,
	}

	return proto, nil
}

func (c Card) ToProtoP() (*financePb.PaymentMethod_Card, error) {
	proto, err := c.ToProto()
	return &proto, err
}
