package business

import (
	"context"

	"github.com/kkjhamb01/courier-management/common/logger"
	pb "github.com/kkjhamb01/courier-management/party/proto"
	"github.com/kkjhamb01/courier-management/uaa/proto"
)

func (s *Service) InterServiceGetProfileAdditionalInfo(ctx context.Context, in *pb.InterServiceGetProfileAdditionalInfoRequest) (*pb.GetProfileAdditionalInfoResponse, error) {
	logger.Infof("InterServiceGetProfileAdditionalInfo userId = %v, type = %v", in.GetUserId(), in.GetType())
	return s.getProfileAdditionalInfoByUserid(in.GetUserId(), in.Type)
}

func (s *Service) InterServiceUpdateProfileAdditionalInfo(ctx context.Context, in *pb.InterServiceUpdateProfileAdditionalInfoRequest) (*pb.UpdateProfileAdditionalInfoResponse, error) {
	logger.Infof("InterServiceUpdateProfileAdditionalInfo userId = %v, info = %v", in.GetUserId(), in.GetInfo())
	var in2 *pb.UpdateProfileAdditionalInfoRequest
	if in.GetIdCard() != nil {
		in2 = &pb.UpdateProfileAdditionalInfoRequest{
			Info: &pb.UpdateProfileAdditionalInfoRequest_IdCard{
				IdCard: in.GetIdCard(),
			},
		}
	} else if in.GetDrivingLicense() != nil {
		in2 = &pb.UpdateProfileAdditionalInfoRequest{
			Info: &pb.UpdateProfileAdditionalInfoRequest_DrivingLicense{
				DrivingLicense: in.GetDrivingLicense(),
			},
		}
	} else if in.GetDriverBackground() != nil {
		in2 = &pb.UpdateProfileAdditionalInfoRequest{
			Info: &pb.UpdateProfileAdditionalInfoRequest_DriverBackground{
				DriverBackground: in.GetDriverBackground(),
			},
		}
	} else if in.GetResidenceCard() != nil {
		in2 = &pb.UpdateProfileAdditionalInfoRequest{
			Info: &pb.UpdateProfileAdditionalInfoRequest_ResidenceCard{
				ResidenceCard: in.GetResidenceCard(),
			},
		}
	} else if in.GetBankAccount() != nil {
		in2 = &pb.UpdateProfileAdditionalInfoRequest{
			Info: &pb.UpdateProfileAdditionalInfoRequest_BankAccount{
				BankAccount: in.GetBankAccount(),
			},
		}
	} else if in.GetAddress() != nil {
		in2 = &pb.UpdateProfileAdditionalInfoRequest{
			Info: &pb.UpdateProfileAdditionalInfoRequest_Address{
				Address: in.GetAddress(),
			},
		}
	} else if in.GetMot() != nil {
		in2 = &pb.UpdateProfileAdditionalInfoRequest{
			Info: &pb.UpdateProfileAdditionalInfoRequest_Mot{
				Mot: in.GetMot(),
			},
		}
	}
	return s.updateProfileAdditionalInfoByUserId(ctx, in.GetUserId(), in2)
}

func (s *Service) InterServiceFindCourierAccounts(ctx context.Context, in *pb.InterServiceFindCourierAccountsRequest) (*pb.FindCourierAccountsResponse, error) {
	logger.Infof("InterServiceFindCourierAccounts userId = %v, phoneNumber = %v, name = %v, email = %v", in.GetUserId(), in.GetPhoneNumber(), in.GetName(), in.GetEmail())
	var in2 *pb.FindCourierAccountsRequest
	if in.GetUserId() != "" {
		in2 = &pb.FindCourierAccountsRequest{
			Filter: &pb.FindCourierAccountsRequest_UserId{
				UserId: in.GetUserId(),
			},
		}
	} else if in.GetPhoneNumber() != "" {
		in2 = &pb.FindCourierAccountsRequest{
			Filter: &pb.FindCourierAccountsRequest_PhoneNumber{
				PhoneNumber: in.GetPhoneNumber(),
			},
		}
	} else if in.GetName() != "" {
		in2 = &pb.FindCourierAccountsRequest{
			Filter: &pb.FindCourierAccountsRequest_Name{
				Name: in.GetName(),
			},
		}
	} else if in.GetEmail() != "" {
		in2 = &pb.FindCourierAccountsRequest{
			Filter: &pb.FindCourierAccountsRequest_Email{
				Email: in.GetEmail(),
			},
		}
	} else {
		return nil, proto.InvalidArgument.ErrorMsg("invalid request")
	}
	return s.FindCourierAccounts(ctx, in2)
}

func (s *Service) InterServiceGetDocumentsOfUser(ctx context.Context, in *pb.InterServiceGetDocumentsOfUserRequest) (*pb.GetDocumentsOfUserResponse, error) {
	logger.Infof("InterServiceFindCourierAccounts userId = %v, type = %v, dataType = %v", in.GetUserId(), in.GetType(), in.GetDataType())
	return s.getDocumentsOfUserById(in.GetUserId(), in.GetType(), in.GetDataType())
}
