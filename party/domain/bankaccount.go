package domain

type BankAccount struct {
	UserID           				string				`gorm:"type:string;size:64;primaryKey;not null;column:user_id"`
	BankName       					string				`gorm:"type:string;size:64;not null;column:bank_name"`
	AccountNumber       			string				`gorm:"type:string;size:64;not null;column:account_number"`
	AccountHolderName       		string				`gorm:"type:string;size:64;not null;column:account_holder_name"`
	SortCode       					string				`gorm:"type:string;size:64;not null;column:sort_code"`
}

func (BankAccount) TableName() string {
	return "courier_bank_account"
}