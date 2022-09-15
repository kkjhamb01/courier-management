package api

import (
	"context"
	"errors"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/delivery/business"
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"google.golang.org/protobuf/types/known/durationpb"
)

func (s serverImpl) GetCourierCompletedRequests(ctx context.Context, req *deliveryPb.GetCourierCompletedRequestsRequest) (*deliveryPb.GetCourierCompletedRequestsResponse, error) {

	logger.Infof("GetCourierCompletedRequests request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate GetCourierCompletedRequestsRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value("user").(security.User)

	completedRequests, err := business.GetCourierCompletedRequests(ctx, tokenUser.Id, req.From.AsTime(), req.To.AsTime())
	if err != nil {
		logger.Error("failed to get courier completed request", err)
		return nil, err
	}

	return &deliveryPb.GetCourierCompletedRequestsResponse{
		CompletedRequests: int32(completedRequests),
	}, nil
}

func (s serverImpl) GetCustomerCompletedRequests(ctx context.Context, req *deliveryPb.GetCustomerCompletedRequestsRequest) (*deliveryPb.GetCustomerCompletedRequestsResponse, error) {

	logger.Infof("GetCustomerCompletedRequests request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate GetCustomerCompletedRequestsRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value("user").(security.User)

	completedRequests, err := business.GetCustomerCompletedRequests(ctx, tokenUser.Id, req.From.AsTime(), req.To.AsTime())
	if err != nil {
		logger.Error("failed to get customer completed request", err)
		return nil, err
	}

	return &deliveryPb.GetCustomerCompletedRequestsResponse{
		CompletedRequests: int32(completedRequests),
	}, nil
}

func (s serverImpl) GetCourierRequestsDuration(ctx context.Context, req *deliveryPb.GetCourierRequestsDurationRequest) (*deliveryPb.GetCourierRequestsDurationResponse, error) {

	logger.Infof("GetCourierRequestsDuration request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate GetCourierRequestsDurationRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value("user").(security.User)

	totalTime, err := business.GetCourierRequestsDuration(ctx, tokenUser.Id, req.From.AsTime(), req.To.AsTime())
	if err != nil {
		logger.Error("failed to get courier requests duration", err)
		return nil, err
	}

	return &deliveryPb.GetCourierRequestsDurationResponse{
		TimeRange: durationpb.New(totalTime),
	}, nil
}

func (s serverImpl) GetCustomerRequestsDuration(ctx context.Context, req *deliveryPb.GetCustomerRequestsDurationRequest) (*deliveryPb.GetCustomerRequestsDurationResponse, error) {

	logger.Infof("GetCustomerRequestsDuration request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate GetCustomerRequestsDurationRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value("user").(security.User)

	totalTime, err := business.GetCustomerRequestsDuration(ctx, tokenUser.Id, req.From.AsTime(), req.To.AsTime())
	if err != nil {
		logger.Error("failed to get customer requests duration", err)
		return nil, err
	}

	return &deliveryPb.GetCustomerRequestsDurationResponse{
		TimeRange: durationpb.New(totalTime),
	}, nil
}

func (s serverImpl) GetCourierRequestsMileage(ctx context.Context, req *deliveryPb.GetCourierRequestsMileageRequest) (*deliveryPb.GetCourierRequestsMileageResponse, error) {

	logger.Infof("GetCourierRequestsMileage request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate GetCourierRequestsMileageRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value("user").(security.User)

	totalMileage, err := business.GetCourierRequestsMileage(ctx, tokenUser.Id, req.From.AsTime(), req.To.AsTime())
	if err != nil {
		logger.Error("failed to get courier requests mileage", err)
		return nil, err
	}

	return &deliveryPb.GetCourierRequestsMileageResponse{
		Mileage: int32(totalMileage),
	}, nil
}

func (s serverImpl) GetCustomerRequestsMileage(ctx context.Context, req *deliveryPb.GetCustomerRequestsMileageRequest) (*deliveryPb.GetCustomerRequestsMileageResponse, error) {

	logger.Infof("GetCustomerRequestsMileage request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate GetCustomerRequestsMileageRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value("user").(security.User)

	totalMileage, err := business.GetCustomerRequestsMileage(ctx, tokenUser.Id, req.From.AsTime(), req.To.AsTime())
	if err != nil {
		logger.Error("failed to get customer requests mileage", err)
		return nil, err
	}

	return &deliveryPb.GetCustomerRequestsMileageResponse{
		Mileage: int32(totalMileage),
	}, nil
}

func (s serverImpl) GetCourierRequestsHistory(ctx context.Context, req *deliveryPb.GetCourierRequestsHistoryRequest) (*deliveryPb.GetCourierRequestsHistoryResponse, error) {

	logger.Infof("GetCourierRequestsHistory request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate GetCourierRequestsHistoryRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value("user").(security.User)

	requestsHistory, err := business.GetCourierRequestsHistory(ctx, tokenUser.Id, int(req.PageSize), int(req.PageNumber))
	if err != nil {
		logger.Error("failed to get courier requests history", err)
		return nil, err
	}

	return &deliveryPb.GetCourierRequestsHistoryResponse{
		Items: requestsHistory,
	}, nil
}

func (s serverImpl) GetCustomerRequestsHistory(ctx context.Context, req *deliveryPb.GetCustomerRequestsHistoryRequest) (*deliveryPb.GetCustomerRequestsHistoryResponse, error) {

	logger.Infof("GetCustomerRequestsHistory request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate GetCustomerRequestsHistoryRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value("user").(security.User)

	requestsHistory, err := business.GetCustomerRequestsHistory(ctx, tokenUser.Id, int(req.PageSize), int(req.PageNumber))
	if err != nil {
		logger.Error("failed to get customer requests history", err)
		return nil, err
	}

	return &deliveryPb.GetCustomerRequestsHistoryResponse{
		Items: requestsHistory,
	}, nil
}

func (s serverImpl) GetCourierRequestDetails(ctx context.Context, req *deliveryPb.GetCourierRequestDetailsRequest) (*deliveryPb.GetCourierRequestDetailsResponse, error) {

	logger.Infof("GetCourierRequestDetails request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate GetCourierRequestDetailsRequest", err)
		return nil, err
	}

	requestDetails, err := business.GetCourierRequestDetails(ctx, req.RequestId)
	if err != nil {
		logger.Error("failed to get courier request details", err)
		return nil, err
	}

	return &deliveryPb.GetCourierRequestDetailsResponse{
		Request: &requestDetails,
	}, nil
}

func (s serverImpl) GetCustomerRequestDetails(ctx context.Context, req *deliveryPb.GetCustomerRequestDetailsRequest) (*deliveryPb.GetCustomerRequestDetailsResponse, error) {

	logger.Infof("GetCustomerRequestDetails request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate GetCustomerRequestDetailsRequest", err)
		return nil, err
	}

	response, err := business.GetCustomerRequestDetails(ctx, req.RequestId)
	if err != nil {
		logger.Error("failed to get customer request details", err)
		return nil, err
	}

	return &response, nil
}

func (s serverImpl) GetRequestHistory(ctx context.Context, req *deliveryPb.GetRequestHistoryRequest) (*deliveryPb.GetRequestHistoryResponse, error) {

	//logger.Infof("GetRequestHistory request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate GetCustomerRequestsHistoryRequest", err)
		return nil, err
	}

	response, err := business.GetRequestHistory(ctx, req.RequestId)
	if err != nil {
		logger.Error("failed to get requests history", err)
		return nil, err
	}

	return response, nil
}
