package services

import (
	"os"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"google.golang.org/grpc"
)

var partyConnection *grpc.ClientConn

func ConnectToParty() (*grpc.ClientConn, error) {
	if partyConnection == nil {
		var err error
		host := os.Getenv("PARTY_SERVICE_SERVICE_HOST")
		address := host + config.Party().Server.Address
		partyConnection, err = grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			logger.Errorf("did not connect to party: %s", err, address)
			return nil, err
		}

		logger.Infof("connection established. Address: %v" + address)
	}

	return partyConnection, nil
}
