package domain

type ResidenceCard struct {
	UserID           				string				`gorm:"type:string;size:64;primaryKey;not null;column:user_id"`
	Number       					string				`gorm:"type:string;size:64;not null;column:number"`
	ExpirationDate       			string				`gorm:"type:string;size:64;not null;column:expiration_date"`
	IssueDate       				string				`gorm:"type:string;size:64;not null;column:issue_date"`
}

func (ResidenceCard) TableName() string {
	return "courier_residence_card"
}