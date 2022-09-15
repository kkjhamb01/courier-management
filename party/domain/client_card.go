package domain

type ClientCard struct {
	UserID           	string				`gorm:"type:string;size:64;primaryKey;not null;column:user_id"`
	CardNumber          string				`gorm:"type:string;size:64;column:card_number"`
	IssueDate    		string				`gorm:"type:string;size:64;column:issue_date"`
	CVV		    		string				`gorm:"type:string;size:64;column:cvv"`
	ZipCode	    		string				`gorm:"type:string;size:64;column:zip_code"`
	Country	    		string				`gorm:"type:string;size:64;column:country"`
}

func (ClientCard) TableName() string {
	return "client_card"
}