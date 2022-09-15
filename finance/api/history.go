package api

import (
	"context"
	"errors"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/finance/business"
	financePb "github.com/kkjhamb01/courier-management/grpc/finance/go"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s serverImpl) GetCourierPayable(ctx context.Context, req *emptypb.Empty) (*financePb.GetCourierPayableResponse, error) {

	logger.Infof("GetCourierPayable request = %+v", req)

	if ctx.Value(CtxKeyUser) == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value(CtxKeyUser).(security.User)

	amount, currency, err := business.GetCourierPayable(ctx, tokenUser.Id)
	if err != nil {
		logger.Error("failed to get courier payable", err)
		return nil, err
	}

	logger.Info("GetCustomerPayableResponse successfully called", tag.Obj("request", req))
	return &financePb.GetCourierPayableResponse{
		Price:    amount,
		Currency: currency,
	}, nil
}

func (s serverImpl) GetRequestTransactions(ctx context.Context, req *financePb.GetRequestTransactionsRequest) (*financePb.GetRequestTransactionsResponse, error) {

	logger.Infof("GetRequestTransactions request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("GetRequestTransactions request validation failed", err, tag.Obj("request", req))
		return nil, err
	}

	transactions, err := business.GetRequestTransactions(ctx, req.RequestId)
	if err != nil {
		logger.Error("failed to get request transactions", err)
		return nil, err
	}

	logger.Info("GetRequestTransactionsResponse successfully called", tag.Obj("request", req))
	return &financePb.GetRequestTransactionsResponse{
		Transactions: transactions,
	}, nil
}

func (s serverImpl) GetTransactionsPaidByCustomer(ctx context.Context, req *financePb.GetTransactionsPaidByCustomerRequest) (*financePb.GetTransactionsPaidByCustomerResponse, error) {

	logger.Infof("GetTransactionsPaidByCustomer request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("GetTransactionsPaidByCustomer request validation failed", err, tag.Obj("request", req))
		return nil, err
	}

	if ctx.Value(CtxKeyUser) == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value(CtxKeyUser).(security.User)

	responses, err := business.GetTransactionsPaidByCustomer(ctx, tokenUser.Id, req.From.AsTime(), req.To.AsTime())
	if err != nil {
		logger.Error("failed to get customer transactions", err)
		return nil, err
	}

	logger.Info("GetTransactionsPaidByCustomer successfully called", tag.Obj("request", req))
	return &financePb.GetTransactionsPaidByCustomerResponse{
		Items: responses,
	}, nil
}

func (s serverImpl) GetTransactionsPaidToCourier(ctx context.Context, req *financePb.GetTransactionsPaidToCourierRequest) (*financePb.GetTransactionsPaidToCourierResponse, error) {

	logger.Infof("GetTransactionsPaidToCourier request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("GetTransactionsPaidToCourier request validation failed", err, tag.Obj("request", req))
		return nil, err
	}

	if ctx.Value(CtxKeyUser) == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value(CtxKeyUser).(security.User)

	transactions, err := business.GetTransactionsPaidToCourier(ctx, tokenUser.Id, req.From.AsTime(), req.To.AsTime(), int(req.PageNumber), int(req.PageSize))
	if err != nil {
		logger.Error("failed to get courier paid transactions", err)
		return nil, err
	}

	logger.Info("GetTransactionsPaidToCourier successfully called", tag.Obj("request", req))
	return &financePb.GetTransactionsPaidToCourierResponse{
		Transactions: transactions,
	}, nil
}

func (s serverImpl) GetAmountPaidToCourier(ctx context.Context, req *financePb.GetAmountPaidToCourierRequest) (*financePb.GetAmountPaidToCourierResponse, error) {

	logger.Infof("GetAmountPaidToCourier request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("GetAmountPaidToCourier request validation failed", err, tag.Obj("request", req))
		return nil, err
	}

	if ctx.Value(CtxKeyUser) == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value(CtxKeyUser).(security.User)

	totalAmount, currency, err := business.GetAmountPaidToCourier(ctx, tokenUser.Id, req.From.AsTime(), req.To.AsTime())
	if err != nil {
		logger.Error("failed to get courier paid amount", err)
		return nil, err
	}

	logger.Info("GetAmountPaidToCourier successfully called", tag.Obj("request", req))
	return &financePb.GetAmountPaidToCourierResponse{
		Amount:   totalAmount,
		Currency: currency,
	}, nil
}
