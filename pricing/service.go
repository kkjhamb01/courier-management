package pricing

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/kkjhamb01/courier-management/pricing/api"
)

func init() {
	service.RegisterService(pricingService{})
}

type pricingService struct {
}

func (s pricingService) Start(ctx context.Context) {
	// start offering gRPC server
	api.CreateGrpcServer()
}

func (s pricingService) Stop() error {
	api.StopGrpcServer()
	return nil
}

func (s pricingService) Name() string {
	return service.Pricing
}

func (s pricingService) RelDir() string {
	return service.Pricing
}
