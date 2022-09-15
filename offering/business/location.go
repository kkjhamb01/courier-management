package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	commonPb "gitlab.artin.ai/backend/courier-management/grpc/common/go"
	"gitlab.artin.ai/backend/courier-management/offering/storage"
)

func SetCourierLocation(ctx context.Context, location commonPb.Location, accessToken string, courierId string) (returnErr error) {

	logger.Infof("SetCourierLocation courierId = %v, location = %+v", courierId, location)

	tx := storage.CreateTx()

	defer func() {
		if returnErr != nil {
			rollbackErr := tx.Rollback()
			logger.Error("failed to rollback tx", rollbackErr)
			return
		}

		returnErr = tx.Commit(ctx)
		if returnErr != nil {
			logger.Error("failed to commit tx", returnErr)
		}
	}()

	if err := location.Validate(); err != nil {
		logger.Error("location is not valid", err)
		return err
	}

	courierType, err := GetCourierType(ctx, accessToken, courierId)
	if err != nil {
		logger.Error("failed to get courier type", err)
		return err
	}

	if err := tx.SetCourierLocationByCourierType(ctx, location, courierType, courierId); err != nil {
		logger.Error("failed to store courier location", err)
		return err
	}

	return nil
}

func GetCourierLocation(ctx context.Context, courierId string) (_ commonPb.Location, returnErr error) {

	//logger.Infof("GetCourierLocation courierId = %v", courierId)

	tx := storage.CreateTx()

	defer func() {
		if returnErr != nil {
			rollbackErr := tx.Rollback()
			logger.Error("failed to rollback tx", rollbackErr)
			return
		}

		returnErr = tx.Commit(ctx)
		if returnErr != nil {
			logger.Error("failed to commit tx", returnErr)
		}
	}()

	location, err := tx.GetCourierLocation(ctx, courierId)
	if err != nil {
		logger.Error("failed to fetch courier location", err)
		return commonPb.Location{}, err
	}

	return location, nil
}

func GetNearbyCouriers(ctx context.Context, location commonPb.Location, radiusMeter int, courierType commonPb.VehicleType) (_ []commonPb.CourierLocation, returnErr error) {

	logger.Infof("GetNearbyCouriers radiusMeter = %v, courierType = %v, location = %+v", radiusMeter, courierType, location)

	tx := storage.CreateTx()

	defer func() {
		if returnErr != nil {
			rollbackErr := tx.Rollback()
			logger.Error("failed to rollback tx", rollbackErr)
			return
		}

		returnErr = tx.Commit(ctx)
		if returnErr != nil {
			logger.Error("failed to commit tx", returnErr)
		}
	}()

	courierLocations, err := tx.GetNearbyCouriersByCourierType(ctx, location, radiusMeter, courierType)
	if err != nil {
		logger.Error("failed to fetch nearby courierLocations", err)
		return nil, err
	}

	return courierLocations, nil
}
