package services

import (
	"os"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"google.golang.org/grpc"
)

var DeliveryConnection *grpc.ClientConn

func ConnectToDelivery() (*grpc.ClientConn, error) {
	if DeliveryConnection == nil {
		var err error
		host := os.Getenv("DELIVERY_SERVICE_SERVICE_HOST")
		address := host + config.Delivery().GrpcPort
		DeliveryConnection, err = grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			logger.Fatalf("did not connect to the delivery service at: %s, error: %s", address, err)
			return nil, err
		}

		logger.Infof("connection established. Address: %v" + address)
	}

	return DeliveryConnection, nil
}
