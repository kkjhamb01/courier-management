package business

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/logger"
	pb "github.com/kkjhamb01/courier-management/party/proto"
)

func (s *Service) OpenGetCourierAccount(ctx context.Context, in *pb.OpenGetCourierAccountRequest) (*pb.GetCourierAccountResponse, error) {
	logger.Infof("OpenGetCourierAccount userId = %v", in.GetUserId())
	return s.getCourierAccountById(ctx, in.GetUserId())
}

func (s *Service) OpenGetProfileAdditionalInfo(ctx context.Context, in *pb.OpenGetProfileAdditionalInfoRequest) (*pb.GetProfileAdditionalInfoResponse, error) {
	logger.Infof("OpenGetProfileAdditionalInfo userId = %v, type = %v", in.GetUserId(), in.GetType())
	return s.getProfileAdditionalInfoByUserid(in.GetUserId(), in.Type)
}

func (s *Service) OpenGetDocumentsOfUser(ctx context.Context, in *pb.OpenGetDocumentsOfUserRequest) (*pb.GetDocumentsOfUserResponse, error) {
	logger.Infof("OpenGetDocumentsOfUser userId = %v, type = %v, data_type = %v", in.GetUserId(), in.GetType(), in.GetDataType())
	return s.getDocumentsOfUserById(in.GetUserId(), in.GetType(), in.GetDataType())
}

func (s *Service) OpenGetUserAccount(ctx context.Context, in *pb.OpenGetUserAccountRequest) (*pb.GetUserAccountResponse, error) {
	logger.Infof("OpenGetUserAccount userId = %v", in.GetUserId())
	return s.getUserAccountById(ctx, in.GetUserId())
}

func (s *Service) OpenGetCourierPublicInfo(ctx context.Context, in *pb.OpenGetCourierPublicInfoRequest) (*pb.OpenGetCourierPublicInfoResponse, error) {
	logger.Infof("OpenGetCourierPublicInfo courierId = %v", in.GetCourierId())

	var profile *pb.CourierProfile
	var registrationNumber string
	var profilePicture *pb.ProfilePicture

	res, err := s.getCourierAccountById(ctx, in.GetCourierId())
	if err != nil {
		logger.Errorf("OpenGetCourierPublicInfo error in get profile %v", err)
		return nil, err
	}
	if res != nil {
		profile = res.GetProfile()
	}

	mot, err := s.getProfileAdditionalInfoByUserid(in.GetCourierId(), pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_MOT)
	if err != nil {
		logger.Errorf("OpenGetCourierPublicInfo error in get mot %v", err)
	}
	if mot != nil && mot.GetMot() != nil {
		registrationNumber = mot.GetMot().GetRegistrationNumber()
	}

	pp, err := s.getProfileAdditionalInfoByUserid(in.GetCourierId(), pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_PROFILE_PICTURE)
	if err != nil {
		logger.Errorf("OpenGetCourierPublicInfo error in get profile picture %v", err)
	}
	if pp != nil && pp.GetProfilePicture() != nil {
		profilePicture = pp.GetProfilePicture()
	}

	return &pb.OpenGetCourierPublicInfoResponse{
		Profile:            profile,
		RegistrationNumber: registrationNumber,
		ProfilePicture:     profilePicture,
	}, nil
}
