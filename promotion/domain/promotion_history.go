package domain

import "time"

type PromotionHistory struct {
	PromotionId             int64                		 	`gorm:"primaryKey;not null;column:promotion_id"`
	UserId                	string                		 	`gorm:"primaryKey;not null;column:user_id"`
	TransactionId  			string               		 	`gorm:"column:transaction_id"`
	Date		  	        time.Time               		 `gorm:"column:date"`
}

func (PromotionHistory) TableName() string {
	return "promotion_history"
}