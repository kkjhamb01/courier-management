package domain

import (
	"database/sql"
)

type ClientUser struct {
	ID           			string					`gorm:"type:string;size:64;primaryKey;not null;column:id"`
	FirstName    			sql.NullString			`gorm:"type:string;size:64;default:'';column:first_name"`
	LastName     			string					`gorm:"type:string;size:64;not null;column:last_name"`
	Email        			sql.NullString			`gorm:"type:string;size:64;unique;column:email"`
	PhoneNumber  			string					`gorm:"type:string;size:32;unique;not null;column:phone_number"`
	Status       			int32					`gorm:"type:int32;default:0;not null;column:status"`
	PaymentMethod			sql.NullInt32			`gorm:"type:int32;column:payment_method"`
	ClientCards				[]ClientCard			`gorm:"foreignKey:UserID;references:ID"`
	Claims					[]ClientClaim			`gorm:"foreignKey:UserID;references:ID"`
	BirthDate        		sql.NullString   		`gorm:"type:string;size:32;column:birth_date"`
	Address           		ClientAddress   		`gorm:"foreignKey:UserID;references:ID"`
	Referral     			string					`gorm:"type:string;size:16;not null;column:referral"`
}

func (ClientUser) TableName() string {
	return "client"
}