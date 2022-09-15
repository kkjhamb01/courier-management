package business

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	commonPb "gitlab.artin.ai/backend/courier-management/grpc/common/go"
	"gitlab.artin.ai/backend/courier-management/offering/db"
	"gitlab.artin.ai/backend/courier-management/offering/model"
	"gitlab.artin.ai/backend/courier-management/offering/services"
	"gitlab.artin.ai/backend/courier-management/offering/storage"
	"gitlab.artin.ai/backend/courier-management/party/proto"
	"time"
)

func GetCourierType(ctx context.Context, accessToken string, courierId string) (_ commonPb.VehicleType, _err error) {

	logger.Infof("GetCourierType courierId = %v", courierId)

	tx := storage.CreateTx()
	defer func() {
		if _err == nil {
			_err = tx.Commit(ctx)
		}

		var rollbackErr error
		if _err != nil {
			rollbackErr = tx.Rollback()
		}
		if rollbackErr != nil {
			_err = multierror.Append(_err, rollbackErr).ErrorOrNil()
		}
	}()

	var courierType commonPb.VehicleType
	//courierType, err := tx.GetCourierType(ctx, courierId)
	//if err == nil {
	//	return courierType, nil
	//}

	conn, err := services.ConnectToParty()
	if err != nil {
		logger.Error("failed to establish a connection to the party service", err)
		return commonPb.VehicleType_ANY, err
	}

	partyC := proto.NewCourierAccountServiceClient(conn)
	courierAccount, err := partyC.GetCourierAccount(ctx, &proto.GetCourierAccountRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		logger.Error("failed to get courier account", err)
		return commonPb.VehicleType_ANY, err
	}

	switch courierAccount.Profile.TransportType {
	case proto.TransportationType_TRANSPORTATION_TYPE_BICYCLE:
		courierType = commonPb.VehicleType_BICYCLE
	case proto.TransportationType_TRANSPORTATION_TYPE_CAR:
		courierType = commonPb.VehicleType_CAR
	case proto.TransportationType_TRANSPORTATION_TYPE_MOTORBIKE:
		courierType = commonPb.VehicleType_MOTORBIKE
	case proto.TransportationType_TRANSPORTATION_TYPE_VAN:
		switch courierAccount.Profile.TransportSize {
		case proto.TransportationSize_TRANSPORTATION_SIZE_SMALL:
			courierType = commonPb.VehicleType_SMALL_VAN
		case proto.TransportationSize_TRANSPORTATION_SIZE_MEDIUM:
			courierType = commonPb.VehicleType_MEDIUM_VAN
		case proto.TransportationSize_TRANSPORTATION_SIZE_LARGE:
			courierType = commonPb.VehicleType_LARGE_VAN
		}

	default:
		err = fmt.Errorf("the courier type is not known, party.transportationType: %v", courierAccount.Profile.TransportType)
		logger.Error("no courier type is matched", err)
		return commonPb.VehicleType_ANY, err
	}

	err = tx.SetCourierType(ctx, courierId, courierType)
	if err != nil {
		logger.Error("failed to store courier type in cache", err)
	}

	return courierType, err
}

func IsCourierActive(ctx context.Context, courierId string) (_ bool, _err error) {

	logger.Infof("IsCourierActive courierId = %v", courierId)

	if s, err := db.OfferRedisClient().Get(ctx, fmt.Sprintf("alive-%s", courierId)).Result(); err != nil || s == "" {
		logger.Debugf("IsCourierActive is not alive courierId = %v, err = %v", courierId, err)
		return false, err
	}

	conn, err := services.ConnectToParty()
	if err != nil {
		logger.Error("failed to establish a connection to the party service", err)
		return false, err
	}

	partyC := proto.NewUserStatusServiceClient(conn)
	courierStatus, err := partyC.GetCourierUserStatus(ctx, &proto.GetCourierUserStatusRequest{
		UserId: courierId,
	})
	if err != nil {
		logger.Error("failed to get courier account", err)
		return false, err
	}

	return courierStatus.Status == proto.UserStatus_USER_STATUS_AVAILABLE, nil
}

func OnCourierStatusChange(ctx context.Context, courierId string, newStatus commonPb.CourierStatus, updatedAt time.Time) error {

	logger.Infof("OnCourierStatusChange courierId = %v, newStatus = %v", courierId, newStatus)

	courierLog := model.CourierStatusLog{
		CourierId: courierId,
		Status:    newStatus.String(),
		Time:      updatedAt,
	}

	result := db.MariaDbClient().Create(&courierLog)
	if result.Error != nil {
		logger.Error("failed to insert courier log", result.Error)
		return result.Error
	}

	return nil
}
