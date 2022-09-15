package rating

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/service"
	"gitlab.artin.ai/backend/courier-management/rating/api"
)

func init() {
	service.RegisterService(ratingService{})
}

type ratingService struct {
}

func (s ratingService) Start(ctx context.Context) {
	api.CreateApiServer()
}

func (s ratingService) Stop() error {
	api.StopGrpcServer()
	return nil
}

func (s ratingService) Name() string {
	return service.Rating
}

func (s ratingService) RelDir() string {
	return service.Rating
}
