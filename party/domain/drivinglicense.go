package domain

type DrivingLicense struct {
	UserID           			string				`gorm:"type:string;size:64;primaryKey;not null;column:user_id"`
	DrivingLicenseNumber       	string				`gorm:"type:string;size:64;not null;column:driving_license_number"`
	ExpirationDate       		string				`gorm:"type:string;size:64;not null;column:expiration_date"`
}

func (DrivingLicense) TableName() string {
	return "courier_driving_license"
}