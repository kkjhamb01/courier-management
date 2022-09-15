package api

import (
	"context"
	"errors"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/finance/business"
	financePb "github.com/kkjhamb01/courier-management/grpc/finance/go"
	"github.com/kkjhamb01/courier-management/uaa/security"
)

func (s serverImpl) CreatePayment(ctx context.Context, req *financePb.CreatePaymentRequest) (*financePb.CreatePaymentResponse, error) {

	logger.Infof("CreatePayment request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("CreatePayment request validation failed", err, tag.Obj("request", req))
		return nil, err
	}

	if ctx.Value(CtxKeyUser) == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	tokenUser := ctx.Value(CtxKeyUser).(security.User)

	clientSecret, err := business.CreatePayment(ctx, int64(req.Price*100), req.Currency, tokenUser.Id, req.CourierId, req.PaymentMethodId, req.RequestId)
	if err != nil {
		logger.Error("failed to get customer saved cards", err)
		return nil, err
	}

	logger.Info("CreatePayment successfully called", tag.Obj("request", req))
	return &financePb.CreatePaymentResponse{
		ClientSecret: clientSecret,
	}, nil
}
