package domain

type IDCard struct {
	UserID           	string				`gorm:"type:string;size:64;primaryKey;not null;column:user_id"`
	FirstName           string				`gorm:"type:string;size:64;not null;column:first_name"`
	LastName           	string				`gorm:"type:string;size:64;not null;column:last_name"`
	Number           	string				`gorm:"type:string;size:64;not null;column:number"`
	ExpirationDate      string				`gorm:"type:string;size:64;not null;column:expiration_date"`
	IssuePlace          string				`gorm:"type:string;size:64;not null;column:issue_place"`
	Type          		int32				`gorm:"type:int;not null;column:type"`
}

func (IDCard) TableName() string {
	return "courier_id_card"
}