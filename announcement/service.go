package announcement

import (
	"context"

	"github.com/kkjhamb01/courier-management/announcement/api"
	"github.com/kkjhamb01/courier-management/common/service"
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
