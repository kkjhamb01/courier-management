package business

import (
	"context"
	"errors"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/finance/services"
	"gitlab.artin.ai/backend/courier-management/party/proto"
)

func getCustomerInfo(ctx context.Context, accessToken string) (*proto.UserProfile, error) {

	logger.Infof("getCustomerInfo")

	conn, err := services.ConnectToParty()
	if err != nil {
		logger.Error("failed to connect to the party service", err)
		return nil, err
	}
	partyC := proto.NewUserAccountServiceClient(conn)

	req := &proto.GetUserAccountRequest{
		AccessToken: accessToken,
	}

	response, err := partyC.GetUserAccount(ctx, req)
	if err != nil {
		logger.Error("failed to get user account from the party service", err)
		return nil, err
	}

	if response == nil || response.Profile == nil {
		err = errors.New("party response is nil")
		logger.Error("no profile was returned from the party service", err, tag.Obj("request", req), tag.Obj("response", response))
		return nil, err
	}

	return response.Profile, nil
}
