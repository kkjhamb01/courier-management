package announcement

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/announcement/api"
	"gitlab.artin.ai/backend/courier-management/common/service"
)

func init() {
	service.RegisterService(announcementService{})
}

type announcementService struct {
}

func (s announcementService) Start(ctx context.Context) {
	api.CreateApiServer()
}

func (s announcementService) Stop() error {
	api.StopGrpcServer()
	return nil
}

func (s announcementService) Name() string {
	return service.Announcement
}

func (s announcementService) RelDir() string {
	return service.Announcement
}
