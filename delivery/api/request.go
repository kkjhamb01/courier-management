package api

import (
	"context"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/delivery/business"
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s serverImpl) CreateRequest(ctx context.Context, req *deliveryPb.CreateRequestRequest) (*deliveryPb.CreateRequestResponse, error) {

	logger.Infof("CreateRequest request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate CreateRequestRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value("user").(security.User)

	var schedule *time.Time
	if req.Schedule != nil {
		scheduleTime := req.Schedule.AsTime()
		schedule = &scheduleTime
	}

	var requiredWorkers int32
	if req.RequiredWorkers != nil {
		requiredWorkers = req.RequiredWorkers.Value
	}

	request, err := business.CreateRequest(ctx, tokenUser.Id, req.VehicleType, req.Origin, req.Destinations, schedule, requiredWorkers)
	if err != nil {
		logger.Error("failed to create request", err)
		return nil, err
	}

	return &deliveryPb.CreateRequestResponse{
		CreatedRequest: &request,
	}, nil
}

func (s serverImpl) AcceptRequest(ctx context.Context, req *deliveryPb.AcceptRequestRequest) (*empty.Empty, error) {

	logger.Infof("AcceptRequest request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate AcceptRequestRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		logger.Error("failed to get user", err)
		return nil, err
	}

	tokenUser := ctx.Value("user").(security.User)

	var desc string
	if req.Description != nil {
		desc = req.Description.Value
	}

	err := business.AcceptRequest(ctx, req.Id, tokenUser.Id, desc)
	if err != nil {
		logger.Error("failed to accept request", err)
		return nil, err
	}

	logger.Info("AcceptRequest was called successfully")

	return &empty.Empty{}, nil
}

func (s serverImpl) RejectRequest(ctx context.Context, req *deliveryPb.RejectRequestRequest) (*empty.Empty, error) {

	logger.Infof("RejectRequest request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate RejectRequestRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	tokenUser := ctx.Value("user").(security.User)

	var desc string
	if req.Description != nil {
		desc = req.Description.Value
	}

	err := business.RejectRequest(ctx, req.Id, tokenUser.Id, desc)
	if err != nil {
		logger.Error("failed to accept request", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s serverImpl) RequestArrivedOrigin(ctx context.Context, req *deliveryPb.RequestArrivedOriginRequest) (*empty.Empty, error) {

	logger.Infof("RequestArrivedOrigin request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate RequestArrivedOrigin", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	var desc string
	if req.Description != nil {
		desc = req.Description.Value
	}
	err := business.RequestArrivedOrigin(ctx, req.Id, desc)
	if err != nil {
		logger.Error("failed to process request arrived", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s serverImpl) RequestArrivedDestination(ctx context.Context, req *deliveryPb.RequestArrivedDestinationRequest) (*empty.Empty, error) {

	logger.Infof("RequestArrivedDestination request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate RequestArrivedDestination", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	var desc string
	if req.Description != nil {
		desc = req.Description.Value
	}
	err := business.RequestArrivedDestination(ctx, req.Id, int(req.DestinationOrder), desc)
	if err != nil {
		logger.Error("failed to process request arrived", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s serverImpl) RequestPickedUp(ctx context.Context, req *deliveryPb.RequestPickedUpRequest) (*empty.Empty, error) {

	logger.Infof("RequestPickedUp request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate RejectRequestRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	var desc string
	if req.Description != nil {
		desc = req.Description.Value
	}

	err := business.RequestPickedUp(ctx, req.Id, req.Name, req.Signature, desc)
	if err != nil {
		logger.Error("failed to process request picked up", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s serverImpl) RequestNavigatingToSender(ctx context.Context, req *deliveryPb.RequestNavigatingToSenderRequest) (*empty.Empty, error) {

	logger.Infof("RequestNavigatingToSender request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate RequestNavigatingToSender", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	err := business.RequestNavigatingToSender(ctx, req.Id)
	if err != nil {
		logger.Error("failed to process request navigating to sender", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s serverImpl) RequestNavigatingToReceiver(ctx context.Context, req *deliveryPb.RequestNavigatingToReceiverRequest) (*empty.Empty, error) {

	logger.Infof("RequestNavigatingToReceiver request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate RequestNavigatingToReceiverRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	err := business.RequestNavigatingToReceiver(ctx, req.Id, int(req.TargetDestinationOrder))
	if err != nil {
		logger.Error("failed to process request navigating to sender", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s serverImpl) RequestDelivered(ctx context.Context, req *deliveryPb.RequestDeliveredRequest) (*empty.Empty, error) {

	logger.Infof("RequestDelivered request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate RejectRequestRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	var desc string
	if req.Description != nil {
		desc = req.Description.Value
	}

	err := business.RequestDelivered(ctx, req.Id, int(req.DestinationOrder), desc, req.Name, req.Signature)
	if err != nil {
		logger.Error("failed to process delivered request", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s serverImpl) RequestSenderNotAvailable(ctx context.Context, req *deliveryPb.RequestSenderNotAvailableRequest) (*empty.Empty, error) {

	logger.Infof("RequestSenderNotAvailable request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate RejectRequestRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	var desc string
	if req.Description != nil {
		desc = req.Description.Value
	}

	err := business.RequestSenderNotAvailable(ctx, req.Id, desc)
	if err != nil {
		logger.Error("failed to process request sender not available event", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s serverImpl) RequestReceiverNotAvailable(ctx context.Context, req *deliveryPb.RequestReceiverNotAvailableRequest) (*empty.Empty, error) {

	logger.Infof("RequestReceiverNotAvailable request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate RejectRequestRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	var desc string
	if req.Description != nil {
		desc = req.Description.Value
	}

	err := business.RequestReceiverNotAvailable(ctx, req.Id, int(req.DestinationOrder), desc)
	if err != nil {
		logger.Error("failed to process request receiver not available", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s serverImpl) RequestCompleted(ctx context.Context, req *deliveryPb.RequestCompletedRequest) (*empty.Empty, error) {

	logger.Infof("RequestCompleted request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate RequestCompletedRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	var desc string
	if req.Description != nil {
		desc = req.Description.Value
	}

	err := business.RequestCompleted(ctx, req.Id, desc)
	if err != nil {
		logger.Error("failed to process request completed", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s serverImpl) CancelRequest(ctx context.Context, req *deliveryPb.CancelRequestRequest) (*empty.Empty, error) {

	logger.Infof("CancelRequest request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate CancelRequestRequest", err)
		return nil, err
	}

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	tokenUser := ctx.Value("user").(security.User)

	var desc string
	if req.Description != nil {
		desc = req.Description.Value
	}
	err := business.CancelRequest(ctx, req.Id, req.CancelReason, req.CancelledBy, tokenUser.Id, desc)
	if err != nil {
		logger.Error("failed to accept request", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s serverImpl) GetRequests(ctx context.Context, req *deliveryPb.GetRequestsRequest) (*deliveryPb.GetRequestsResponse, error) {

	logger.Infof("GetRequests request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("failed to validate GetRequestRequest", err)
		return nil, err
	}

	request, err := business.GetRequests(ctx, req)
	if err != nil {
		logger.Error("failed to get requests", err)
		return nil, err
	}

	return &deliveryPb.GetRequestsResponse{
		Requests: request,
	}, nil
}

func (s serverImpl) GetCourierActiveRequest(ctx context.Context, req *emptypb.Empty) (*deliveryPb.GetCourierActiveRequestResponse, error) {

	logger.Infof("GetCourierActiveRequest request = %+v", req)

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value("user").(security.User)

	response, err := business.GetCourierActiveRequest(ctx, tokenUser.Id)
	if err != nil {
		logger.Error("failed to GetCourierActiveRequest", err)
		return nil, err
	}

	return response, nil
}

func (s serverImpl) GetCustomerActiveRequest(ctx context.Context, req *emptypb.Empty) (*deliveryPb.GetCustomerActiveRequestResponse, error) {

	logger.Infof("GetCustomerActiveRequest request = %+v", req)

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}
	tokenUser := ctx.Value("user").(security.User)

	response, err := business.GetCustomerActiveRequest(ctx, tokenUser.Id)
	if err != nil {
		logger.Error("failed to GetCustomerActiveRequest", err)
		return nil, err
	}

	return response, nil
}
