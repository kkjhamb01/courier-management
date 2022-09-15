package domain

import "database/sql"

type ClientAddress struct {
	UserID           	string				`gorm:"type:string;size:64;primaryKey;not null;column:user_id"`
	Street           	sql.NullString		`gorm:"type:string;size:64;column:street"`
	Building    		sql.NullString		`gorm:"type:string;size:64;column:building"`
	City    			sql.NullString		`gorm:"type:string;size:64;column:city"`
	County    			sql.NullString		`gorm:"type:string;size:64;column:county"`
	PostCode    		sql.NullString		`gorm:"type:string;size:64;column:post_code"`
	AddressDetails    	sql.NullString		`gorm:"type:string;size:64;column:address_details"`
}

func (ClientAddress) TableName() string {
	return "client_address"
}