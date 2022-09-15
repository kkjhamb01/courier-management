package domain

import (
	"bytes"
	"encoding/json"
	"strings"
)

type RateItemFeedback struct {
	RateId           	int64		`gorm:"type:bigint;size:64;primaryKey;not null;column:rate_id"`
	Feedback  	        int			`gorm:"type:int;primaryKey;not null;column:feedback"`
	Positive  	        FeedbackType			`gorm:"type:int;not null;column:positive"`
}

func (RateItemFeedback) TableName() string {
	return "rate_item_feedback"
}

type FeedbackType int32

const (
	FEEDBACK_TYPE_UNKNOWN FeedbackType = iota
	FEEDBACK_TYPE_POSITIVE
	FEEDBACK_TYPE_NEGATIVE
)

func (s FeedbackType) String() string {
	return feedbackToString[s]
}

var feedbackToString = map[FeedbackType]string{
	FEEDBACK_TYPE_UNKNOWN:  "",
	FEEDBACK_TYPE_POSITIVE:  "positive",
	FEEDBACK_TYPE_NEGATIVE: "negative",
}

var feedbackToID = map[string]FeedbackType{
	"":  FEEDBACK_TYPE_UNKNOWN,
	"positive":  FEEDBACK_TYPE_POSITIVE,
	"negative": FEEDBACK_TYPE_NEGATIVE,
}

func (s FeedbackType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(feedbackToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *FeedbackType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = feedbackToID[strings.ToLower(j)]
	return nil
}
