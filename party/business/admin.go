package business

import (
	"context"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/party/domain"
	pb "gitlab.artin.ai/backend/courier-management/party/proto"
	"gitlab.artin.ai/backend/courier-management/uaa/proto"
)

func (s *Service) ServiceGetCourierRegistrationStat(ctx context.Context, in *pb.ServiceGetCourierRegistrationStatRequest) (*pb.GetRegistrationStatResponse, error){
	logger.Infof("ServiceGetCourierRegistrationStat %v ", in)
	return nil, nil
}

func (s *Service) ServiceGetClientRegistrationStat(ctx context.Context, in *pb.ServiceGetClientRegistrationStatRequest) (*pb.GetRegistrationStatResponse, error){
	logger.Infof("ServiceGetClientRegistrationStat %v ", in)
	return nil, nil
}

func (s *Service) ServiceGetCourierRegistration(ctx context.Context, in *pb.ServiceGetCourierRegistrationRequest) (*pb.GetCourierRegistrationResponse, error){
	logger.Infof("ServiceGetCourierRegistration userId = %v, email = %v, name = %v, phone number = %v, pagination = %v",
		in.GetUserId(), in.GetEmail(),
		in.GetName(), in.GetPhoneNumber(), in.GetPagination())

	var users []domain.CourierUser

	var err error

	if in.GetUserId() != ""{
		err = s.db.Model(&domain.CourierUser{}).Where("id = ?", in.GetUserId()).Scan(&users).Error
	} else if in.GetName() != ""{
		err = s.db.Model(&domain.CourierUser{}).Where("first_name LIKE ? OR last_name LIKE ?", "%" + in.GetName() + "%", "%" + in.GetName() + "%").Scan(&users).Error
	} else if in.GetPhoneNumber() != "" || in.GetEmail() != "" {
		var identifier string
		if in.GetPhoneNumber() != "" {
			identifier = in.GetPhoneNumber()
		} else if in.GetEmail() != "" {
			identifier = in.GetEmail()
		}
		err = s.db.Model(&domain.CourierUser{}).Where("id IN (SELECT user_id FROM courier_claim WHERE identifier=?)", identifier).Scan(&users).Error
	} else {
		var offset, limit, page int32
		var order = ""
		if in.GetPagination() != nil {
			page = in.GetPagination().GetPage()
			limit = in.GetPagination().GetLimit()
			if in.GetPagination().GetSort() != ""{
				order = in.GetPagination().GetSort()
				if in.GetPagination().GetSortType() == pb.SortType_SORT_TYPE_DESC {
					order = order + " desc"
				} else{
					order = order + " asc"
				}
			}
		}
		if page == 0 {
			page = 1
		}
		if limit == 0 {
			limit = 20
		}
		offset = (page - 1) * limit
		if order == ""{
			order = "creation_time desc"
		}
		err = s.db.Limit(int(limit)).Offset(int(offset)).Order(order).Model(&domain.CourierUser{}).Scan(&users).Error
	}

	if err != nil{
		return nil, proto.Internal.Error(err)
	}

	if len(users) == 0{
		return nil, proto.NotFound.ErrorMsg("no user found")
	}

	var profiles = make([]*pb.CourierProfileAndStatus, len(users))

	for i,user := range users {
		var citizen bool
		if user.Citizen.Int32 > 0{
			citizen = true
		}
		var statusItemsConv []*pb.ProfileAdditionalInfoStatusItem
		statusItems, err := s.getProfileAdditionalInfoStatusByUserId(user.ID)
		if err != nil{
			return nil, err
		}
		for _, statusItem := range statusItems.GetItems(){
			statusItemsConv = append(statusItemsConv, &pb.ProfileAdditionalInfoStatusItem{
				Type: statusItem.Type,
				Status: statusItem.Status,
				Message: statusItem.Message,
			})
		}
		var citizen2 pb.Boolean
		if citizen{
			citizen2 = pb.Boolean_BOOLEAN_TRUE
		} else {
			citizen2 = pb.Boolean_BOOLEAN_FALSE
		}
		profiles[i] = &pb.CourierProfileAndStatus{
			Profile: &pb.CourierProfile{
				UserId: user.ID,
				PhoneNumber: user.PhoneNumber,
				FirstName: user.FirstName.String,
				LastName: user.LastName,
				Email: user.Email.String,
				BirthDate: user.BirthDate.String,
				TransportType: pb.TransportationType(user.TransportType.Int32),
				TransportSize: pb.TransportationSize(user.TransportSize.Int32),
				Citizen: citizen2,
			},
			UserStatus: pb.UserStatus(user.Status),
			StatusItems: statusItemsConv,
		}
	}

	logger.Debugf("ServiceGetCourierRegistration result = %v", profiles)

	return &pb.GetCourierRegistrationResponse{
		Items: profiles,
	}, nil
}

func (s *Service) ServiceGetClientRegistration(ctx context.Context, in *pb.ServiceGetClientRegistrationRequest) (*pb.GetClientRegistrationResponse, error){
	logger.Infof("ServiceGetClientRegistration userId = %v, email = %v, name = %v, phone number = %v, pagination = %v",
		in.GetUserId(), in.GetEmail(),
		in.GetName(), in.GetPhoneNumber(), in.GetPagination())

	var users []domain.ClientUser

	var err error

	if in.GetUserId() != ""{
		err = s.db.Model(&domain.ClientUser{}).Where("id = ?", in.GetUserId()).Scan(&users).Error
	} else if in.GetName() != ""{
		err = s.db.Model(&domain.ClientUser{}).Where("first_name LIKE ? OR last_name LIKE ?", "%" + in.GetName() + "%", "%" + in.GetName() + "%").Scan(&users).Error
	} else if in.GetPhoneNumber() != "" || in.GetEmail() != "" {
		var identifier string
		if in.GetPhoneNumber() != "" {
			identifier = in.GetPhoneNumber()
		} else if in.GetEmail() != "" {
			identifier = in.GetEmail()
		}
		err = s.db.Model(&domain.ClientUser{}).Where("id IN (SELECT user_id FROM client_claim WHERE identifier=?)", identifier).Scan(&users).Error
	} else {
		var offset, limit, page int32
		var order = ""
		if in.GetPagination() != nil {
			page = in.GetPagination().GetPage()
			limit = in.GetPagination().GetLimit()
			if in.GetPagination().GetSort() != ""{
				order = in.GetPagination().GetSort()
				if in.GetPagination().GetSortType() == pb.SortType_SORT_TYPE_DESC {
					order = order + " desc"
				} else{
					order = order + " asc"
				}
			}
		}
		if page == 0 {
			page = 1
		}
		if limit == 0 {
			limit = 20
		}
		offset = (page - 1) * limit
		if order == ""{
			order = "creation_time desc"
		}
		err = s.db.Limit(int(limit)).Offset(int(offset)).Order(order).Model(&domain.ClientUser{}).Scan(&users).Error
	}

	if err != nil{
		return nil, proto.Internal.Error(err)
	}

	if len(users) == 0{
		return nil, proto.NotFound.ErrorMsg("no user found")
	}

	var profiles = make([]*pb.UserProfile, len(users))

	for i,user := range users {
		profiles[i] = &pb.UserProfile{
			UserId: user.ID,
			PhoneNumber: user.PhoneNumber,
			FirstName: user.FirstName.String,
			LastName: user.LastName,
			Email: user.Email.String,
			PaymentMethod: pb.PaymentMethod(user.PaymentMethod.Int32),
			Code: user.Referral,
		}
	}

	logger.Debugf("ServiceGetClientRegistration result = %v", profiles)

	return &pb.GetClientRegistrationResponse{
		Profiles: profiles,
	}, nil
}

func (s *Service) ServiceGetDocumentsOfUser(ctx context.Context, in *pb.ServiceGetDocumentsOfUserRequest) (*pb.GetDocumentsOfUserResponse, error){
	logger.Infof("ServiceGetDocumentsOfUser userId = %v, type = %v", in.GetUserId(), in.GetType())
	return s.getDocumentsOfUserById(in.GetUserId(), in.GetType(), pb.DocumentDataType_UNKNOWN_DOCUMENT_DATA_TYPE)
}

func (s *Service) ServiceGetProfileAdditionalInfo(ctx context.Context, in *pb.ServiceGetProfileAdditionalInfoRequest) (*pb.GetProfileAdditionalInfoResponse, error){
	logger.Infof("ServiceGetProfileAdditionalInfo userId = %v, type = %v", in.GetUserId(), in.GetType())
	return s.getProfileAdditionalInfoByUserid(in.UserId, in.Type)
}

func (s *Service) ServiceUpdateProfileStatus(ctx context.Context, in *pb.ServiceUpdateProfileStatusRequest) (*pb.UpdateProfileAdditionalInfoStatusResponse, error){
	logger.Infof("ServiceUpdateProfileStatus userId = %v, userStatus = %v, statusItems = %v",
		in.GetUserId(), in.GetUserStatus(), in.GetStatusItems())

	if in.GetStatusItems() != nil && len(in.GetStatusItems()) > 0{
		for _,item := range in.GetStatusItems(){
			var status pb.UpdateProfileAdditionalInfoStatus
			if item.GetStatus() == pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_ACCEPTED{
				status = pb.UpdateProfileAdditionalInfoStatus_UPDATE_PROFILE_ADDITIONAL_INFO_STATUS_ACCEPTED
			} else if item.GetStatus() == pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_REJECTED{
				status = pb.UpdateProfileAdditionalInfoStatus_UPDATE_PROFILE_ADDITIONAL_INFO_STATUS_REJECTED
			} else {
				return nil, proto.InvalidArgument.ErrorMsg("invalid status")
			}
			_,err := s.updateProfileAdditionalInfoStatusById(in.GetUserId(), item.GetType(), status, item.GetMessage())
			if err != nil{
				return nil, err
			}
		}
	}

	if in.GetUserStatus() != pb.UserStatus_UNKNOWN_USER_STATUS{
		_,err := s.UpdateCourierUserStatus(ctx, &pb.UpdateCourierUserStatusRequest{
			UserId: in.GetUserId(),
			Status: in.GetUserStatus(),
		})
		if err != nil{
			return nil, err
		}
	}

	return &pb.UpdateProfileAdditionalInfoStatusResponse{

	}, nil
}