package api

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/finance/business"
	financePb "github.com/kkjhamb01/courier-management/grpc/finance/go"
	"github.com/kkjhamb01/courier-management/uaa/security"
)

func (s serverImpl) GetCustomerPaymentMethods(ctx context.Context, req *financePb.GetCustomerPaymentMethodsRequest) (*financePb.GetCustomerPaymentMethodsResponse, error) {

	logger.Infof("GetCustomerPaymentMethods request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("GetCustomerPaymentMethods request validation failed", err, tag.Obj("request", req))
		return nil, err
	}

	if ctx.Value(CtxKeyUser) == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	tokenUser := ctx.Value(CtxKeyUser).(security.User)

	paymentMethods, err := business.GetCustomerPaymentMethods(ctx, tokenUser.Id)
	if err != nil {
		logger.Error("failed to get customer payment Methods", err)
		return nil, err
	}

	logger.Info("GetCustomerPaymentMethods successfully called", tag.Obj("request", req))
	return &financePb.GetCustomerPaymentMethodsResponse{
		PaymentMethod: paymentMethods,
	}, nil
}

func (s serverImpl) DeleteCustomerPaymentMethod(ctx context.Context, req *financePb.DeleteCustomerPaymentMethodRequest) (*empty.Empty, error) {

	logger.Infof("DeleteCustomerPaymentMethod request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("DeleteCustomerPaymentMethod request validation failed", err, tag.Obj("request", req))
		return nil, err
	}

	if ctx.Value(CtxKeyUser) == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	tokenUser := ctx.Value(CtxKeyUser).(security.User)

	err := business.DeleteCustomerPaymentMethod(ctx, tokenUser.Id, req.PaymentMethodId)
	if err != nil {
		logger.Error("failed to get customer saved cards", err)
		return nil, err
	}

	logger.Info("DeleteCustomerPaymentMethod successfully called", tag.Obj("request", req))
	return &empty.Empty{}, nil
}

func (s serverImpl) SetDefaultPaymentMethod(ctx context.Context, req *financePb.SetDefaultPaymentMethodRequest) (*empty.Empty, error) {

	logger.Infof("SetDefaultPaymentMethod request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("SetDefaultPaymentMethod request validation failed", err, tag.Obj("request", req))
		return nil, err
	}

	if ctx.Value(CtxKeyUser) == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	tokenUser := ctx.Value(CtxKeyUser).(security.User)

	err := business.SetDefaultPaymentMethod(ctx, tokenUser.Id, req.PaymentMethodId)
	if err != nil {
		logger.Error("failed to set default payment method", err)
		return nil, err
	}

	logger.Info("SetDefaultPaymentMethod successfully called", tag.Obj("request", req))
	return &empty.Empty{}, nil
}
