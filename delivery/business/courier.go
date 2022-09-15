package business

import (
	"context"
	"errors"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/delivery/services"
	partypb "github.com/kkjhamb01/courier-management/party/proto"
	"github.com/kkjhamb01/courier-management/uaa/proto"
)

func getCourierProfile(ctx context.Context, courierId string) (*partypb.CourierProfile, error) {

	logger.Infof("getCourierProfile courierId = %v", courierId)

	conn, err := services.ConnectToParty()
	if err != nil {
		logger.Error("failed to connect to the party service", err)
		return nil, err
	}
	interServiceClient := partypb.NewInterServiceClient(conn)
	courierData, err := interServiceClient.InterServiceFindCourierAccounts(ctx, &partypb.InterServiceFindCourierAccountsRequest{
		Filter: &partypb.InterServiceFindCourierAccountsRequest_UserId{
			UserId: courierId,
		},
	})
	if err != nil {
		logger.Error("failed to get courier details from the party service", err)
		return nil, err
	}
	if len(courierData.Profiles) == 0 {
		err = errors.New("no profile returned from the party service")
		logger.Error("failed to get courier details from the party service", err)
		return nil, proto.NotFound.Error(err)
	}

	return courierData.Profiles[0], nil
}

func getCourierPlateNumber(ctx context.Context, courierId string) (string, error) {
	conn, err := services.ConnectToParty()
	if err != nil {
		logger.Error("failed to connect to the party service", err)
		return "", err
	}

	// get plate number
	partyService := partypb.NewInterServiceClient(conn)
	plateResponse, err := partyService.InterServiceGetProfileAdditionalInfo(ctx, &partypb.InterServiceGetProfileAdditionalInfoRequest{
		UserId: courierId,
		Type:   partypb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_MOT,
	})
	if err != nil {
		logger.Error("failed to get profile addition info from the party service", err)
		return "", err
	}

	return plateResponse.GetMot().RegistrationNumber, nil
}

func getCourierPhoto(ctx context.Context, courierId string) ([]byte, error) {
	conn, err := services.ConnectToParty()
	if err != nil {
		logger.Error("failed to connect to the party service", err)
		return nil, err
	}

	partyService := partypb.NewInterServiceClient(conn)

	photoResponse, err := partyService.InterServiceGetDocumentsOfUser(ctx, &partypb.InterServiceGetDocumentsOfUserRequest{
		UserId:   courierId,
		Type:     partypb.DocumentInfoType_DOCUMENT_INFO_TYPE_PROFILE_PICTURE,
		DataType: partypb.DocumentDataType_DOCUMENT_DATA_TYPE_DATA,
	})
	if err != nil {
		logger.Error("failed to get profile documents from the party service", err)
		return nil, err
	} else if len(photoResponse.Documents) == 0 {
		err = errors.New("InterServiceGetDocumentsOfUserResponse.Documents is empty")
		logger.Error("no document returned from the party service", err)
		return nil, proto.NotFound.Error(err)
	}

	return photoResponse.Documents[0].Data, nil

}
