package domain

type Device struct {
	DeviceId  	         	string				`gorm:"type:string;size:128;primaryKey;not null;column:device_id"`
	PhoneNumber           	string				`gorm:"type:string;size:64;not null;column:phone_number"`
	Manufacturer           	string				`gorm:"type:string;size:64;not null;column:manufacturer"`
	DeviceModel           	string				`gorm:"type:string;size:64;not null;column:device_model"`
	DeviceOs    	       	int32				`gorm:"type:string;size:64;not null;column:device_os"`
	DeviceVersion           string				`gorm:"type:string;size:64;not null;column:device_version"`
	DeviceToken           	string				`gorm:"type:string;size:255;not null;column:device_token"`
}

func (Device) TableName() string {
	return "device"
}