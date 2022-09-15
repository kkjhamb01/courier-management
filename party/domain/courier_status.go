package domain

type CourierStatus struct {
	UserID           	string				`gorm:"type:string;size:64;primaryKey;not null;column:user_id"`
	StatusType          int32				`gorm:"type:int;primaryKey;not null;column:status_type"`
	Status          	int32				`gorm:"type:int;not null;column:status"`
	Message          	string				`gorm:"type:string;column:message"`
}

func (CourierStatus) TableName() string {
	return "courier_status"
}