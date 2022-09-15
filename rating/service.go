package rating

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/service"
	"github.com/kkjhamb01/courier-management/rating/api"
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
