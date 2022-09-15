package domain

type ClientClaim struct {
	UserID           	string		`gorm:"type:string;size:64;primaryKey;not null;column:user_id"`
	ClaimType  	        int			`gorm:"type:int;primaryKey;not null;column:claim_type"`
	Identifier    		string		`gorm:"type:string;size:64;not null;column:identifier"`
}

func (ClientClaim) TableName() string {
	return "client_claim"
}