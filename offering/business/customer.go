package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/offering/services"
	partyPb "gitlab.artin.ai/backend/courier-management/party/proto"
)

func IsCustomerActive(ctx context.Context, customerId string) (_ bool, _err error) {

	logger.Infof("IsCustomerActive customerId = %v", customerId)

	conn, err := services.ConnectToParty()
	if err != nil {
		logger.Error("failed to establish a connection to the party service", err)
		return false, err
	}

	partyC := partyPb.NewUserStatusServiceClient(conn)
	courierStatus, err := partyC.GetClientUserStatus(ctx, &partyPb.GetClientUserStatusRequest{
		UserId: customerId,
	})
	if err != nil {
		logger.Error("failed to get client account", err)
		return false, err
	}

	return courierStatus.Status == partyPb.UserStatus_USER_STATUS_AVAILABLE, nil
}

func getCustomerProfile(ctx context.Context, customerId string) (*partyPb.UserProfile, error) {

	logger.Infof("getCustomerProfile customerId = %v", customerId)

	conn, err := services.ConnectToParty()
	if err != nil {
		logger.Error("failed to connect to the party service", err)
		return nil, err
	}
	interServiceClient := partyPb.NewInterServiceClient(conn)
	response, err := interServiceClient.OpenGetUserAccount(ctx, &partyPb.OpenGetUserAccountRequest{
		UserId: customerId,
	})
	if err != nil {
		logger.Error("failed to get customer details from the party service", err)
		return nil, err
	}

	return response.Profile, nil
}
