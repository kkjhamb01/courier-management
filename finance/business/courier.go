package business

import (
	"context"
	"errors"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/common/logger/tag"
	"github.com/kkjhamb01/courier-management/finance/services"
	"github.com/kkjhamb01/courier-management/party/proto"
)

func getCourierInfo(ctx context.Context, accessToken string) (*proto.CourierProfile, error) {

	logger.Infof("getCourierInfo")

	conn, err := services.ConnectToParty()
	if err != nil {
		logger.Error("failed to connect to the party service", err)
		return nil, err
	}
	partyC := proto.NewCourierAccountServiceClient(conn)

	req := &proto.GetCourierAccountRequest{
		AccessToken: accessToken,
	}

	response, err := partyC.GetCourierAccount(ctx, req)
	if err != nil {
		logger.Error("failed to get courier account from the party service", err)
		return nil, err
	}

	if response == nil || response.Profile == nil {
		err = errors.New("party response is nil")
		logger.Error("no profile was returned from the party service", err, tag.Obj("request", req), tag.Obj("response", response))
		return nil, err
	}

	return response.Profile, nil
}
