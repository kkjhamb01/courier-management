package api

import (
	"context"
	"testing"
	"time"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	commonPb "github.com/kkjhamb01/courier-management/grpc/common/go"
	pricingPb "github.com/kkjhamb01/courier-management/grpc/pricing/go"
	"google.golang.org/grpc"
)

func init() {
	config.InitTestConfig()
	logger.InitLogger()
}

func TestGetPrice(t *testing.T) {
	destinations2 := make([]*commonPb.Location, 1)
	destinations2[0] = &commonPb.Location{
		Lat: 29.611436342227933,
		Lon: 52.48521186411381,
	}
	query := &pricingPb.CalculateCourierPriceRequest{
		VehicleType:     commonPb.VehicleType_ANY,
		RequiredWorkers: 0,
		Source: &commonPb.Location{
			Lat: 29.6056488,
			Lon: 52.484686100000005,
		},
		Destinations: destinations2,
	}
	conn := getPricingConn()
	defer conn.Close()
	clientDeadline := time.Now().Add(time.Duration(60000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	c := pricingPb.NewPricingClient(conn)

	res, err := c.CalculateCourierPrice(ctx, query)

	logger.Infof("res = %v, err = %v", res, err)
}

func getPricingConn() *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	opts = append(opts, grpc.WithBlock())
	logger.Infof("try to connect to pricing service %v", config.GetData().Pricing.GrpcPort)
	conn, err := grpc.Dial(":"+config.GetData().Pricing.GrpcPort, opts...)
	if err != nil {
		logger.Errorf("cannot connect to pricing: %v", err)
	}
	return conn
}
