package domain

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"strings"
)

type RateItem struct {
	ID                   int64                		 `gorm:"type:bigint;size:64;primaryKey;not null;column:id"`
	Rater     	     	 string               		 `gorm:"type:string;size:64;not null;column:rater"`
	Rated  	         	 string               		 `gorm:"type:string;size:64;not null;column:rated"`
	Ride	             string               		 `gorm:"type:string;size:64;not null;column:ride"`
	RaterType     	     RateeType               	 `gorm:"type:tinyint;not null;column:rater_type"`
	RatedType  	         RateeType               	 `gorm:"type:tinyint;not null;column:rated_type"`
	RateValue            int32                		 `gorm:"type:tinyint;default:0;not null;column:rate_value"`
	Message              sql.NullString    			 `gorm:"type:string;size:64;column:message"`
	Feedbacks            []RateItemFeedback       `gorm:"foreignKey:RateId;references:ID"`
}

func (RateItem) TableName() string {
	return "rate_item"
}

type RateeType int32

const (
	RATEE_TYPE_UNKNOWN RateeType = iota
	RATEE_TYPE_COURIER
	RATEE_TYPE_CLIENT
)

func (s RateeType) String() string {
	return raterTypeToString[s]
}

var raterTypeToString = map[RateeType]string{
	RATEE_TYPE_UNKNOWN:  "",
	RATEE_TYPE_COURIER:  "courier",
	RATEE_TYPE_CLIENT: "client",
}

var raterTypeToID = map[string]RateeType{
	"":  RATEE_TYPE_UNKNOWN,
	"courier":  RATEE_TYPE_COURIER,
	"client": RATEE_TYPE_CLIENT,
}

func (s RateeType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(raterTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *RateeType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = raterTypeToID[strings.ToLower(j)]
	return nil
}