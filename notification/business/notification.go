package business

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/notification/domain"
	pb "github.com/kkjhamb01/courier-management/notification/proto"
	proto "github.com/kkjhamb01/courier-management/uaa/proto"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Service) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("Register user_id = %v, device_id = %v, device_model = %v, device_os = %v, device_token = %v, device_version = %v, manufacturer = %v",
		tokenUser.Id, in.GetDeviceId(), in.GetDeviceModel(), in.GetDeviceOs(),
		in.GetDeviceToken(), in.GetDeviceVersion(), in.GetManufacturer())

	var phoneNumber string

	for _, claim := range tokenUser.Claims {
		if claim.ClaimType == security.CLAIM_TYPE_PHONE_NUMBER {
			phoneNumber = claim.Identifier
		}
	}

	if phoneNumber == "" {
		logger.Debugf("Register request doesn't have phone_number")
		return nil, proto.Unauthenticated.ErrorMsg("unauthorized phone number")
	}

	logger.Debugf("Register request with phone_number %v", phoneNumber)

	err := s.db.Transaction(func(tx *gorm.DB) error {
		return tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "device_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"phone_number", "manufacturer", "device_model", "device_os", "device_version", "device_token"}),
		}).Create(&domain.Device{
			DeviceId:      in.DeviceId,
			PhoneNumber:   phoneNumber,
			Manufacturer:  in.Manufacturer,
			DeviceModel:   in.DeviceModel,
			DeviceOs:      int32(in.DeviceOs),
			DeviceVersion: in.DeviceVersion,
			DeviceToken:   in.DeviceToken,
		}).Error
	})

	if err != nil {
		logger.Errorf("Register error in registering device", err)
		return nil, proto.Internal.Error(err)
	}

	return &pb.RegisterResponse{}, nil

}

func (s *Service) Unregister(ctx context.Context, in *pb.UnregisterRequest) (*pb.UnregisterResponse, error) {

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("Unregister user_id = %v, device_id = %v", tokenUser.Id, in.GetDeviceId())

	var phoneNumber string

	for _, claim := range tokenUser.Claims {
		if claim.ClaimType == security.CLAIM_TYPE_PHONE_NUMBER {
			phoneNumber = claim.Identifier
		}
	}

	if phoneNumber == "" {
		logger.Debugf("Unregister request doesn't have phone_number")
		return nil, proto.Unauthenticated.ErrorMsg("unauthorized phone number")
	}

	logger.Debugf("Unregister request with phone_number %v", phoneNumber)

	err := s.db.Transaction(func(tx *gorm.DB) error {
		return tx.Delete(&domain.Device{
			DeviceId: in.DeviceId,
		}).Error
	})

	if err != nil {
		logger.Errorf("Unregister cannot perform action", err)
		return nil, proto.Internal.Error(err)
	}

	return &pb.UnregisterResponse{}, nil

}
