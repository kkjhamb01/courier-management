package domain

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"strings"
)

type PromotionUser struct {
	PromotionId             int64                		 	`gorm:"primaryKey;not null;column:promotion_id"`
	UserId                	string                		 	`gorm:"primaryKey;not null;column:user_id"`
	Metadata                sql.NullString                	`gorm:"column:metadata"`
	Status  	         	PromotionStatus             	`gorm:"column:status"`
	Promotion  	         	Promotion
}

func (PromotionUser) TableName() string {
	return "promotion_user"
}


type PromotionStatus int32

const (
	PROMOTION_STATUS_UNKNOWN PromotionStatus = iota
	PROMOTION_STATUS_AVAILABLE
	PROMOTION_STATUS_CONSUMED
	PROMOTION_STATUS_NOT_AVAILABLE
)

func (s PromotionStatus) String() string {
	return promotionStatusToString[s]
}

var promotionStatusToString = map[PromotionStatus]string{
	PROMOTION_STATUS_UNKNOWN:  "",
	PROMOTION_STATUS_AVAILABLE:  "available",
	PROMOTION_STATUS_CONSUMED: "consumed",
	PROMOTION_STATUS_NOT_AVAILABLE: "not available",
}

var promotionStatusToID = map[string]PromotionStatus{
	"":  PROMOTION_STATUS_UNKNOWN,
	"available":  PROMOTION_STATUS_AVAILABLE,
	"consumed": PROMOTION_STATUS_CONSUMED,
	"not available": PROMOTION_STATUS_NOT_AVAILABLE,
}

func (s PromotionStatus) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(promotionStatusToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *PromotionStatus) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = promotionStatusToID[strings.ToLower(j)]
	return nil
}