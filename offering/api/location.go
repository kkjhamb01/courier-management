package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	commonPb "gitlab.artin.ai/backend/courier-management/grpc/common/go"
	"gitlab.artin.ai/backend/courier-management/grpc/offering/go"
	"gitlab.artin.ai/backend/courier-management/offering/business"
	"gitlab.artin.ai/backend/courier-management/uaa/security"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"sync"
	"time"
)

func (s serverImpl) SetCourierLiveLocation(stream offeringPb.Offering_SetCourierLiveLocationServer) error {
	ctx := stream.Context()

	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return err
	}

	tokenUser, ok := ctx.Value("user").(security.User)
	if !ok {
		err := errors.New("failed to cast user value found in context to security.User")
		logger.Error("failed to get user", err)
		return err
	}
	courierId := tokenUser.Id

	logger.Infof("SetCourierLiveLocation courierId = %v", courierId)

	accessToken, ok := ctx.Value("access_token").(string)
	if !ok {
		err := errors.New("failed to cast access_token value found in context to string")
		logger.Error("failed to get accessToken", err)
		return err
	}

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			logger.Info("SetCourierLiveLocation done")
			return nil
		}

		//logger.Info("SetCourierLiveLocation sending result", tag.Obj("request", request))

		if err != nil {
			//logger.Error("failed to receive the SetLiveLocation request", err)
			// TODO ignore the send() error for now
			stream.Send(&offeringPb.SetCourierLiveLocationResponse{
				Successful: false,
				Message:    err.Error(),
			})
			continue
		}

		if err := request.Validate(); err != nil {
			logger.Error("the request is not valid", err, tag.Obj("req", request))
			// TODO ignore the send() error for now
			stream.Send(&offeringPb.SetCourierLiveLocationResponse{
				Successful: false,
				Message:    err.Error(),
			})
			continue
		}

		//logger.Infof("SetCourierLiveLocation SetCourierLocation courierId = %v, request = %+v", courierId, request)

		err = business.SetCourierLocation(ctx, *request.GetLocation(), accessToken, courierId)
		if err != nil {
			logger.Errorf("failed to store the courier location : %v", err, request.GetLocation())
			// TODO ignore the send() error for now
			stream.Send(&offeringPb.SetCourierLiveLocationResponse{
				Successful: false,
				Message:    err.Error(),
			})
			continue
		}

		err = stream.Send(&offeringPb.SetCourierLiveLocationResponse{
			Successful: true,
			Message:    "done",
		})
		if err != nil {
			logger.Errorf("failed to send the SetLiveLocation result : %v", err, request.GetLocation())
		}
	}
}

func (s serverImpl) GetCourierLiveLocation(request *offeringPb.GetCourierLiveLocationRequest, stream offeringPb.Offering_GetCourierLiveLocationServer) error {
	ctx := stream.Context()

	if err := request.Validate(); err != nil {
		logger.Error("the request is not valid", err, tag.Obj("req", request))
		return status.Error(codes.InvalidArgument, err.Error())
	}

	ticker := time.NewTicker(time.Duration(request.IntervalSeconds) * time.Second)

	logger.Infof("GetCourierLiveLocation request = %+v", request)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(ctx context.Context) {
		for {
			select {
			case <-ticker.C:
				//logger.Info("GetCourierLiveLocation sending result", tag.Obj("request", request))
				location, err := business.GetCourierLocation(ctx, request.CourierId)
				if err != nil {
					logger.Error("failed to get courier location", err)
					//do nothing to tolerate error
					continue
				}

				timestamp, err := ptypes.TimestampProto(time.Now())
				if err != nil {
					logger.Error("failed to convert time to timestamp", err)
					//do nothing to tolerate error
					continue
				}

				//logger.Info("return result", tag.Obj("location", location))
				err = stream.Send(&offeringPb.GetCourierLiveLocationResponse{
					Location: &location,
					Time:     timestamp,
				})
				if err != nil {
					logger.Error("failed to get the location", err)
					//do nothing to tolerate error
					continue
				}
			case <-ctx.Done():
				wg.Done()
				logger.Info("GetCourierLiveLocation Done")
				return
			}
		}
	}(ctx)

	wg.Wait()

	return nil
}

func (s serverImpl) SetCourierLocation(ctx context.Context, request *offeringPb.SetCourierLocationRequest) (*emptypb.Empty, error) {
	if ctx.Value("user") == nil {
		err := errors.New("no user found in context")
		return nil, err
	}

	tokenUser, ok := ctx.Value("user").(security.User)
	if !ok {
		err := errors.New("failed to cast user value found in context to security.User")
		logger.Error("failed to get user", err)
		return nil, err
	}
	courierId := tokenUser.Id

	logger.Infof("SetCourierLocation courierId = %v, request = %+v", courierId, request)

	accessToken, ok := ctx.Value("access_token").(string)
	if !ok {
		err := errors.New("failed to cast access_token value found in context to string")
		logger.Error("failed to get accessToken", err)
		return nil, err
	}

	err := business.SetCourierLocation(ctx,
		*request.Location,
		accessToken,
		courierId)
	if err != nil {
		logger.Error("setCourierLocation error.", err, tag.Obj("request", request))
		return nil, err
	}

	logger.Info("setCourierLocation successfully called", tag.Obj("request", request))
	return &emptypb.Empty{}, nil
}

func (s serverImpl) GetNearbyCouriers(ctx context.Context, request *offeringPb.GetNearbyCouriersRequest) (*offeringPb.GetNearbyCouriersResponse, error) {

	logger.Infof("GetNearbyCouriers request = %+v", request)

	courierLocations, err := business.GetNearbyCouriers(ctx,
		*request.Location,
		int(request.RadiusMeter),
		request.VehicleType)
	if err != nil {
		logger.Error("getNearByCouriers error", err, tag.Obj("request", request))
		return nil, err
	}

	var couriersPointer = make([]*commonPb.Courier, len(courierLocations))
	for i, courierLocation := range courierLocations {
		courierLocationCopy := courierLocation
		couriersPointer[i] = courierLocationCopy.Courier
	}
	logger.Infof("api.GetNearbyCouriers, courierLocations: %+v", couriersPointer)

	for _, courier := range couriersPointer {
		isActive, err := business.IsCourierActive(ctx, courier.Id)
		if err != nil {
			logger.Error("failed to check if the courierLocation is active", err)
			return nil, err
		}
		if !isActive {
			err = fmt.Errorf("the courierLocation %v is not active", courier.Id)
			logger.Error("failed to get nearby courierLocations", err)
			return nil, err
		}
	}

	return &offeringPb.GetNearbyCouriersResponse{
		Couriers: couriersPointer,
	}, nil
}
