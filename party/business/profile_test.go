package business

import (
	"context"
	"fmt"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/party/db"
	"gitlab.artin.ai/backend/courier-management/party/domain"
	_ "gitlab.artin.ai/backend/courier-management/party/proto"
	pb "gitlab.artin.ai/backend/courier-management/party/proto"
	"gitlab.artin.ai/backend/courier-management/uaa/security"
	"testing"
)

func init(){
	config.InitTestConfig()
	logger.InitLogger()
}

func TestProfile1(t *testing.T) {
	t.Skip()
	config := config.Party().Database
	db,err := db.NewOrm(config)
	if err != nil{
		logger.Fatalf("cannot connect to database", err)
	}
	id := "bc0b62c7-55cc-2657-54cb-71be785387f8"
	result := db.First(&domain.CourierUser{
		ID: id,
	})
	fmt.Printf("rows : %v ", result.RowsAffected)
	fmt.Printf("error : %v ", result.Error)

	fmt.Println("map")
	result2 := map[string]interface{}{}
	db.Model(&domain.CourierUser{}).First(&result2)
	fmt.Printf("result2 : %v ", result2)

	fmt.Println("single entity")
	u := domain.CourierUser{}
	db.Model(&domain.CourierUser{}).Where("id = ?", id).Scan(&u)
	fmt.Println(u)

	fmt.Println("list of entities with address")
	users2 := []domain.CourierUser{}
	db.Model(&domain.CourierUser{}).Joins("CourierAddress").Scan(&users2)
	fmt.Println(users2)

	fmt.Println("single entity with document with preload")
	u2 := domain.CourierUser{}
	db.Model(&domain.CourierUser{}).Preload("Documents").Where("id = ?", id).Find(&u2)
	fmt.Println(u2)

	fmt.Println("single entity with document with dirty join")
	rows, err := db.Model(&domain.CourierUser{}).Joins("JOIN document ON user.id = document.user_id").Select("user.id, user.first_name, user.last_name, user.email, user.phone_number, user.birth_date, user.status, user.transport_type, document.object_id, document.document_info_type, document.document_type, document.file_type").Rows()
	defer rows.Close()
	u3 := &domain.CourierUser{}
	u3.Documents = make([]domain.Document, 0)
	for rows.Next() {
		d := domain.Document{}
		err = rows.Scan(&u3.ID, &u3.FirstName, &u3.LastName, &u3.Email, &u3.PhoneNumber, &u3.BirthDate, &u3.Status, &u3.TransportType, &d.ObjectId, &d.DocumentInfoType, &d.DocumentType, &d.FileType)
		if err != nil {
			fmt.Printf("error %e", err)
		}
		u3.Documents = append(u3.Documents, d)
	}
	fmt.Println(u3)

	fmt.Println("search not existing user")
	var count int64
	u4 := domain.CourierUser{}
	db.Model(&domain.CourierUser{}).Where("id = ?", id).Find(&u4).Count(&count)
	fmt.Println(count)
}

func TestProfile2(t *testing.T) {
	t.Skip()
	service := NewService(config.GetData(), config.Jwt())
	ctx := context.WithValue(context.Background(), "user", security.User{
		Id: "e63c0b87-098a-4927-94e6-1bd817c48e5d",
		PhoneNumber: "989126031724",
		Name: "Behnam Nikbakht",
		Roles: []security.Role{security.Role_COURIER},
	})
	/*response, err := service.CreateCourierAccount(ctx, &pb.CreateCourierAccountRequest{
		FirstName: "Behnam",
		LastName: "Nikbakht",
		Email: "behnam.nikbakht@gmail.com",
		BirthDate: "1985",
		City: "Tehran",
	})

	/*response, err := service.UpdateCourierAccount(ctx, &pb.UpdateCourierAccountRequest{
		City: "Tehran2",
		FirstName: "Behnam2",
	})*/

	//response, err := service.GetIdCard("e63c0b87-098a-4927-94e6-1bd817c48e5d")
	//fmt.Printf("response = %v , err = %v", response, err)

	response, err := service.UpdateProfileAdditionalInfo(ctx, &pb.UpdateProfileAdditionalInfoRequest{
		Info: &pb.UpdateProfileAdditionalInfoRequest_IdCard{
			IdCard: &pb.IDCard{
				FirstName: "beh12",
				LastName: "nik1",
				Number: "num1",
				ExpirationDate: "exp1",
				IssuePlace: "isp1",
			},
		},
	})
	fmt.Printf("response = %v , err = %v", response, err)

	service.UpdateProfileAdditionalInfo(ctx, &pb.UpdateProfileAdditionalInfoRequest{
		Info: &pb.UpdateProfileAdditionalInfoRequest_DrivingLicense{
			DrivingLicense: &pb.DrivingLicense{
				DrivingLicenseNumber: "dn",
			},
		},
	})

	service.UpdateProfileAdditionalInfo(ctx, &pb.UpdateProfileAdditionalInfoRequest{
		Info: &pb.UpdateProfileAdditionalInfoRequest_DriverBackground{
			DriverBackground: &pb.DriverBackground{
				NationalInsuranceNumber: "nin",
			},
		},
	})

	service.UpdateProfileAdditionalInfo(ctx, &pb.UpdateProfileAdditionalInfoRequest{
		Info: &pb.UpdateProfileAdditionalInfoRequest_ResidenceCard{
			ResidenceCard: &pb.ResidenceCard{
				Number: "n1",
				ExpirationDate: "ed1",
				IssueDate: "id1",
			},
		},
	})

	service.UpdateProfileAdditionalInfo(ctx, &pb.UpdateProfileAdditionalInfoRequest{
		Info: &pb.UpdateProfileAdditionalInfoRequest_BankAccount{
			BankAccount: &pb.BankAccount{
				BankName: "bcn1",
				AccountNumber: "an1",
				AccountHolderName: "ahn1",
				SortCode: "sc1",
			},
		},
	})

	service.UpdateProfileAdditionalInfo(ctx, &pb.UpdateProfileAdditionalInfoRequest{
		Info: &pb.UpdateProfileAdditionalInfoRequest_Address{
			Address: &pb.Address{
				Street: "st1",
				Building: "b1",
				City: "c1",
				County: "co1",
				PostCode: "pc1",
			},
		},
	})

	response2, _ := service.GetProfileAdditionalInfo(ctx, &pb.GetProfileAdditionalInfoRequest{
		Type: pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ID_CARD,
	})
	fmt.Printf("idcard = %v", response2)

	response2, _ = service.GetProfileAdditionalInfo(ctx, &pb.GetProfileAdditionalInfoRequest{
		Type: pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVING_LICENSE,
	})
	fmt.Printf("drivers license = %v", response2)

	response2, _ = service.GetProfileAdditionalInfo(ctx, &pb.GetProfileAdditionalInfoRequest{
		Type: pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_DRIVER_BACKGROUND,
	})
	fmt.Printf("drivers background = %v", response2)

	response2, _ = service.GetProfileAdditionalInfo(ctx, &pb.GetProfileAdditionalInfoRequest{
		Type: pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_RESIDENCE_CARD,
	})
	fmt.Printf("residence card = %v", response2)

	response2, _ = service.GetProfileAdditionalInfo(ctx, &pb.GetProfileAdditionalInfoRequest{
		Type: pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_BANK_ACCOUNT,
	})
	fmt.Printf("bank account = %v", response2)

	response2, _ = service.GetProfileAdditionalInfo(ctx, &pb.GetProfileAdditionalInfoRequest{
		Type: pb.AdditionalInfoType_ADDITIONAL_INFO_TYPE_ADDRESS,
	})
	fmt.Printf("address = %v", response2)

	response3,_ := service.FindAccount(ctx, &pb.FindAccountRequest{
		Filter: &pb.FindAccountRequest_UserId{
			UserId: "e63c0b87-098a-4927-94e6-1bd817c48e5d",
		},
	})
	fmt.Printf("FindCourierAccount by id = %v", response3)

	response3,_ = service.FindAccount(ctx, &pb.FindAccountRequest{
		Filter: &pb.FindAccountRequest_PhoneNumber{
			PhoneNumber: "989126031724",
		},
	})
	fmt.Printf("FindCourierAccount by phonenumber = %v", response3)

	response4,err := service.GetCourierAccount(ctx, &pb.GetCourierAccountRequest{
	})
	fmt.Printf("GetCourierAccount = %v, err = %v", response4, err)

	response5,_ := service.FindCourierAccounts(ctx, &pb.FindCourierAccountsRequest{
		Filter: &pb.FindCourierAccountsRequest_UserId{
			UserId: "e63c0b87-098a-4927-94e6-1bd817c48e5d",
		},
	})
	fmt.Printf("FindCourierAccounts by userid = %v", response5)

	response5,_ = service.FindCourierAccounts(ctx, &pb.FindCourierAccountsRequest{
		Filter: &pb.FindCourierAccountsRequest_Name{
			Name: "behnam",
		},
	})
	fmt.Printf("FindCourierAccounts by name = %v", response5)

	/*service.Upload(ctx, &pb.UploadDocumentRequest{
		Document: &pb.Document{
			InfoType: pb.DocumentInfoType_DOCUMENT_INFO_TYPE_ID_CARD,
			DocType: pb.DocumentType_DOCUMENT_TYPE_PASSPORT,
			FileType: "pdf",
			Data: []byte("cGFzc3BvcnQgZG9jdW1lbnQK"),
		},
	})*/

	/*service.Delete(ctx, &pb.DeleteDocumentRequest{
		ObjectId: "58c107f5-fa42-557d-feac-5d621af62850",
	})*/

	response9,err := service.DirectDownload(ctx, &pb.DirectDownloadRequest{
		ObjectId: "9fc8360d-468b-2065-0cf4-a2dcfcf2d7ea",
	})
	fmt.Printf("DirectDownload = %v, err = %v", response9, err)


}