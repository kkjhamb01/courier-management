package business

import (
	"context"
	"database/sql"

	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/party/domain"
	pb "github.com/kkjhamb01/courier-management/party/proto"
	"github.com/kkjhamb01/courier-management/uaa/proto"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Service) CreateUserAccount(ctx context.Context, in *pb.CreateUserAccountRequest) (*pb.CreateUserAccountResponse, error) {
	logger.Infof("CreateUserAccount firstName = %v, lastName = %v, email = %v, birthDate = %v",
		in.GetFirstName(), in.GetLastName(), in.GetEmail(), in.GetBirthDate())

	tokenUser := ctx.Value("user").(security.User)

	existingUser := domain.ClientUser{}

	if err := s.db.Model(&domain.ClientUser{}).Where("id = ?", tokenUser.Id).Find(&existingUser).Error; err != nil {
		logger.Debugf("CreateUserAccount userId = %v, error = %v", tokenUser.Id, err)
		return nil, proto.Internal.Error(err)
	}

	if existingUser.ID != "" {
		logger.Debugf("CreateUserAccount user already exists %v", tokenUser.Id)
		return nil, proto.AlreadyExists.ErrorMsg("user already exists")
	}

	var authorizedPhoneNumber = false
	for _, claim := range tokenUser.Claims {
		if claim.ClaimType == security.CLAIM_TYPE_PHONE_NUMBER {
			authorizedPhoneNumber = true
		}
	}

	if !authorizedPhoneNumber {
		logger.Debugf("CreateUserAccount token doesn't have any authorized phone number")
		return nil, proto.Unauthenticated.ErrorMsg("token doesn't have any authorized phone number")
	}

	var referred domain.ClientUser

	if in.GetReferral() != "" {
		err := s.db.Model(&domain.ClientUser{}).Where("referral = ?", in.GetReferral()).Find(&referred).Error

		if err != nil {
			logger.Errorf("CreateUserAccount error in query referral", err)
			return nil, proto.Internal.Error(err)
		}

		if referred.ID == "" {
			logger.Debugf("CreateUserAccount referral not found")
			return nil, proto.InvalidArgument.ErrorMsg("referral not found")
		}
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// insert user
		if err := tx.Create(&domain.ClientUser{
			ID:          tokenUser.Id,
			FirstName:   sql.NullString{String: in.FirstName, Valid: in.FirstName != ""},
			LastName:    in.LastName,
			Email:       sql.NullString{String: in.Email, Valid: in.Email != ""},
			PhoneNumber: tokenUser.PhoneNumber,
			Status:      int32(pb.UserStatus_USER_STATUS_AVAILABLE),
			BirthDate:   sql.NullString{String: in.BirthDate, Valid: in.BirthDate != ""},
			Referral:    s.CalculateReferral(tokenUser.Id),
		}).Error; err != nil {
			logger.Errorf("CreateUserAccount cannot perform update", err)
			return proto.Internal.Error(err)
		}

		// insert claims
		var claims = make([]*domain.ClientClaim, len(tokenUser.Claims))
		for i, claim := range tokenUser.Claims {
			claims[i] = &domain.ClientClaim{
				UserID:     tokenUser.Id,
				ClaimType:  int(claim.ClaimType),
				Identifier: claim.Identifier,
			}
		}
		if err := tx.CreateInBatches(claims, len(claims)).Error; err != nil {
			return proto.Internal.Error(err)
		}

		return nil
	})

	if err != nil {
		logger.Errorf("CreateUserAccount error in creating profile", err)
		return nil, proto.Internal.Error(err)
	}

	if in.GetReferral() != "" {
		err = s.promotionApi.AssignUserReferral(tokenUser.Id, in.GetReferral(), referred.ID)

		if err != nil {
			logger.Errorf("CreateUserAccount error in connection to promotion service", err)
			return nil, proto.Internal.Error(err)
		}
	}

	return &pb.CreateUserAccountResponse{}, nil
}

func (s *Service) GetUserAccount(ctx context.Context, in *pb.GetUserAccountRequest) (*pb.GetUserAccountResponse, error) {

	tokenUser := ctx.Value("user").(security.User)

	logger.Debugf("GetUserAccount userId = %v", tokenUser.Id)

	return s.getUserAccountById(ctx, tokenUser.Id)
}

func (s *Service) getUserAccountById(ctx context.Context, userId string) (*pb.GetUserAccountResponse, error) {

	logger.Debugf("getUserAccountById userId = %v", userId)

	var user domain.ClientUser

	err := s.db.Model(&domain.ClientUser{}).Where("id = ?", userId).Find(&user).Error

	if err != nil {
		logger.Debugf("GetUserAccount user = %v, error = %v", user, err)
		return nil, proto.Internal.Error(err)
	}

	if user.ID == "" {
		logger.Debugf("GetUserAccount user not found %v", user)
		return nil, proto.NotFound.ErrorMsg("user not found")
	}

	logger.Debugf("GetUserAccount user = %v", user)

	authorizedClaims, err := s.getAuthorizedClientClaims(userId)

	if err != nil {
		logger.Infof("GetUserAccount error in get authorized claims = %v", err)
		return nil, proto.Internal.Error(err)
	}

	return &pb.GetUserAccountResponse{
		Profile: &pb.UserProfile{
			UserId:           user.ID,
			PhoneNumber:      user.PhoneNumber,
			FirstName:        user.FirstName.String,
			LastName:         user.LastName,
			Email:            user.Email.String,
			PaymentMethod:    pb.PaymentMethod(user.PaymentMethod.Int32),
			BirthDate:        user.BirthDate.String,
			Code:             user.Referral,
			AuthorizedClaims: authorizedClaims,
		},
	}, nil
}

func (s *Service) FindUserAccounts(ctx context.Context, in *pb.FindUserAccountsRequest) (*pb.FindUserAccountsResponse, error) {
	logger.Debugf("FindUserAccounts userId = %v, email = %v, name = %v, phoneNumber = %v ",
		in.GetUserId(), in.GetEmail(), in.GetName(), in.GetPhoneNumber())

	var users []domain.ClientUser

	var err error

	if in.GetUserId() != "" {
		err = s.db.Model(&domain.ClientUser{}).Where("id = ?", in.GetUserId()).Scan(&users).Error
	} else if in.GetName() != "" {
		err = s.db.Model(&domain.ClientUser{}).Where("first_name LIKE ? OR last_name LIKE ?", "%"+in.GetName()+"%", "%"+in.GetName()+"%").Scan(&users).Error
	} else {
		var identifier string
		if in.GetPhoneNumber() != "" {
			identifier = in.GetPhoneNumber()
		} else if in.GetEmail() != "" {
			identifier = in.GetEmail()
		}
		err = s.db.Model(&domain.ClientUser{}).Where("id IN (SELECT user_id FROM client_claim WHERE identifier=?)", identifier).Scan(&users).Error
	}

	if err != nil {
		return nil, proto.Internal.Error(err)
	}

	if len(users) == 0 {
		return nil, proto.NotFound.ErrorMsg("no user found")
	}

	var profiles = make([]*pb.UserProfile, len(users))

	for i, user := range users {

		authorizedClaims, err := s.getAuthorizedClientClaims(user.ID)

		if err != nil {
			logger.Infof("FindUserAccounts error in get authorized claims = %v", err)
			return nil, proto.Internal.Error(err)
		}

		profiles[i] = &pb.UserProfile{
			UserId:           user.ID,
			PhoneNumber:      user.PhoneNumber,
			FirstName:        user.FirstName.String,
			LastName:         user.LastName,
			Email:            user.Email.String,
			PaymentMethod:    pb.PaymentMethod(user.PaymentMethod.Int32),
			BirthDate:        user.BirthDate.String,
			Code:             user.Referral,
			AuthorizedClaims: authorizedClaims,
		}
	}

	logger.Debugf("FindUserAccounts result = %v", profiles)

	return &pb.FindUserAccountsResponse{
		Profiles: profiles,
	}, nil
}

func (s *Service) UpdateUserAccount(ctx context.Context, in *pb.UpdateUserAccountRequest) (*pb.UpdateUserAccountResponse, error) {
	logger.Infof("UpdateUserAccount firstName = %v, lastName = %v, email = %v, birthDate = %v, paymentMethod = %v",
		in.GetFirstName(), in.GetLastName(), in.GetEmail(), in.GetBirthDate(), in.GetPaymentMethod())

	tokenUser := ctx.Value("user").(security.User)

	user, _ := s.GetUser(tokenUser.Id)
	if user == nil {
		logger.Debugf("UpdateUserAccount user does not exist %v", tokenUser.Id)
		return nil, proto.NotFound.ErrorMsg("user does not exist")
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// update user
		updatedUser := domain.ClientUser{
			ID: user.ID,
		}
		if in.FirstName != "" {
			updatedUser.FirstName = sql.NullString{String: in.FirstName, Valid: true}
		}

		if in.LastName != "" {
			updatedUser.LastName = in.LastName
		}

		if in.Email != "" {
			updatedUser.Email = sql.NullString{String: in.Email, Valid: true}
		}

		if in.PaymentMethod != 0 {
			updatedUser.PaymentMethod = sql.NullInt32{Int32: int32(in.PaymentMethod), Valid: true}
		}

		if in.BirthDate != "" {
			updatedUser.BirthDate = sql.NullString{String: in.BirthDate, Valid: true}
		}

		if err := tx.Model(&updatedUser).Updates(&updatedUser).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		logger.Errorf("UpdateUserAccount error in updating profile", err)
		return nil, proto.Internal.Error(err)
	}

	return &pb.UpdateUserAccountResponse{}, nil
}

func (s *Service) UpdateUserCard(ctx context.Context, in *pb.UpdateUserCardRequest) (*pb.UpdateUserCardResponse, error) {

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("UpdateUserCard userId = %v, cardNumber = %v, cvv = %v, country = %v, issueDate = %v, zipCode = %v",
		tokenUser.Id, in.GetCardNumber(), in.GetCvv(), in.GetCountry(), in.GetIssueDate(), in.GetZipCode())

	user, _ := s.GetUser(tokenUser.Id)
	if user == nil {
		logger.Debugf("UpdateUserCard user does not exist %v", tokenUser.Id)
		return nil, proto.NotFound.ErrorMsg("user does not exist")
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		return tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "card_number"}},
			DoUpdates: clause.AssignmentColumns([]string{"issue_date", "cvv", "zip_code", "country"}),
		}).Create(&domain.ClientCard{
			UserID:     tokenUser.Id,
			CardNumber: in.CardNumber,
			IssueDate:  in.IssueDate,
			CVV:        in.Cvv,
			ZipCode:    in.ZipCode,
			Country:    in.Country,
		}).Error
	})

	if err != nil {
		logger.Errorf("UpdateUserCard error in updating user card", err)
		return nil, proto.Internal.Error(err)
	}

	return &pb.UpdateUserCardResponse{}, nil
}

func (s *Service) DeleteUserCard(ctx context.Context, in *pb.DeleteUserCardRequest) (*pb.DeleteUserCardResponse, error) {

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("DeleteUserCard userId = %v, cardNumber = %v", tokenUser.Id, in.GetCardNumber())

	user, _ := s.GetUser(tokenUser.Id)
	if user == nil {
		logger.Debugf("DeleteUserCard user does not exist %v", tokenUser.Id)
		return nil, proto.NotFound.ErrorMsg("user does not exist")
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		return tx.Model(&domain.ClientCard{}).Delete(&domain.ClientCard{
			UserID:     tokenUser.Id,
			CardNumber: in.CardNumber,
		}).Error
	})

	if err != nil {
		logger.Errorf("DeleteUserCard error in updating user card", err)
		return nil, proto.Internal.Error(err)
	}

	return &pb.DeleteUserCardResponse{}, nil
}

func (s *Service) GetUserCard(ctx context.Context, in *pb.GetUserCardRequest) (*pb.GetUserCardResponse, error) {

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("GetUserCard userId = %v", tokenUser.Id)

	user, _ := s.GetUser(tokenUser.Id)
	if user == nil {
		logger.Debugf("GetUserCard user does not exist %v", tokenUser.Id)
		return nil, proto.NotFound.ErrorMsg("user does not exist")
	}

	var cards []*domain.ClientCard
	if err := s.db.Model(&domain.ClientCard{}).Where("user_id = ?", user.ID).Find(&cards).Error; err != nil {
		logger.Debugf("GetUserCard cannot perform query %v, error = %v", tokenUser.Id, err)
		return nil, proto.Internal.Error(err)
	}
	if len(cards) == 0 {
		logger.Debugf("GetUserCard empty query result %v", tokenUser.Id)
		return nil, proto.NotFound.ErrorNoMsg()
	}

	var cardsDto = make([]*pb.UserCard, len(cards))
	for i, card := range cards {
		cardsDto[i] = &pb.UserCard{
			UserId:     tokenUser.Id,
			CardNumber: card.CardNumber,
			IssueDate:  card.IssueDate,
			Cvv:        card.CVV,
			ZipCode:    card.ZipCode,
			Country:    card.Country,
		}
	}

	logger.Infof("GetUserCard result = %v", cardsDto)

	return &pb.GetUserCardResponse{
		Cards: cardsDto,
	}, nil
}

func (s *Service) UpdateUserPhoneNumber(ctx context.Context, in *pb.UpdateUserPhoneNumberRequest) (*pb.UpdateUserPhoneNumberResponse, error) {

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("UpdateUserPhoneNumber userId = %v", tokenUser.Id)

	user, _ := s.GetUser(tokenUser.Id)
	if user == nil {
		logger.Debugf("UpdateUserPhoneNumber user does not exist %v", tokenUser.Id)
		return nil, proto.NotFound.ErrorMsg("user does not exist")
	}

	newUser, err := s.jwtUtils.ValidateUnsigned(in.GetNewAccessToken(), false)
	if newUser == nil || err != nil || len(newUser.Claims) == 0 {
		logger.Debugf("UpdateUserPhoneNumber invalid new access token %v", in.GetNewAccessToken())
		return nil, proto.InvalidArgument.ErrorMsg("invalid new access token")
	}

	var newPhoneNumber string

	for _, claim := range newUser.Claims {
		if claim.ClaimType == security.CLAIM_TYPE_PHONE_NUMBER {
			newPhoneNumber = claim.Identifier
			break
		}
	}

	if newPhoneNumber == "" {
		logger.Debugf("UpdateUserPhoneNumber new token does not have an authorized phone number")
		return nil, proto.InvalidArgument.ErrorMsg("new token does not have an authorized phone number")
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		updatedUser := domain.ClientUser{
			ID: user.ID,
		}
		updatedUser.PhoneNumber = newPhoneNumber

		if err := tx.Model(&updatedUser).Updates(&updatedUser).Error; err != nil {
			return err
		}

		updatedClaim := domain.ClientClaim{
			UserID:     user.ID,
			ClaimType:  int(security.CLAIM_TYPE_PHONE_NUMBER),
			Identifier: newPhoneNumber,
		}

		if err := tx.Model(&updatedClaim).Updates(&updatedClaim).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		logger.Errorf("UpdateUserPhoneNumber error in updating phone number", err)
		return nil, proto.Internal.Error(err)
	}

	return &pb.UpdateUserPhoneNumberResponse{}, nil

}

func (s *Service) GetUserAddress(ctx context.Context, in *pb.GetUserAddressRequest) (*pb.GetUserAddressResponse, error) {

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("GetUserAddress userId = %v", tokenUser.Id)

	address := domain.ClientAddress{}
	if err := s.db.Model(&domain.ClientAddress{}).Where("user_id = ?", tokenUser.Id).Find(&address).Error; err != nil {
		logger.Infof("GetUserAddress error in performing query %v, error = %v", tokenUser.Id, err)
		return nil, proto.Internal.Error(err)
	}
	if address.UserID == "" {
		logger.Infof("GetUserAddress empty result %v", tokenUser.Id)
		return nil, proto.NotFound.ErrorMsg("address not found")
	}

	logger.Infof("GetUserAddress result = %v", address)

	return &pb.GetUserAddressResponse{
		Address: &pb.UserAddress{
			Street:         address.Street.String,
			Building:       address.Building.String,
			City:           address.City.String,
			County:         address.County.String,
			PostCode:       address.PostCode.String,
			AddressDetails: address.AddressDetails.String,
		},
	}, nil
}

func (s *Service) UpdateUserAddress(ctx context.Context, in *pb.UpdateUserAddressRequest) (*pb.UpdateUserAddressResponse, error) {

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("UpdateUserAddress userId = %v, address = %v", tokenUser.Id, in.GetAddress())

	address := in.GetAddress()

	err := s.db.Transaction(func(tx *gorm.DB) error {
		return tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"street", "building", "city", "county", "post_code", "address_details"}),
		}).Create(&domain.ClientAddress{
			UserID:         tokenUser.Id,
			Street:         sql.NullString{String: address.Street, Valid: address.Street != ""},
			Building:       sql.NullString{String: address.Building, Valid: address.Building != ""},
			City:           sql.NullString{String: address.City, Valid: address.City != ""},
			County:         sql.NullString{String: address.County, Valid: address.County != ""},
			PostCode:       sql.NullString{String: address.PostCode, Valid: address.PostCode != ""},
			AddressDetails: sql.NullString{String: address.AddressDetails, Valid: address.AddressDetails != ""},
		}).Error
	})

	if err != nil {
		logger.Errorf("UpdateUserAddress error in updating user address", err)
		return nil, proto.Internal.Error(err)
	}

	return &pb.UpdateUserAddressResponse{}, nil
}

func (s *Service) getAuthorizedClientClaims(userId string) ([]*pb.AuthorizedClaim, error) {
	var claims []domain.ClientClaim

	err := s.db.Model(&domain.ClientClaim{}).Where("user_id = ?", userId).Scan(&claims).Error

	if err != nil {
		return nil, err
	}

	var result = make([]*pb.AuthorizedClaim, len(claims))

	for i, claim := range claims {
		result[i] = &pb.AuthorizedClaim{
			Identifier: claim.Identifier,
			Type:       pb.ClaimType(claim.ClaimType),
		}
	}

	return result, nil

}
