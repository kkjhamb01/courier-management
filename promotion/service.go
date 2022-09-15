package promotion

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/service"
	"gitlab.artin.ai/backend/courier-management/promotion/api"
)

func init() {
	service.RegisterService(promotionService{})
}

type promotionService struct {
}

func (s promotionService) Start(ctx context.Context) {
	api.CreateApiServer()
}

func (s promotionService) Stop() error {
	api.StopGrpcServer()
	return nil
}

func (s promotionService) Name() string {
	return service.Promotion
}

func (s promotionService) RelDir() string {
	return service.Promotion
}
