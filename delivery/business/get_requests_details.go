package business

import (
	"context"
	"errors"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/delivery/db"
	"github.com/kkjhamb01/courier-management/delivery/model"
	"github.com/kkjhamb01/courier-management/delivery/services"
	deliveryPb "github.com/kkjhamb01/courier-management/grpc/delivery/go"
	"github.com/kkjhamb01/courier-management/party/proto"
	uaa "github.com/kkjhamb01/courier-management/uaa/proto"
)

func GetCourierRequestDetails(ctx context.Context, requestId string) (deliveryPb.Request, error) {

	logger.Infof("GetCourierRequestDetails requestId = %v", requestId)

	var request model.Request
	err := db.MariaDbClient().
		Preload("Locations").
		Where("ID = ?", requestId).
		First(&request).
		Error
	if err != nil {
		logger.Error("failed to fetch courier requests", err)
		return deliveryPb.Request{}, uaa.Internal.Error(err)
	}

	return request.ToProto(), nil
}

func GetCustomerRequestDetails(ctx context.Context, requestId string) (deliveryPb.GetCustomerRequestDetailsResponse, error) {

	logger.Infof("GetCourierRequestDetails requestId = %v", requestId)

	var request model.Request
	err := db.MariaDbClient().
		Preload("Locations").
		Where("ID = ?", requestId).
		First(&request).
		Error
	if err != nil {
		logger.Error("failed to fetch courier requests", err)
		return deliveryPb.GetCustomerRequestDetailsResponse{}, uaa.Internal.Error(err)
	}

	conn, err := services.ConnectToParty()
	if err != nil {
		logger.Error("failed to connect to the party service", err)
		return deliveryPb.GetCustomerRequestDetailsResponse{}, uaa.Internal.Error(err)
	}

	// get plate number
	// TODO skipping party service errors for now to make testing easier, change it back later
	var plateNumber string
	var photo []byte
	var fullName string
	var phoneNumber string

	if request.CourierId != "" {
		partyService := proto.NewInterServiceClient(conn)
		plateResponse, err := partyService.InterServiceGetProfileAdditionalInfo(ctx, &proto.InterServiceGetProfileAdditionalInfoRequest{
			UserId: request.CourierId,
			Type:   proto.AdditionalInfoType_ADDITIONAL_INFO_TYPE_MOT,
		})
		if err != nil {
			logger.Error("failed to get profile addition info from the party service", err)
			return deliveryPb.GetCustomerRequestDetailsResponse{}, uaa.Internal.Error(err)
		} else {
			plateNumber = plateResponse.GetMot().RegistrationNumber
		}

		// get photo
		// TODO Skipping party service errors for now to make testing easier, change it back later

		photoResponse, err := partyService.InterServiceGetDocumentsOfUser(ctx, &proto.InterServiceGetDocumentsOfUserRequest{
			UserId:   request.CourierId,
			Type:     proto.DocumentInfoType_DOCUMENT_INFO_TYPE_PROFILE_PICTURE,
			DataType: proto.DocumentDataType_DOCUMENT_DATA_TYPE_DATA,
		})
		if err != nil {
			logger.Error("failed to get profile documents from the party service", err)
			return deliveryPb.GetCustomerRequestDetailsResponse{}, uaa.Internal.Error(err)
		} else if len(photoResponse.Documents) == 0 {
			err = errors.New("InterServiceGetDocumentsOfUserResponse.Documents is empty")
			logger.Error("no document returned from the party service", err)
			//return deliveryPb.GetCustomerRequestDetailsResponse{}, uaa.NotFound.Error(err)
		} else {
			photo = photoResponse.Documents[0].Data
		}

		// get name and phone number
		// TODO Skipping party service errors for now to make testing easier, change it back later
		profileResponse, err := partyService.InterServiceFindCourierAccounts(ctx, &proto.InterServiceFindCourierAccountsRequest{
			Filter: &proto.InterServiceFindCourierAccountsRequest_UserId{
				UserId: request.CourierId,
			},
		})
		if err != nil {
			logger.Error("failed to get profile info from the party service", err)
			return deliveryPb.GetCustomerRequestDetailsResponse{}, uaa.Internal.Error(err)
		} else if len(profileResponse.Profiles) == 0 {
			err = errors.New("InterServiceFindCourierAccountsResponse.Profiles is empty")
			logger.Error("no profile returned from the party service", err)
			return deliveryPb.GetCustomerRequestDetailsResponse{}, uaa.NotFound.Error(err)
		} else {
			fullName = profileResponse.Profiles[0].FirstName + " " + profileResponse.Profiles[0].LastName
			phoneNumber = profileResponse.Profiles[0].PhoneNumber
		}

	}

	return deliveryPb.GetCustomerRequestDetailsResponse{
		Request:            request.ToProtoP(),
		Plate:              plateNumber,
		CourierPhoto:       photo,
		CourierPhoneNumber: phoneNumber,
		CourierName:        fullName,
	}, nil
}
