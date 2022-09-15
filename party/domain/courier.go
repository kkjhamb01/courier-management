package domain

import (
	"database/sql"
)

type CourierUser struct {
	ID                   string               `gorm:"type:string;size:64;primaryKey;not null;column:id"`
	FirstName            sql.NullString       `gorm:"type:string;size:64;default:'';column:first_name"`
	LastName             string               `gorm:"type:string;size:64;not null;column:last_name"`
	Email                sql.NullString       `gorm:"type:string;size:64;unique;column:email"`
	PhoneNumber          string               `gorm:"type:string;size:32;unique;not null;column:phone_number"`
	BirthDate            sql.NullString       `gorm:"type:string;size:32;column:birth_date"`
	Status               int32                `gorm:"type:int32;default:0;not null;column:status"`
	TransportType        sql.NullInt32        `gorm:"type:int32;column:transport_type"`
	TransportSize        sql.NullInt32        `gorm:"type:int32;column:transport_size"`
	Citizen              sql.NullInt32        `gorm:"type:int32;column:citizen"`
	Address              CourierAddress       `gorm:"foreignKey:UserID;references:ID"`
	Claims               []CourierClaim       `gorm:"foreignKey:UserID;references:ID"`
	IDCards              []IDCard             `gorm:"foreignKey:UserID;references:ID"`
	DrivingLicense       DrivingLicense       `gorm:"foreignKey:UserID;references:ID"`
	DriverBackground     DriverBackground     `gorm:"foreignKey:UserID;references:ID"`
	ResidenceCard        ResidenceCard        `gorm:"foreignKey:UserID;references:ID"`
	BankAccount          BankAccount          `gorm:"foreignKey:UserID;references:ID"`
	Documents            []Document           `gorm:"foreignKey:UserID;references:ID"`
	CourierStatusList    []CourierStatus      `gorm:"foreignKey:UserID;references:ID"`
	CourierMot           CourierMot           `gorm:"foreignKey:UserID;references:ID"`
}

func (CourierUser) TableName() string {
	return "courier"
}
