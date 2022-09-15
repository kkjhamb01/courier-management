package services

import (
	"os"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"google.golang.org/grpc"
)

var PartyConnection *grpc.ClientConn

func ConnectToParty() (*grpc.ClientConn, error) {
	if PartyConnection == nil {
		var err error
		host := os.Getenv("PARTY_SERVICE_SERVICE_HOST")
		address := host + config.Party().Server.Address
		PartyConnection, err = grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			logger.Fatalf("did not connect to the party service at: %s, error: %s", address, err)
			return nil, err
		}

		logger.Infof("connection established. Address: %v" + address)
	}

	return PartyConnection, nil
}
