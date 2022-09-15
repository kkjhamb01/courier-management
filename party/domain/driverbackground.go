package domain

type DriverBackground struct {
	UserID           				string				`gorm:"type:string;size:64;primaryKey;not null;column:user_id"`
	NationalInsuranceNumber       	string				`gorm:"type:string;size:64;not null;column:national_insurance_number"`
	UploadDbsLater       			int32				`gorm:"type:tinyint;column:upload_dbs_later"`
}

func (DriverBackground) TableName() string {
	return "courier_driver_background"
}