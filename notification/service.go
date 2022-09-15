package notification

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/kkjhamb01/courier-management/notification/api"
	"github.com/kkjhamb01/courier-management/notification/messaging"
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
