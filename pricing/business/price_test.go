package business

import (
	"context"
	"testing"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	commonPb "github.com/kkjhamb01/courier-management/grpc/common/go"
	pricingPb "github.com/kkjhamb01/courier-management/grpc/pricing/go"
)

func init() {
	config.InitTestConfig()
	logger.InitLogger()
}

func TestDistanceApi(t *testing.T) {
	response, err := CalculateCourierPrice(context.Background(), &pricingPb.CalculateCourierPriceRequest{
		VehicleType:     commonPb.VehicleType_CAR,
		RequiredWorkers: 10,
		Source: &commonPb.Location{
			Lat: 35,
			Lon: 52,
		},
		Destinations: []*commonPb.Location{
			{
				Lat: 36,
				Lon: 53,
			},
		},
	})
	logger.Infof("err = %v", err)
	logger.Infof("response = %v", response.Amount)
}
