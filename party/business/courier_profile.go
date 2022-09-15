package business

import (
	"context"
	"database/sql"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/party/domain"
	pb "gitlab.artin.ai/backend/courier-management/party/proto"
	"gitlab.artin.ai/backend/courier-management/uaa/proto"
	"gitlab.artin.ai/backend/courier-management/uaa/security"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Service) CreateCourierAccount(ctx context.Context, in *pb.CreateCourierAccountRequest) (*pb.CreateCourierAccountResponse, error) {
	logger.Infof("CreateCourierAccount firstName = %v, lastName = %v, birthDate = %v, email = %v, citizen = %v",
		in.GetFirstName(), in.GetLastName(), in.GetBirthDate(),
		in.GetEmail(), in.GetCitizen())

	tokenUser := ctx.Value("user").(security.User)

	existingUser := domain.CourierUser{}

	if err := s.db.Model(&domain.CourierUser{}).Where("id = ?", tokenUser.Id).Find(&existingUser).Error; err != nil{
		logger.Debugf("CreateCourierAccount cannot query %v, err = %v", tokenUser.Id, err)
		return nil, proto.Internal.Error(err)
	}

	if existingUser.ID != "" {
		logger.Debugf("CreateCourierAccount user already exists %v", tokenUser.Id)
		return nil, proto.AlreadyExists.ErrorMsg("user already exists")
	}

	var authorizedPhoneNumber = false
	for _,claim := range tokenUser.Claims {
		if claim.ClaimType == security.CLAIM_TYPE_PHONE_NUMBER{
			authorizedPhoneNumber = true
		}
	}

	if !authorizedPhoneNumber {
		logger.Debugf("CreateCourierAccount token doesn't have any authorized phone number")
		return nil, proto.Unauthenticated.ErrorMsg("token doesn't have any authorized phone number")
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// insert user
		if err := tx.Create(&domain.CourierUser{
			ID:          tokenUser.Id,
			FirstName:   sql.NullString{String: in.FirstName, Valid: in.FirstName != ""},
			LastName:    in.LastName,
			Email:       sql.NullString{String: in.Email, Valid: in.Email != ""},
			PhoneNumber: tokenUser.PhoneNumber,
			Status:      int32(pb.UserStatus_UNKNOWN_USER_STATUS),
			BirthDate:   sql.NullString{String: in.BirthDate, Valid: in.BirthDate != ""},
			Citizen:     sql.NullInt32{Int32: in.Citizen.ToInt(), Valid: in.Citizen.Valid()},
		}).Error; err != nil {
			logger.Errorf("CreateCourierAccount cannot perform update", err)
			return proto.Internal.Error(err)
		}

		// insert claims
		var claims = make([]*domain.CourierClaim, len(tokenUser.Claims))
		for i, claim := range tokenUser.Claims {
			claims[i] = &domain.CourierClaim{
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
		logger.Errorf("CreateCourierAccount error in creating profile", err)
		return nil, proto.Internal.Error(err)
	}

	logger.Debugf("CreateCourierAccount courier created successfully %v", tokenUser.Id)

	return &pb.CreateCourierAccountResponse{}, nil

}

func (s *Service) UpdateCourierAccount(ctx context.Context, in *pb.UpdateCourierAccountRequest) (*pb.UpdateCourierAccountResponse, error) {
	logger.Infof("UpdateCourierAccount firstName = %v, lastName = %v, birthDate = %v, email = %v, citizen = %v, transport type = %v, transport size = %v",
		in.GetFirstName(), in.GetLastName(), in.GetBirthDate(),
		in.GetEmail(), in.GetCitizen(), in.GetTransportationType(), in.GetTransportSize())
	tokenUser := ctx.Value("user").(security.User)

	user, _ := s.GetCourier(tokenUser.Id)
	if user == nil {
		logger.Debugf("UpdateCourierAccount user does not exist %v", tokenUser.Id)
		return nil, proto.NotFound.ErrorMsg("user does not exist")
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// update user
		updatedUser := domain.CourierUser{
			ID: user.ID,
		}
		if in.FirstName != "" {
			updatedUser.FirstName = sql.NullString{String: in.FirstName, Valid: true}
		}

		if in.Citizen.Valid() {
			updatedUser.Citizen = sql.NullInt32{Int32: in.Citizen.ToInt(), Valid: in.Citizen.Valid()}
		}

		if in.LastName != "" {
			updatedUser.LastName = in.LastName
		}

		if in.Email != "" {
			updatedUser.Email = sql.NullString{String: in.Email, Valid: true}
		}

		if in.BirthDate != "" {
			updatedUser.BirthDate = sql.NullString{String: in.BirthDate, Valid: true}
		}

		if in.TransportationType != 0 {
			updatedUser.TransportType = sql.NullInt32{Int32: int32(in.TransportationType), Valid: true}
		}

		if in.TransportSize != 0 {
			updatedUser.TransportSize = sql.NullInt32{Int32: int32(in.TransportSize), Valid: true}
		}

		if err := tx.Model(&updatedUser).Updates(&updatedUser).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		logger.Errorf("UpdateCourierAccount error in updating profile", err)
		return nil, proto.Internal.Error(err)
	}

	logger.Debugf("UpdateCourierAccount update courier successfully %v", tokenUser.Id)

	return &pb.UpdateCourierAccountResponse{}, nil

}

func (s *Service) profileInfoStatus(userId string, infoType pb.AdditionalInfoType) (pb.ProfileAdditionalInfoStatus, error) {
	var docType = infoType.DocumentInfoType()
	var validDocTypes = infoType.ValidDocumentTypes()

	// first check driver background, because it might be checked as upload documents later
	if infoType == pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVER_BACKGROUND{
		driverBackground, err := s.GetDriverBackground(userId)
		if err != nil {
			// if not found
			if e, ok := status.FromError(err); ok {
				if uint32(e.Code()) == uint32(proto.NotFound){
					return pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_EMPTY, nil
				}
			}
			return pb.ProfileAdditionalInfoStatus_UNKNOWN_PROFILE_ADDITIONAL_INFO_STATUS,err
		}
		if driverBackground != nil{
			if driverBackground.UploadDbsLater.ToBool() {
				return pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_COMPLETED, nil
			}
		} else {
			return pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_EMPTY, nil
		}
	}

	var err error
	switch infoType {
	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ID_CARD:
		_, err = s.GetIdCard(userId)
		break
	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVING_LICENSE:
		_, err = s.GetDriversLicense(userId)
		break
	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_RESIDENCE_CARD:
		_, err = s.GetResidenceCard(userId)
		break
	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_BANK_ACCOUNT:
		_, err = s.GetBankAccount(userId)
		break
	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ADDRESS:
		_, err = s.GetAddress(userId)
		break
	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_MOT:
		_, err = s.GetMot(userId)
		break
	}

	var emptyEntity = false

	if err != nil {
		if e, ok := status.FromError(err); ok {
			if uint32(e.Code()) == uint32(proto.NotFound){
				emptyEntity = true
			}
		} else {
			return pb.ProfileAdditionalInfoStatus_UNKNOWN_PROFILE_ADDITIONAL_INFO_STATUS, err
		}
	}

	if docType != pb.DocumentInfoType_UNKNOWN_DOCUMENT_INFO_TYPE && validDocTypes != nil {

		res, _ := s.getDocumentsOfUserById(userId, docType, pb.DocumentDataType_UNKNOWN_DOCUMENT_DATA_TYPE)

		if res == nil{
			if emptyEntity{
				return pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_EMPTY, nil
			}
			return pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_INCOMPLETED, nil
		}

		var documentTypeList = make([]pb.DocumentType, len(res.Documents))

		for i,d := range res.Documents{
			documentTypeList[i] = d.DocType
		}

		if infoType == pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ID_CARD{
			if pb.DocumentType_DOCUMENT_TYPE_PASSPORT.InArray(documentTypeList) ||
				(pb.DocumentType_DOCUMENT_TYPE_NATIONAL_ID_FRONT.InArray(documentTypeList) &&
					pb.DocumentType_DOCUMENT_TYPE_NATIONAL_ID_BACK.InArray(documentTypeList)){
			} else {
				if emptyEntity{
					return pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_EMPTY, nil
				}
				return pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_INCOMPLETED, nil
			}
		} else {
			for _, vtype := range validDocTypes {
				if !vtype.InArray(documentTypeList) {
					if emptyEntity{
						return pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_EMPTY, nil
					}
					return pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_INCOMPLETED, nil
				}
			}
		}
	}

	if emptyEntity{
		return pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_INCOMPLETED, nil
	}
	return pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_COMPLETED, nil

}

func (s *Service) UpdateProfileAdditionalInfo(ctx context.Context, in *pb.UpdateProfileAdditionalInfoRequest) (*pb.UpdateProfileAdditionalInfoResponse, error) {
	tokenUser := ctx.Value("user").(security.User)
	logger.Infof("UpdateProfileAdditionalInfo userId = %v", tokenUser.Id)
	return s.updateProfileAdditionalInfoByUserId(ctx, tokenUser.Id, in)
}

func (s *Service) updateProfileAdditionalInfoByUserId(ctx context.Context, userId string, in *pb.UpdateProfileAdditionalInfoRequest) (*pb.UpdateProfileAdditionalInfoResponse, error) {
	var user *domain.CourierUser
	err := s.db.Model(&domain.CourierUser{}).Where("id = ?", userId).Find(&user).Error

	if err != nil {
		logger.Infof("UpdateProfileAdditionalInfo cannot perform query userId = %v, error = %v", userId, err)
		return nil, proto.Internal.Error(err)
	} else if user.ID == "" {
		logger.Infof("UpdateProfileAdditionalInfo user not found = %v", userId)
		return nil, proto.NotFound.ErrorMsg("user does not exist")
	}

	var infoType = pb.AdditionalInfoType_UNKNOWN_ADDITIONAL_INFO_TYPE

	if in.GetIdCard() != nil {
		idCard := in.GetIdCard()
		logger.Debugf("UpdateProfileAdditionalInfo update id card %v", idCard)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"first_name", "last_name", "number", "expiration_date", "issue_place", "type"}),
			}).Create(&domain.IDCard{
				UserID:         userId,
				FirstName:      idCard.FirstName,
				LastName:       idCard.LastName,
				Number:         idCard.Number,
				ExpirationDate: idCard.ExpirationDate,
				IssuePlace:     idCard.IssuePlace,
				Type:           int32(idCard.Type),
			}).Error

			if err != nil {
				return err
			}

			infoType = pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ID_CARD

			return nil

		})

		if err != nil {
			logger.Errorf("UpdateProfileAdditionalInfo error in creating idCard", err)
		}

	} else if in.GetDrivingLicense() != nil {
		driversLicense := in.GetDrivingLicense()
		logger.Debugf("UpdateProfileAdditionalInfo update drivers license %v", driversLicense)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"driving_license_number", "expiration_date"}),
			}).Create(&domain.DrivingLicense{
				UserID:               userId,
				DrivingLicenseNumber: driversLicense.DrivingLicenseNumber,
				ExpirationDate:       driversLicense.ExpirationDate,
			}).Error

			if err != nil {
				return err
			}

			infoType = pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVING_LICENSE

			return nil
		})

		if err != nil {
			logger.Errorf("UpdateProfileAdditionalInfo error in creating driversLicense", err)
		}

	} else if in.GetDriverBackground() != nil {
		driverBackground := in.GetDriverBackground()
		logger.Debugf("UpdateProfileAdditionalInfo update driver background %v", driverBackground)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"national_insurance_number", "upload_dbs_later"}),
			}).Create(&domain.DriverBackground{
				UserID:                  userId,
				NationalInsuranceNumber: driverBackground.NationalInsuranceNumber,
				UploadDbsLater:          driverBackground.UploadDbsLater.ToInt(),
			}).Error

			if err != nil {
				return err
			}

			infoType = pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVER_BACKGROUND

			return nil
		})

		if err != nil {
			logger.Errorf("UpdateProfileAdditionalInfo error in creating driverBackground", err)
		}

	} else if in.GetResidenceCard() != nil {
		residenceCard := in.GetResidenceCard()
		logger.Debugf("UpdateProfileAdditionalInfo update residence card %v", residenceCard)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"number", "expiration_date", "issue_date"}),
			}).Create(&domain.ResidenceCard{
				UserID:         userId,
				Number:         residenceCard.Number,
				ExpirationDate: residenceCard.ExpirationDate,
				IssueDate:      residenceCard.IssueDate,
			}).Error

			if err != nil {
				return err
			}

			infoType = pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_RESIDENCE_CARD

			return nil
		})

		if err != nil {
			logger.Errorf("UpdateProfileAdditionalInfo error in creating residenceCard", err)
		}

	} else if in.GetBankAccount() != nil {
		bankAccount := in.GetBankAccount()
		logger.Debugf("UpdateProfileAdditionalInfo update bank account %v", bankAccount)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"bank_name", "account_number", "account_holder_name", "sort_code"}),
			}).Create(&domain.BankAccount{
				UserID:            userId,
				BankName:          bankAccount.BankName,
				AccountNumber:     bankAccount.AccountNumber,
				AccountHolderName: bankAccount.AccountHolderName,
				SortCode:          bankAccount.SortCode,
			}).Error

			if err != nil {
				return err
			}

			infoType = pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_BANK_ACCOUNT

			return nil
		})

		if err != nil {
			logger.Errorf("UpdateProfileAdditionalInfo error in creating bankAccount", err)
		}

	} else if in.GetAddress() != nil {
		address := in.GetAddress()
		logger.Debugf("UpdateProfileAdditionalInfo update address %v", address)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"street", "building", "city", "county", "post_code", "address_details"}),
			}).Create(&domain.CourierAddress{
				UserID:         userId,
				Street:         sql.NullString{String: address.Street, Valid: address.Street != ""},
				Building:       sql.NullString{String: address.Building, Valid: address.Building != ""},
				City:           sql.NullString{String: address.City, Valid: address.City != ""},
				County:         sql.NullString{String: address.County, Valid: address.County != ""},
				PostCode:       sql.NullString{String: address.PostCode, Valid: address.PostCode != ""},
				AddressDetails: sql.NullString{String: address.AddressDetails, Valid: address.AddressDetails != ""},
			}).Error

			if err != nil {
				return err
			}

			infoType = pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ADDRESS

			return nil
		})

	} else if in.GetMot() != nil {
		registrationNumber := in.GetMot().GetRegistrationNumber()
		logger.Debugf("UpdateProfileAdditionalInfo update mot registration number = %v", registrationNumber)
		mot, err := s.SearchMot(ctx, &pb.SearchMotRequest{
			AccessToken:        in.AccessToken,
			RegistrationNumber: registrationNumber,
		})

		if err != nil {
			return nil, err
		}

		logger.Debugf("UpdateProfileAdditionalInfo retrieved mot is %v", mot)

		err = s.db.Transaction(func(tx *gorm.DB) error {
			err = tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}, {Name: "registration_number"}},
				DoUpdates: clause.AssignmentColumns([]string{"co2_emissions", "engine_capacity", "euro_status", "marked_for_export", "fuel_type", "mot_status", "revenue_weight", "colour", "make", "type_approval", "year_of_manufacture", "tax_due_date", "tax_status", "date_of_last_v5c_issued", "real_driving_emissions", "wheelplan", "month_of_first_registration"}),
			}).Create(&domain.CourierMot{
				UserID:                   userId,
				RegistrationNumber:       mot.Mot.RegistrationNumber,
				Co2Emissions:             mot.Mot.Co2Emissions,
				EngineCapacity:           mot.Mot.EngineCapacity,
				EuroStatus:               mot.Mot.EuroStatus,
				MarkedForExport:          mot.Mot.MarkedForExport.ToBool(),
				FuelType:                 mot.Mot.FuelType,
				MotStatus:                domain.MotStatus(mot.Mot.MotStatus),
				RevenueWeight:            mot.Mot.RevenueWeight,
				Colour:                   mot.Mot.Colour,
				Make:                     mot.Mot.Make,
				TypeApproval:             mot.Mot.TypeApproval,
				YearOfManufacture:        mot.Mot.YearOfManufacture,
				TaxDueDate:               mot.Mot.TaxDueDate,
				TaxStatus:                domain.TaxStatus(mot.Mot.TaxStatus),
				DateOfLastV5CIssued:      mot.Mot.DateOfLastV5CIssued,
				RealDrivingEmissions:     mot.Mot.RealDrivingEmissions,
				Wheelplan:                mot.Mot.Wheelplan,
				MonthOfFirstRegistration: mot.Mot.MonthOfFirstRegistration,
			}).Error

			if err != nil {
				return err
			}

			infoType = pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_MOT

			return nil
		})

	}

	if err != nil {
		logger.Errorf("UpdateProfileAdditionalInfo error in updating profile additional info", err)
		return nil, proto.Internal.Error(err)
	}

	if infoType != pb.AdditionalInfoType_UNKNOWN_ADDITIONAL_INFO_TYPE {
		newStatus, err := s.profileInfoStatus(userId, infoType)
		if err != nil {
			return nil, proto.Internal.Error(err)
		}
		if newStatus != pb.ProfileAdditionalInfoStatus_UNKNOWN_PROFILE_ADDITIONAL_INFO_STATUS {
			err = s.db.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}, {Name: "status_type"}},
				DoUpdates: clause.AssignmentColumns([]string{"status", "message"}),
			}).Create(&domain.CourierStatus{
				UserID:     userId,
				StatusType: int32(infoType),
				Status:     int32(newStatus),
				Message:    "",
			}).Error

			if err != nil {
				return nil, proto.Internal.Error(err)
			}
		}
	}

	logger.Debugf("UpdateProfileAdditionalInfo update additional info successfully %v", userId)

	return &pb.UpdateProfileAdditionalInfoResponse{}, nil
}

func (s *Service) GetProfileAdditionalInfo(ctx context.Context, in *pb.GetProfileAdditionalInfoRequest) (*pb.GetProfileAdditionalInfoResponse, error) {
	tokenUser := ctx.Value("user").(security.User)
	return s.getProfileAdditionalInfoByUserid(tokenUser.Id, in.Type)
}

func (s *Service) getProfileAdditionalInfoByUserid(userId string, reqType pb.AdditionalInfoType) (*pb.GetProfileAdditionalInfoResponse, error) {
	logger.Infof("getProfileAdditionalInfoByUserid userId = %v, type = %v", userId, reqType)

	switch reqType {
	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ID_CARD:
		result, err := s.GetIdCard(userId)
		if err != nil || result == nil {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find id_card", err)
			return nil, err
		}
		if err = s.findDocuments(userId, result, false); err != nil {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find documents", err)
			return nil, err
		}
		logger.Debugf("getProfileAdditionalInfoByUserid id card is %v", result)
		return &pb.GetProfileAdditionalInfoResponse{
			Info: &pb.GetProfileAdditionalInfoResponse_IdCard{
				IdCard: result,
			},
		}, nil
	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVING_LICENSE:
		result, err := s.GetDriversLicense(userId)
		if err != nil || result == nil {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find drivers_license", err)
			return nil, err
		}
		if err = s.findDocuments(userId, result, false); err != nil {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find documents", err)
			return nil, err
		}
		logger.Debugf("getProfileAdditionalInfoByUserid driving license is %v", result)
		return &pb.GetProfileAdditionalInfoResponse{
			Info: &pb.GetProfileAdditionalInfoResponse_DrivingLicense{
				DrivingLicense: result,
			},
		}, nil
	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVER_BACKGROUND:
		result, err := s.GetDriverBackground(userId)
		if err != nil || result == nil {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find driver_background", err)
			return nil, err
		}
		if err = s.findDocuments(userId, result, false); err != nil {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find documents", err)
			return nil, err
		}
		logger.Debugf("getProfileAdditionalInfoByUserid driver background is %v", result)
		return &pb.GetProfileAdditionalInfoResponse{
			Info: &pb.GetProfileAdditionalInfoResponse_DriverBackground{
				DriverBackground: result,
			},
		}, nil
	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_RESIDENCE_CARD:
		result, err := s.GetResidenceCard(userId)
		if err != nil || result == nil {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find residence_card", err)
			return nil, err
		}
		if err = s.findDocuments(userId, result, false); err != nil {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find documents", err)
			return nil, err
		}
		logger.Debugf("getProfileAdditionalInfoByUserid residence card is %v", result)
		return &pb.GetProfileAdditionalInfoResponse{
			Info: &pb.GetProfileAdditionalInfoResponse_ResidenceCard{
				ResidenceCard: result,
			},
		}, nil
	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_BANK_ACCOUNT:
		result, err := s.GetBankAccount(userId)
		if err != nil || result == nil {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find bank_account", err)
			return nil, err
		}
		if err = s.findDocuments(userId, result, false); err != nil {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find documents", err)
			return nil, err
		}
		logger.Debugf("getProfileAdditionalInfoByUserid bank account is %v", result)
		return &pb.GetProfileAdditionalInfoResponse{
			Info: &pb.GetProfileAdditionalInfoResponse_BankAccount{
				BankAccount: result,
			},
		}, nil
	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ADDRESS:
		result, err := s.GetAddress(userId)
		if err != nil || result == nil {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find address", err)
			return nil, err
		}
		if err = s.findDocuments(userId, result, false); err != nil {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find documents", err)
			return nil,err
		}
		logger.Debugf("getProfileAdditionalInfoByUserid address is %v", result)
		return &pb.GetProfileAdditionalInfoResponse{
			Info: &pb.GetProfileAdditionalInfoResponse_Address{
				Address: result,
			},
		}, nil
	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_INSURANCE_CERTIFICATE:
		result := &pb.InsuranceCertificate{}
		if err := s.findDocuments(userId, result, false); err != nil {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find documents", err)
			return nil, err
		}
		if result.GetDocumentIds() == nil || len(result.GetDocumentIds()) == 0{
			logger.Debugf("getProfileAdditionalInfoByUserid cannot find insurance certificate")
			return nil, proto.NotFound.ErrorNoMsg()
		}
		logger.Debugf("getProfileAdditionalInfoByUserid insurance certificate is %v", result)
		return &pb.GetProfileAdditionalInfoResponse{
			Info: &pb.GetProfileAdditionalInfoResponse_InsuranceCertificate{
				InsuranceCertificate: result,
			},
		}, nil
	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_MOT:
		result, err := s.GetMot(userId)
		if err != nil || result == nil || len(result) < 1 {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find mot", err)
			return nil, err
		}
		logger.Debugf("getProfileAdditionalInfoByUserid mot is %v", result)
		return &pb.GetProfileAdditionalInfoResponse{
			Info: &pb.GetProfileAdditionalInfoResponse_Mot{
				Mot: result[0],
			},
		}, nil

	case pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_PROFILE_PICTURE:
		result := &pb.ProfilePicture{}
		if err := s.findDocuments(userId, result, true); err != nil {
			logger.Errorf("getProfileAdditionalInfoByUserid cannot find documents", err)
			return nil,err
		}
		if result.GetDocumentIds() == nil || len(result.GetDocumentIds()) == 0{
			logger.Debugf("getProfileAdditionalInfoByUserid cannot find profile picture")
			return nil, proto.NotFound.ErrorNoMsg()
		}
		logger.Debugf("getProfileAdditionalInfoByUserid profile picture is %v", result)
		return &pb.GetProfileAdditionalInfoResponse{
			Info: &pb.GetProfileAdditionalInfoResponse_ProfilePicture{
				ProfilePicture: result,
			},
		}, nil

	case pb.AdditionalInfoType_UNKNOWN_ADDITIONAL_INFO_TYPE:
		result := &pb.ProfileAdditionalInfo{}
		if item,_ := s.getProfileAdditionalInfoByUserid(userId, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ID_CARD); item != nil {
			result.IdCard = item.GetIdCard()
		}
		if item,_ := s.getProfileAdditionalInfoByUserid(userId, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVING_LICENSE); item != nil {
			result.DrivingLicense = item.GetDrivingLicense()
		}
		if item,_ := s.getProfileAdditionalInfoByUserid(userId, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVER_BACKGROUND); item != nil {
			result.DriverBackground = item.GetDriverBackground()
		}
		if item,_ := s.getProfileAdditionalInfoByUserid(userId, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_RESIDENCE_CARD); item != nil {
			result.ResidenceCard = item.GetResidenceCard()
		}
		if item,_ := s.getProfileAdditionalInfoByUserid(userId, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_BANK_ACCOUNT); item != nil {
			result.BankAccount = item.GetBankAccount()
		}
		if item,_ := s.getProfileAdditionalInfoByUserid(userId, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ADDRESS); item != nil {
			result.Address = item.GetAddress()
		}
		if item,_ := s.getProfileAdditionalInfoByUserid(userId, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_INSURANCE_CERTIFICATE); item != nil {
			result.InsuranceCertificate = item.GetInsuranceCertificate()
		}
		if item,_ := s.getProfileAdditionalInfoByUserid(userId, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_MOT); item != nil {
			result.Mot = item.GetMot()
		}
		if item,_ := s.getProfileAdditionalInfoByUserid(userId, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_PROFILE_PICTURE); item != nil {
			result.ProfilePicture = item.GetProfilePicture()
		}
		return &pb.GetProfileAdditionalInfoResponse{
			Info: &pb.GetProfileAdditionalInfoResponse_ProfileAdditionalInfo{
				ProfileAdditionalInfo: result,
			},
		}, nil
	}



	return nil, nil
}

func (s *Service) GetIdCard(userId string) (*pb.IDCard, error) {
	logger.Debugf("GetIdCard userId = %v ", userId)

	idCard := domain.IDCard{}

	if err := s.db.Model(&domain.IDCard{}).Where("user_id = ?", userId).Find(&idCard).Error; err != nil{
		return nil, proto.Internal.Error(err)
	}

	if idCard.UserID == "" {
		return nil, proto.NotFound.ErrorMsg("idCard not found")
	}
	logger.Debugf("GetIdCard retrieved id card = %v ", idCard)
	return &pb.IDCard{
		FirstName:      idCard.FirstName,
		LastName:       idCard.LastName,
		Number:         idCard.Number,
		ExpirationDate: idCard.ExpirationDate,
		IssuePlace:     idCard.IssuePlace,
		Type:           pb.IDCardType(idCard.Type),
	}, nil
}

func (s *Service) GetDriversLicense(userId string) (*pb.DrivingLicense, error) {
	logger.Debugf("GetDriversLicense userId = %v ", userId)

	driverLicense := domain.DrivingLicense{}

	if err := s.db.Model(&domain.DrivingLicense{}).Where("user_id = ?", userId).Find(&driverLicense).Error; err != nil{
		return nil, proto.Internal.Error(err)
	}

	if driverLicense.UserID == "" {
		return nil, proto.NotFound.ErrorMsg("driverLicense not found")
	}
	logger.Infof("GetDriversLicense retrieved driverLicense = %v", driverLicense)
	return &pb.DrivingLicense{
		DrivingLicenseNumber: driverLicense.DrivingLicenseNumber,
		ExpirationDate:       driverLicense.ExpirationDate,
	}, nil
}

func (s *Service) GetDriverBackground(userId string) (*pb.DriverBackground, error) {
	logger.Debugf("GetDriverBackground userId = %v ", userId)

	driverBackground := domain.DriverBackground{}

	if err := s.db.Model(&domain.DriverBackground{}).Where("user_id = ?", userId).Find(&driverBackground).Error; err != nil{
		return nil, proto.Internal.Error(err)
	}

	if driverBackground.UserID == "" {
		return nil, proto.NotFound.ErrorMsg("GetDriverBackground driverBackground not found")
	}
	logger.Infof("GetDriverBackground retrieved driver background = %v", driverBackground)
	var uploadDbsLater = pb.Boolean_UNKNOWN_BOOLEAN
	uploadDbsLater = uploadDbsLater.FromInt(driverBackground.UploadDbsLater)
	return &pb.DriverBackground{
		NationalInsuranceNumber: driverBackground.NationalInsuranceNumber,
		UploadDbsLater:          uploadDbsLater,
	}, nil
}

func (s *Service) GetResidenceCard(userId string) (*pb.ResidenceCard, error) {
	logger.Debugf("GetResidenceCard userId = %v ", userId)

	residenceCard := domain.ResidenceCard{}

	if err := s.db.Model(&domain.ResidenceCard{}).Where("user_id = ?", userId).Find(&residenceCard).Error; err != nil{
		return nil, proto.Internal.Error(err)
	}

	if residenceCard.UserID == "" {
		return nil, proto.NotFound.ErrorMsg("residenceCard not found")
	}
	logger.Infof("GetResidenceCard retrieved residence card = %v", residenceCard)
	return &pb.ResidenceCard{
		Number:         residenceCard.Number,
		ExpirationDate: residenceCard.ExpirationDate,
		IssueDate:      residenceCard.IssueDate,
	}, nil
}

func (s *Service) GetBankAccount(userId string) (*pb.BankAccount, error) {
	logger.Debugf("GetBankAccount userId = %v ", userId)

	bankAccount := domain.BankAccount{}

	if err := s.db.Model(&domain.BankAccount{}).Where("user_id = ?", userId).Find(&bankAccount).Error; err != nil{
		return nil, proto.Internal.Error(err)
	}

	if bankAccount.UserID == "" {
		return nil, proto.NotFound.ErrorMsg("bankAccount not found")
	}
	logger.Infof("GetBankAccount retrieved bank account = %v", bankAccount)
	return &pb.BankAccount{
		BankName:          bankAccount.BankName,
		AccountNumber:     bankAccount.AccountNumber,
		AccountHolderName: bankAccount.AccountHolderName,
		SortCode:          bankAccount.SortCode,
	}, nil
}

func (s *Service) GetAddress(userId string) (*pb.Address, error) {
	logger.Debugf("GetAddress userId = %v ", userId)

	address := domain.CourierAddress{}

	if err := s.db.Model(&domain.CourierAddress{}).Where("user_id = ?", userId).Find(&address).Error; err != nil{
		return nil, proto.Internal.Error(err)
	}

	if address.UserID == "" {
		return nil, proto.NotFound.ErrorMsg("address not found")
	}
	logger.Infof("GetAddress retrieved address = %v", address)
	return &pb.Address{
		Street:         address.Street.String,
		Building:       address.Building.String,
		City:           address.City.String,
		County:         address.County.String,
		PostCode:       address.PostCode.String,
		AddressDetails: address.AddressDetails.String,
	}, nil
}

func (s *Service) findDocuments(userId string, entity hasDocument, loadData bool) error {
	logger.Debugf("findDocuments userId = %v ", userId)

	var documents []*domain.Document
	err := s.db.Model(&domain.Document{}).Where("user_id=? AND document_info_type=?", userId, entity.GetDocumentInfoType()).Scan(&documents).Error

	if err != nil {
		return proto.Internal.Error(err)
	}

	if len(documents) > 0 {
		var documents2 = make([]*pb.DocumentInfo, len(documents))
		for i,d := range documents{
			var data []byte
			if loadData {
				documentData := domain.DocumentData{}
				s.db.Model(&domain.DocumentData{}).Where("object_id=?", d.ObjectId).Find(&documentData)
				data = documentData.Data
			}
			documents2[i] = &pb.DocumentInfo{
				ObjectId: d.ObjectId,
				InfoType: pb.DocumentInfoType(d.DocumentInfoType),
				DocType: pb.DocumentType(d.DocumentType),
				FileType: d.FileType.String,
				Data: data,
				CreationDate: d.CreationTime.Format("2006-01-02 15:04:05"),
			}
		}
		entity.SetDocumentIds(documents2)
	}

	return nil
}

func (s *Service) GetCourierAccount(ctx context.Context, in *pb.GetCourierAccountRequest) (*pb.GetCourierAccountResponse, error) {

	tokenUser := ctx.Value("user").(security.User)
	logger.Infof("GetCourierAccount %v", tokenUser.Id)

	return s.getCourierAccountById(ctx, tokenUser.Id)
}

func (s *Service) getCourierAccountById(ctx context.Context, userId string) (*pb.GetCourierAccountResponse, error) {

	logger.Infof("GetCourierAccountById %v", userId)

	var user domain.CourierUser

	err := s.db.Model(&domain.CourierUser{}).Where("id = ?", userId).Find(&user).Error

	if err != nil {
		logger.Infof("GetCourierAccount error in db %v, error = %v", userId, err)
		return nil, proto.Internal.Error(err)
	}

	if user.ID == "" {
		logger.Infof("GetCourierAccount courier not found %v", userId)
		return nil, proto.NotFound.ErrorMsg("courier not found")
	}

	logger.Debugf("GetCourierAccount retrieved courier account %v", user)

	var citizen = pb.Boolean_UNKNOWN_BOOLEAN
	citizen = citizen.FromInt(user.Citizen.Int32)

	authorizedClaims,err := s.getAuthorizedCourierClaims(userId)

	if err != nil {
		logger.Infof("GetCourierAccount error in get authorized claims = %v", err)
		return nil, proto.Internal.Error(err)
	}

	return &pb.GetCourierAccountResponse{
		Profile: &pb.CourierProfile{
			UserId:        user.ID,
			PhoneNumber:   user.PhoneNumber,
			FirstName:     user.FirstName.String,
			LastName:      user.LastName,
			Email:         user.Email.String,
			BirthDate:     user.BirthDate.String,
			TransportType: pb.TransportationType(user.TransportType.Int32),
			TransportSize: pb.TransportationSize(user.TransportSize.Int32),
			Citizen:       citizen,
			AuthorizedClaims: authorizedClaims,
		},
	}, nil

}

func (s *Service) FindCourierAccounts(ctx context.Context, in *pb.FindCourierAccountsRequest) (*pb.FindCourierAccountsResponse, error) {
	logger.Debugf("FindCourierAccounts email = %v, userId = %v, name = %v, phone number = %v", in.GetEmail(), in.GetUserId(),
		in.GetName(), in.GetPhoneNumber())

	var users []domain.CourierUser

	var err error

	if in.GetUserId() != "" {
		err = s.db.Model(&domain.CourierUser{}).Where("id = ?", in.GetUserId()).Scan(&users).Error
	} else if in.GetName() != "" {
		err = s.db.Model(&domain.CourierUser{}).Where("first_name LIKE ? OR last_name LIKE ?", "%"+in.GetName()+"%", "%"+in.GetName()+"%").Scan(&users).Error
	} else {
		var identifier string
		if in.GetPhoneNumber() != "" {
			identifier = in.GetPhoneNumber()
		} else if in.GetEmail() != "" {
			identifier = in.GetEmail()
		}
		err = s.db.Model(&domain.CourierUser{}).Where("id IN (SELECT user_id FROM courier_claim WHERE identifier=?)", identifier).Scan(&users).Error
	}

	if err != nil {
		return nil, proto.Internal.Error(err)
	}

	if len(users) == 0{
		return nil, proto.NotFound.ErrorMsg("no user found")
	}

	logger.Debugf("FindCourierAccounts retrieved users = %v", users)

	var profiles = make([]*pb.CourierProfile, len(users))

	for i, user := range users {
		var citizen = pb.Boolean_UNKNOWN_BOOLEAN
		citizen = citizen.FromInt(user.Citizen.Int32)

		authorizedClaims,err := s.getAuthorizedCourierClaims(user.ID)

		if err != nil {
			logger.Infof("FindCourierAccounts error in get authorized claims = %v", err)
			return nil, proto.Internal.Error(err)
		}

		profiles[i] = &pb.CourierProfile{
			UserId:        user.ID,
			PhoneNumber:   user.PhoneNumber,
			FirstName:     user.FirstName.String,
			LastName:      user.LastName,
			Email:         user.Email.String,
			BirthDate:     user.BirthDate.String,
			TransportType: pb.TransportationType(user.TransportType.Int32),
			TransportSize: pb.TransportationSize(user.TransportSize.Int32),
			Citizen:       citizen,
			AuthorizedClaims: authorizedClaims,
		}
	}

	return &pb.FindCourierAccountsResponse{
		Profiles: profiles,
	}, nil
}

func (s *Service) DeleteProfileAdditionalInfo(ctx context.Context, in *pb.DeleteProfileAdditionalInfoRequest) (*pb.DeleteProfileAdditionalInfoResponse, error) {

	tokenUser := ctx.Value("user").(security.User)

	logger.Debugf("DeleteProfileAdditionalInfo id = %v, type = %v ", tokenUser.Id, in.GetType())

	var err error

	var infoTypeList = make([]pb.AdditionalInfoType, 0)

	if in.GetType() == pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ID_CARD ||
		in.GetType() == pb.AdditionalInfoType_UNKNOWN_ADDITIONAL_INFO_TYPE {
		infoTypeList = append(infoTypeList, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ID_CARD)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Delete(&domain.IDCard{
				UserID: tokenUser.Id,
			}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}

			if err := tx.Model(&domain.Document{}).Where("user_id = ? and document_info_type = ?",
				tokenUser.Id, int(pb.DocumentInfoType_DOCUMENT_INFO_TYPE_ID_CARD)).
				Delete(&domain.Document{}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}

			return nil
		})
	}
	if in.GetType() == pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVING_LICENSE ||
		in.GetType() == pb.AdditionalInfoType_UNKNOWN_ADDITIONAL_INFO_TYPE {
		infoTypeList = append(infoTypeList, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVING_LICENSE)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Delete(&domain.DrivingLicense{
				UserID: tokenUser.Id,
			}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}

			if err := tx.Model(&domain.Document{}).Where("user_id = ? and document_info_type = ?",
				tokenUser.Id, int(pb.DocumentInfoType_DOCUMENT_INFO_TYPE_DRIVING_LICENSE)).
				Delete(&domain.Document{}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}

			return nil
		})
	}
	if in.GetType() == pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVER_BACKGROUND ||
		in.GetType() == pb.AdditionalInfoType_UNKNOWN_ADDITIONAL_INFO_TYPE {
		infoTypeList = append(infoTypeList, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVER_BACKGROUND)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Delete(&domain.DriverBackground{
				UserID: tokenUser.Id,
			}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}

			if err := tx.Model(&domain.Document{}).Where("user_id = ? and document_info_type = ?",
				tokenUser.Id, int(pb.DocumentInfoType_DOCUMENT_INFO_TYPE_DRIVER_BACKGROUND)).
				Delete(&domain.Document{}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}

			return nil
		})
	}
	if in.GetType() == pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_RESIDENCE_CARD ||
		in.GetType() == pb.AdditionalInfoType_UNKNOWN_ADDITIONAL_INFO_TYPE {
		infoTypeList = append(infoTypeList, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_RESIDENCE_CARD)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Delete(&domain.ResidenceCard{
				UserID: tokenUser.Id,
			}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}

			if err := tx.Model(&domain.Document{}).Where("user_id = ? and document_info_type = ?",
				tokenUser.Id, int(pb.DocumentInfoType_DOCUMENT_INFO_TYPE_RESIDENCE_CARD)).
				Delete(&domain.Document{}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}

			return nil
		})
	}
	if in.GetType() == pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_BANK_ACCOUNT ||
		in.GetType() == pb.AdditionalInfoType_UNKNOWN_ADDITIONAL_INFO_TYPE {
		infoTypeList = append(infoTypeList, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_BANK_ACCOUNT)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Delete(&domain.BankAccount{
				UserID: tokenUser.Id,
			}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}

			if err := tx.Model(&domain.Document{}).Where("user_id = ? and document_info_type = ?",
				tokenUser.Id, int(pb.DocumentInfoType_DOCUMENT_INFO_TYPE_BANK_ACCOUNT)).
				Delete(&domain.Document{}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}

			return nil
		})
	}
	if in.GetType() == pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ADDRESS ||
		in.GetType() == pb.AdditionalInfoType_UNKNOWN_ADDITIONAL_INFO_TYPE {
		infoTypeList = append(infoTypeList, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ADDRESS)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Delete(&domain.CourierAddress{
				UserID: tokenUser.Id,
			}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}

			if err := tx.Model(&domain.Document{}).Where("user_id = ? and document_info_type = ?",
				tokenUser.Id, int(pb.DocumentInfoType_DOCUMENT_INFO_TYPE_ADDRESS)).
				Delete(&domain.Document{}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}

			return nil
		})
	}
	if in.GetType() == pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_MOT ||
		in.GetType() == pb.AdditionalInfoType_UNKNOWN_ADDITIONAL_INFO_TYPE {
		infoTypeList = append(infoTypeList, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_MOT)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Delete(&domain.CourierMot{
				UserID: tokenUser.Id,
			}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}
			return nil
		})
	}
	if in.GetType() == pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_INSURANCE_CERTIFICATE ||
		in.GetType() == pb.AdditionalInfoType_UNKNOWN_ADDITIONAL_INFO_TYPE {
		infoTypeList = append(infoTypeList, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_INSURANCE_CERTIFICATE)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(&domain.Document{}).Where("user_id = ? and document_info_type = ?",
				tokenUser.Id, int(pb.DocumentInfoType_DOCUMENT_INFO_TYPE_INSURANCE_CERTIFICATE)).
				Delete(&domain.Document{}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}
			return nil
		})
	}
	if in.GetType() == pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_PROFILE_PICTURE ||
		in.GetType() == pb.AdditionalInfoType_UNKNOWN_ADDITIONAL_INFO_TYPE {
		infoTypeList = append(infoTypeList, pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_PROFILE_PICTURE)
		err = s.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Model(&domain.Document{}).Where("user_id = ? and document_info_type = ?",
				tokenUser.Id, int(pb.DocumentInfoType_DOCUMENT_INFO_TYPE_PROFILE_PICTURE)).
				Delete(&domain.Document{}).Error; err != nil {
				logger.Errorf("DeleteProfileAdditionalInfo cannot perform delete", err)
				return proto.Internal.Error(err)
			}
			return nil
		})
	}

	if err != nil {
		return nil, proto.Internal.Error(err)
	}

	for _, infoType := range infoTypeList{
		newStatus, err := s.profileInfoStatus(tokenUser.Id, infoType)
		if err != nil {
			return nil, proto.Internal.Error(err)
		}
		if newStatus != pb.ProfileAdditionalInfoStatus_UNKNOWN_PROFILE_ADDITIONAL_INFO_STATUS {
			err = s.db.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}, {Name: "status_type"}},
				DoUpdates: clause.AssignmentColumns([]string{"status", "message"}),
			}).Create(&domain.CourierStatus{
				UserID:     tokenUser.Id,
				StatusType: int32(infoType),
				Status:     int32(newStatus),
				Message:    "",
			}).Error

			if err != nil {
				return nil, proto.Internal.Error(err)
			}
		}
	}

	return &pb.DeleteProfileAdditionalInfoResponse{}, nil
}

func (s *Service) GetProfileAdditionalInfoStatus(ctx context.Context, in *pb.GetProfileAdditionalInfoStatusRequest) (*pb.GetProfileAdditionalInfoStatusResponse, error) {
	tokenUser := ctx.Value("user").(security.User)
	return s.getProfileAdditionalInfoStatusByUserId(tokenUser.Id)
}

func (s *Service) getProfileAdditionalInfoStatusByUserId(userId string) (*pb.GetProfileAdditionalInfoStatusResponse, error) {

	logger.Infof("getProfileAdditionalInfoStatusByUserId userId = %v", userId)

	var user domain.CourierUser

	err := s.db.Model(&domain.CourierUser{}).Where("id = ?", userId).Find(&user).Error

	if err != nil {
		logger.Infof("getProfileAdditionalInfoStatusByUserId userId = %v, error = %v", userId, err)
		return nil, proto.Internal.Error(err)
	}

	if user.ID == "" {
		logger.Infof("getProfileAdditionalInfoStatusByUserId courier not found %v", userId)
		return nil, proto.NotFound.ErrorMsg("courier not found")
	}

	var statusListStored []*domain.CourierStatus

	err = s.db.Model(&domain.CourierStatus{}).Where("user_id = ?", userId).Scan(&statusListStored).Error

	if err != nil {
		return nil, proto.Internal.Error(err)
	}

	var bicycle = user.TransportType.Int32 == int32(pb.TransportationType_TRANSPORTATION_TYPE_BICYCLE)

	var statusList = map[pb.AdditionalInfoType]*pb.GetProfileAdditionalInfoStatusResponseItem{}
	statusList[pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ID_CARD] = &pb.GetProfileAdditionalInfoStatusResponseItem{
		Status: pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_EMPTY,
	}
	statusList[pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVER_BACKGROUND] = &pb.GetProfileAdditionalInfoStatusResponseItem{
		Status: pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_EMPTY,
	}
	statusList[pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_BANK_ACCOUNT] = &pb.GetProfileAdditionalInfoStatusResponseItem{
		Status: pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_EMPTY,
	}
	statusList[pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ADDRESS] = &pb.GetProfileAdditionalInfoStatusResponseItem{
		Status: pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_EMPTY,
	}
	if !bicycle {
		statusList[pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVING_LICENSE] = &pb.GetProfileAdditionalInfoStatusResponseItem{
			Status: pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_EMPTY,
		}
	}
	var citizen = pb.Boolean_UNKNOWN_BOOLEAN
	citizen = citizen.FromInt(user.Citizen.Int32)
	if citizen != pb.Boolean_BOOLEAN_TRUE {
		statusList[pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_RESIDENCE_CARD] = &pb.GetProfileAdditionalInfoStatusResponseItem{
			Status: pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_EMPTY,
		}
	}
	if !bicycle {
		statusList[pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_MOT] = &pb.GetProfileAdditionalInfoStatusResponseItem{
			Status: pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_EMPTY,
		}
	}
	if !bicycle {
		statusList[pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_INSURANCE_CERTIFICATE] = &pb.GetProfileAdditionalInfoStatusResponseItem{
			Status: pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_EMPTY,
		}
	}
	statusList[pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_PROFILE_PICTURE] = &pb.GetProfileAdditionalInfoStatusResponseItem{
		Status: pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_EMPTY,
	}

	for _, status := range statusListStored {
		key := pb.AdditionalInfoType(status.StatusType)
		if _, ok := statusList[key]; ok {
			statusConv := pb.ProfileAdditionalInfoStatus(status.Status)
			if statusConv != pb.ProfileAdditionalInfoStatus_UNKNOWN_PROFILE_ADDITIONAL_INFO_STATUS {
				statusList[key] = &pb.GetProfileAdditionalInfoStatusResponseItem{
					Status:  statusConv,
					Message: status.Message,
				}
			}
		} else {
			logger.Debugf("getProfileAdditionalInfoStatusByUserId invalid status type %v", status.StatusType)
		}
	}

	var items []*pb.GetProfileAdditionalInfoStatusResponseItem
	for k, v := range statusList {
		items = append(items, &pb.GetProfileAdditionalInfoStatusResponseItem{
			Type:    k,
			Status:  v.Status,
			Message: v.Message,
		})
	}

	logger.Debugf("getProfileAdditionalInfoStatusByUserId retrieved user status = %v", items)

	return &pb.GetProfileAdditionalInfoStatusResponse{
		Items: items,
	}, nil
}

func (s *Service) UpdateProfileAdditionalInfoStatus(ctx context.Context, in *pb.UpdateProfileAdditionalInfoStatusRequest) (*pb.UpdateProfileAdditionalInfoStatusResponse, error) {
	tokenUser := ctx.Value("user").(security.User)
	return s.updateProfileAdditionalInfoStatusById(tokenUser.Id, in.GetType(), in.GetStatus(), in.GetMessage())
}

func (s *Service) updateProfileAdditionalInfoStatusById(userId string, statusType pb.AdditionalInfoType, status pb.UpdateProfileAdditionalInfoStatus, message string) (*pb.UpdateProfileAdditionalInfoStatusResponse, error) {
	logger.Infof("updateProfileAdditionalInfoStatusById userId = %v, statusType = %v, status = %v, message = %v ", userId, statusType, status, message)

	var newStatus pb.ProfileAdditionalInfoStatus
	if status == pb.UpdateProfileAdditionalInfoStatus_UPDATE_PROFILE_ADDITIONAL_INFO_STATUS_ACCEPTED {
		newStatus = pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_ACCEPTED
	} else if status == pb.UpdateProfileAdditionalInfoStatus_UPDATE_PROFILE_ADDITIONAL_INFO_STATUS_REJECTED {
		newStatus = pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_REJECTED
	} else {
		return nil, proto.InvalidArgument.ErrorMsg("invalid status")
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		return tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "status_type"}},
			DoUpdates: clause.AssignmentColumns([]string{"status", "message"}),
		}).Create(&domain.CourierStatus{
			UserID:     userId,
			StatusType: int32(statusType),
			Status:     int32(newStatus),
			Message:    message,
		}).Error
	})

	if err != nil {
		return nil, proto.Internal.Error(err)
	}

	return &pb.UpdateProfileAdditionalInfoStatusResponse{}, nil

}

func (s *Service) GetProfileStatus(ctx context.Context, in *pb.GetProfileStatusRequest) (*pb.GetProfileStatusResponse, error) {

	status, err := s.GetProfileAdditionalInfoStatus(ctx, &pb.GetProfileAdditionalInfoStatusRequest{
		AccessToken: in.AccessToken,
	})

	if err != nil {
		return nil, proto.Internal.Error(err)
	}

	var statusItems = make([]pb.ProfileAdditionalInfoStatus,0)

	for _,s := range status.Items{
		if s.Type != pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_PROFILE_PICTURE {
			statusItems = append(statusItems, s.Status)
		}
	}

	countAccepted := 0
	countCompleted := 0
	countRejected := 0

	for _, item := range statusItems {
		if item == pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_ACCEPTED {
			countAccepted = countAccepted + 1
		} else if item == pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_COMPLETED {
			countCompleted = countCompleted + 1
		} else if item == pb.ProfileAdditionalInfoStatus_PROFILE_ADDITIONAL_INFO_STATUS_REJECTED {
			countRejected = countRejected + 1
		}
	}

	var profileStatus pb.ProfileStatus

	if countRejected > 0 {
		profileStatus = pb.ProfileStatus_PROFILE_STATUS_REJECTED
	} else if countAccepted == len(statusItems) {
		profileStatus = pb.ProfileStatus_PROFILE_STATUS_COMPLETED
	} else if countCompleted + countAccepted == len(statusItems) {
		profileStatus = pb.ProfileStatus_PROFILE_STATUS_WAITING_FOR_VERIFY
	} else {
		profileStatus = pb.ProfileStatus_PROFILE_STATUS_IN_PROGRESS
	}

	logger.Debugf("GetProfileStatus result = %v", profileStatus)

	return &pb.GetProfileStatusResponse{
		Status: profileStatus,
	}, nil

}

func (s *Service) UpdateCourierPhoneNumber(ctx context.Context, in *pb.UpdateCourierPhoneNumberRequest) (*pb.UpdateCourierPhoneNumberResponse, error) {

	tokenUser := ctx.Value("user").(security.User)

	logger.Infof("UpdateCourierPhoneNumber %v", tokenUser.Id)

	user, _ := s.GetCourier(tokenUser.Id)
	if user == nil {
		logger.Debugf("UpdateCourierPhoneNumber user does not exist %v", tokenUser.Id)
		return nil, proto.NotFound.ErrorMsg("user does not exist")
	}

	newUser, err := s.jwtUtils.ValidateUnsigned(in.GetNewAccessToken(), false)
	if newUser == nil || err != nil || len(newUser.Claims) == 0 {
		logger.Debugf("UpdateCourierPhoneNumber invalid new access token %v", in.GetNewAccessToken())
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
		logger.Debugf("UpdateCourierPhoneNumber new token does not have an authorized phone number")
		return nil, proto.InvalidArgument.ErrorMsg("new token does not have an authorized phone number")
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		updatedUser := domain.CourierUser{
			ID: user.ID,
		}
		updatedUser.PhoneNumber = newPhoneNumber

		if err := tx.Model(&updatedUser).Updates(&updatedUser).Error; err != nil {
			return err
		}

		updatedClaim := domain.CourierClaim{
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
		logger.Errorf("UpdateCourierPhoneNumber error in updating phone number", err)
		return nil, proto.Internal.Error(err)
	}

	return &pb.UpdateCourierPhoneNumberResponse{}, nil

}

func (s *Service) getAuthorizedCourierClaims(userId string) ([]*pb.AuthorizedClaim, error) {
	var claims []domain.CourierClaim

	err := s.db.Model(&domain.CourierClaim{}).Where("user_id = ?", userId).Scan(&claims).Error

	if err != nil{
		return nil,err
	}

	var result = make([]*pb.AuthorizedClaim, len(claims))

	for i,claim := range claims {
		result[i] = &pb.AuthorizedClaim{
			Identifier: claim.Identifier,
			Type: pb.ClaimType(claim.ClaimType),
		}
	}

	return result, nil

}