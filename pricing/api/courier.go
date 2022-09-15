package api

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	pricingPb "github.com/kkjhamb01/courier-management/grpc/pricing/go"
	"github.com/kkjhamb01/courier-management/pricing/business"
)

func (s serverImpl) CalculateCourierPrice(ctx context.Context, req *pricingPb.CalculateCourierPriceRequest) (*pricingPb.CalculateCourierPriceResponse, error) {

	logger.Infof("CalculateCourierPrice request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("CalculateCourierPrice request validation failed", err, tag.Obj("request", req))
		return nil, err
	}

	price, err := business.CalculateCourierPrice(ctx, req)
	if err != nil {
		logger.Error("CalculateCourierPrice failed to calculate price", err)
		return nil, err
	}

	logger.Info("CalculateCourierPrice successfully called", tag.Obj("request", req))
	return price, nil
}

func (s serverImpl) ReviewCourierPrice(ctx context.Context, req *pricingPb.ReviewCourierPriceRequest) (*pricingPb.ReviewCourierPriceResponse, error) {

	logger.Infof("ReviewCourierPrice request = %+v", req)

	if err := req.Validate(); err != nil {
		logger.Error("ReviewCourierPrice request validation failed", err, tag.Obj("request", req))
		return nil, err
	}

	price, err := business.ReviewCourierPrice(ctx, req)
	if err != nil {
		logger.Error("ReviewCourierPrice failed to review price", err)
		return nil, err
	}

	logger.Info("ReviewCourierPrice successfully called", tag.Obj("request", req))
	return price, nil
}
