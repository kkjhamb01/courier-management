package party

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/service"
	"gitlab.artin.ai/backend/courier-management/party/api"
)

func init() {
	service.RegisterService(partyService{})
}

type partyService struct {
}

func (s partyService) Start(ctx context.Context) {
	api.CreateApiServer()
}

func (s partyService) Stop() error {
	api.StopGrpcServer()
	return nil
}

func (s partyService) Name() string {
	return service.Party
}

func (s partyService) RelDir() string {
	return service.Party
}
