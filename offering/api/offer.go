package api

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	offeringPb "gitlab.artin.ai/backend/courier-management/grpc/offering/go"
	"gitlab.artin.ai/backend/courier-management/offering/business"
)

func (s serverImpl) HadCustomerRideWithCourier(ctx context.Context, req *offeringPb.HadCustomerRideWithCourierRequest) (*offeringPb.HadCustomerRideWithCourierResponse, error) {
	if err := req.Validate(); err != nil {
		logger.Error("request is not valid", err, tag.Obj("req", req))
		return nil, err
	}

	logger.Infof("HadCustomerRideWithCourier request = %+v", req)

	hadRide, err := business.HadCustomerRideWithCourier(ctx, req.CustomerId, req.CourierId, req.OfferId)
	if err != nil {
		logger.Error("failed to check if the customer had the ride with the courier", err)
		return nil, err
	}

	return &offeringPb.HadCustomerRideWithCourierResponse{
		HadRide: hadRide,
	}, nil
}

func (s serverImpl) GetOfferCourierAndCustomer(ctx context.Context, req *offeringPb.GetOfferCourierAndCustomerRequest) (*offeringPb.GetOfferCourierAndCustomerResponse, error) {

	logger.Infof("GetOfferCourierAndCustomer request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("request is not valid", err, tag.Obj("req", req))
		return nil, err
	}

	customerId, courierId, err := business.GetOfferCustomerAndCourier(ctx, req.OfferId)
	if err != nil {
		logger.Error("failed to check if the customer had the ride with the courier", err)
		return nil, err
	}

	return &offeringPb.GetOfferCourierAndCustomerResponse{
		CourierId:  courierId,
		CustomerId: customerId,
	}, nil
}
