package notification

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/service"
	"gitlab.artin.ai/backend/courier-management/notification/api"
	"gitlab.artin.ai/backend/courier-management/notification/messaging"
)

func init() {
	service.RegisterService(notificationService{})
}

type notificationService struct {
}

func (s notificationService) Start(ctx context.Context) {
	messaging.StartSubscriptions()
	api.CreateApiServer()
}

func (s notificationService) Stop() error {
	messaging.StopSubscriptions()
	api.StopGrpcServer()
	return nil
}

func (s notificationService) Name() string {
	return service.Notification
}

func (s notificationService) RelDir() string {
	return service.Notification
}
