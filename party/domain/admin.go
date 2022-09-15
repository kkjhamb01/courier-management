package domain

import "database/sql"

type AdminUser struct {
	UserID          	 	string				`gorm:"type:string;size:64;primaryKey;not null;column:id"`
	FirstName           	sql.NullString		`gorm:"type:string;size:64;column:first_name"`
	LastName    			sql.NullString		`gorm:"type:string;size:64;column:last_name"`
	Username    			string				`gorm:"type:string;size:64;column:username"`
	Password    			string				`gorm:"type:string;size:64;column:password"`
}

func (AdminUser) TableName() string {
	return "admin"
}