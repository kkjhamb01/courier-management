package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/delivery/services"
	partypb "gitlab.artin.ai/backend/courier-management/party/proto"
)

func getCustomerProfile(ctx context.Context, customerId string) (*partypb.UserProfile, error) {

	logger.Infof("getCustomerProfile customerId = %v", customerId)

	conn, err := services.ConnectToParty()
	if err != nil {
		logger.Error("failed to connect to the party service", err)
		return nil, err
	}
	interServiceClient := partypb.NewInterServiceClient(conn)
	response, err := interServiceClient.OpenGetUserAccount(ctx, &partypb.OpenGetUserAccountRequest{
		UserId: customerId,
	})
	if err != nil {
		logger.Error("failed to get customer details from the party service", err)
		return nil, err
	}

	return response.Profile, nil
}
