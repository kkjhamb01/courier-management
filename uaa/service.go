package uaa

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/service"
	"gitlab.artin.ai/backend/courier-management/uaa/api"
)

func init() {
	service.RegisterService(uaaService{})
}

type uaaService struct {
}

func (s uaaService) Start(ctx context.Context) {
	api.StartWebServer()
	api.CreateApiServer()
}

func (s uaaService) Stop() error {
	api.StopGrpcServer()
	return nil
}

func (s uaaService) Name() string {
	return service.Uaa
}

func (s uaaService) RelDir() string {
	return service.Uaa
}
