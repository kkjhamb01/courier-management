package domain

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"strings"
)

type Promotion struct {
	Id                		int64                		 	`gorm:"primaryKey;not null;column:id"`
	Name                	string                		 	`gorm:"not null;column:name"`
	StartDate     	     	sql.NullTime               		`gorm:"column:start_date"`
	ExpDate     	     	sql.NullTime               		`gorm:"column:exp_date"`
	DiscountPercentage  	sql.NullFloat64               	`gorm:"column:discount_percentage"`
	DiscountValue  	        sql.NullFloat64               	`gorm:"column:discount_value"`
	Type  	         		PromotionType               	`gorm:"column:type"`
}

func (Promotion) TableName() string {
	return "promotion"
}

type PromotionType int32

const (
	PROMOTION_TYPE_UNKNOWN PromotionType = iota
	PROMOTION_TYPE_ALL
	PROMOTION_TYPE_REFERRAL
	PROMOTION_TYPE_GROUP
	PROMOTION_TYPE_INDIVIDUAL
	PROMOTION_TYPE_REFERRING
)

func (s PromotionType) String() string {
	return promotionTypeToString[s]
}

var promotionTypeToString = map[PromotionType]string{
	PROMOTION_TYPE_UNKNOWN:  "",
	PROMOTION_TYPE_ALL:  "all",
	PROMOTION_TYPE_REFERRAL: "referral",
	PROMOTION_TYPE_GROUP: "group",
	PROMOTION_TYPE_INDIVIDUAL: "individual",
}

var promotionTypeToID = map[string]PromotionType{
	"":  PROMOTION_TYPE_UNKNOWN,
	"all":  PROMOTION_TYPE_ALL,
	"referral": PROMOTION_TYPE_REFERRAL,
	"group": PROMOTION_TYPE_GROUP,
	"individual": PROMOTION_TYPE_INDIVIDUAL,
}

func (s PromotionType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(promotionTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *PromotionType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = promotionTypeToID[strings.ToLower(j)]
	return nil
}