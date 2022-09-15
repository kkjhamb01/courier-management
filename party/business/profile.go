package business

import (
	"context"
	"errors"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/party/domain"
	pb "gitlab.artin.ai/backend/courier-management/party/proto"
	"gitlab.artin.ai/backend/courier-management/uaa/proto"
)

type hasDocument interface {
	GetDocumentIds() []*pb.DocumentInfo
	SetDocumentIds([]*pb.DocumentInfo)
	GetDocumentInfoType() int32
}

func (s *Service) RegisterClaim(ctx context.Context, in *pb.RegisterClaimRequest) (*pb.RegisterClaimResponse, error){
	logger.Debugf("RegisterClaim type = %v, userId = %v, claim = %v, identifier = %v",
		in.GetType(), in.GetUserId(), in.GetClaim(), in.GetIdentifier())

	var err error

	if in.Type==pb.UserType_USER_TYPE_CURIOUR {
		claim := domain.CourierClaim{
			UserID: in.UserId,
			ClaimType: int(in.GetClaim()),
			Identifier: in.GetIdentifier(),
		}
		err = s.db.Model(&domain.CourierClaim{}).Create(claim).Error
	} else if in.Type==pb.UserType_USER_TYPE_PASSENGER {
		claim := domain.ClientClaim{
			UserID: in.UserId,
			ClaimType: int(in.GetClaim()),
			Identifier: in.GetIdentifier(),
		}
		err = s.db.Model(&domain.ClientClaim{}).Create(claim).Error
	}

	return &pb.RegisterClaimResponse{
	}, err
}

func (s *Service) FindAccount(ctx context.Context, in *pb.FindAccountRequest) (*pb.FindAccountResponse, error) {
	logger.Debugf("FindAccount type = %v, userId = %v, email = %v, phoneNumber = %v, googleId = %v, facebookId = %v",
		in.GetType(), in.GetUserId(), in.GetEmail(), in.GetPhoneNumber(), in.GetGoogleId(), in.GetFacebookId())

	var err error

	if in.Type == pb.UserType_USER_TYPE_CURIOUR{
		var users []domain.CourierUser

		if in.GetUserId() != ""{
			err = s.db.Model(&domain.CourierUser{}).Where("id = ?", in.GetUserId()).Scan(&users).Error
		} else {
			var identifier string
			if in.GetPhoneNumber() != "" {
				identifier = in.GetPhoneNumber()
			} else if in.GetEmail() != "" {
				identifier = in.GetEmail()
			} else if in.GetFacebookId() != ""{
				identifier = in.GetFacebookId()
			} else if in.GetGoogleId() != ""{
				identifier = in.GetGoogleId()
			}
			err = s.db.Model(&domain.CourierUser{}).Where("id IN (SELECT user_id FROM courier_claim WHERE identifier=?)", identifier).Scan(&users).Error
		}

		if err != nil{
			return nil,err
		}

		if len(users) == 0{
			return nil, proto.NotFound.ErrorMsg("no user found")
		}

		var accounts = make([]*pb.Account, len(users))

		for i,user := range users {
			var claims []domain.CourierClaim
			err = s.db.Model(&domain.CourierClaim{}).Where("user_id = ?", user.ID).Scan(&claims).Error
			if err != nil{
				return nil,err
			}
			var aclaims []*pb.Claim
			for _, claim := range claims{
				aclaims = append(aclaims, &pb.Claim{
					ClaimType: pb.ClaimType(claim.ClaimType),
					Identifier: claim.Identifier,
				})
			}
			accounts[i] = &pb.Account{
				UserId: user.ID,
				PhoneNumber: user.PhoneNumber,
				FirstName: user.FirstName.String,
				LastName: user.LastName,
				Email: user.Email.String,
				Claims: aclaims,
			}
		}

		logger.Debugf("FindAccount result = %v", accounts)

		return &pb.FindAccountResponse{
			Users: accounts,
		}, nil
	} else if in.Type == pb.UserType_USER_TYPE_PASSENGER{
		var users []domain.ClientUser

		if in.GetUserId() != ""{
			err = s.db.Model(&domain.ClientUser{}).Where("id = ?", in.GetUserId()).Scan(&users).Error
		} else {
			var identifier string
			if in.GetPhoneNumber() != "" {
				identifier = in.GetPhoneNumber()
			} else if in.GetEmail() != "" {
				identifier = in.GetEmail()
			} else if in.GetFacebookId() != ""{
				identifier = in.GetFacebookId()
			} else if in.GetGoogleId() != ""{
				identifier = in.GetGoogleId()
			}
			err = s.db.Model(&domain.ClientUser{}).Where("id IN (SELECT user_id FROM client_claim WHERE identifier=?)", identifier).Scan(&users).Error
		}

		if err != nil{
			return nil,err
		}

		if len(users) == 0{
			return nil, proto.NotFound.ErrorMsg("no user found")
		}

		var accounts = make([]*pb.Account, len(users))

		for i,user := range users {
			var claims []domain.ClientClaim
			err = s.db.Model(&domain.ClientClaim{}).Where("user_id = ?", user.ID).Scan(&claims).Error
			if err != nil{
				return nil,err
			}
			var aclaims []*pb.Claim
			for _, claim := range claims{
				aclaims = append(aclaims, &pb.Claim{
					ClaimType: pb.ClaimType(claim.ClaimType),
					Identifier: claim.Identifier,
				})
			}
			accounts[i] = &pb.Account{
				UserId: user.ID,
				PhoneNumber: user.PhoneNumber,
				FirstName: user.FirstName.String,
				LastName: user.LastName,
				Email: user.Email.String,
				Claims: aclaims,
			}
		}

		logger.Debugf("FindAccount result = %v", accounts)

		return &pb.FindAccountResponse{
			Users: accounts,
		}, nil
	} else if in.Type == pb.UserType_USER_TYPE_ADMIN{
		var users []domain.AdminUser

		if in.GetUserId() != ""{
			err = s.db.Model(&domain.AdminUser{}).Where("id = ?", in.GetUserId()).Scan(&users).Error
		} else {
			if in.GetUsername() == "" {
				return nil, errors.New("invalid query type")
			}
			err = s.db.Model(&domain.AdminUser{}).Where("username = ?", in.GetUsername()).Scan(&users).Error
		}

		logger.Debugf("FindAccount fetched admin users = %v", users)

		if err != nil{
			return nil,err
		}

		if len(users) == 0{
			return nil, proto.NotFound.ErrorMsg("no user found")
		}

		var accounts = make([]*pb.Account, len(users))

		for i,user := range users {
			var aclaims []*pb.Claim
			aclaims = append(aclaims, &pb.Claim{
				ClaimType:  pb.ClaimType_CLAIM_TYPE_USERNAME,
				Identifier: user.Username,
			})
			aclaims = append(aclaims, &pb.Claim{
				ClaimType:  pb.ClaimType_CLAIM_TYPE_PASSWORD,
				Identifier: user.Password,
			})
			accounts[i] = &pb.Account{
				UserId: user.UserID,
				FirstName: user.FirstName.String,
				LastName: user.LastName.String,
				Claims: aclaims,
			}
		}

		logger.Debugf("FindAccount result = %v", accounts)

		return &pb.FindAccountResponse{
			Users: accounts,
		}, nil
	} else if in.Type == pb.UserType_USER_TYPE_ALL{
		in.Type = pb.UserType_USER_TYPE_PASSENGER
		res,err := s.FindAccount(ctx, in)
		if err == nil{
			return res, nil
		}
		in.Type = pb.UserType_USER_TYPE_CURIOUR
		return s.FindAccount(ctx, in)
	}

	return nil, nil

}

func (s *Service) GetCourier(id string) (*domain.CourierUser, error) {
	logger.Debugf("GetCourier id = %v ", id)

	user := domain.CourierUser{}

	if err := s.db.Model(&domain.CourierUser{}).Where("id = ?", id).Find(&user).Error; err != nil{
		logger.Debugf("GetCourier error in performing query %v, error = %v ", id, err)
		return nil, proto.Internal.Error(err)
	}
	if user.ID == ""{
		return nil, errors.New("user not found")
	}

	logger.Debugf("GetCourier result = %v ", user)

	return &user, nil
}

func (s *Service) GetUser(id string) (*domain.ClientUser, error) {
	logger.Debugf("GetUser id = %v ", id)

	user := domain.ClientUser{}

	if err := s.db.Model(&domain.ClientUser{}).Where("id = ?", id).Find(&user).Error; err != nil{
		return nil, proto.Internal.Error(err)
	}

	if user.ID == ""{
		return nil, errors.New("user not found")
	}

	logger.Debugf("GetUser result = %v ", user)

	return &user, nil
}