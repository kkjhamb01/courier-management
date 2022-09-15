package domain

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"
)

type Announcement struct {
	Id                		int64                		 	`gorm:"primaryKey;not null;column:id"`
	Title                	string                		 	`gorm:"not null;column:title"`
	Text                	string                		 	`gorm:"not null;column:text"`
	Type  	         		AnnouncementType               	`gorm:"column:type"`
	MessageType  	        AnnouncementMessageType         `gorm:"column:message_type"`
	CreationTime            time.Time                		`gorm:"column:creation_time"`
}

func (Announcement) TableName() string {
	return "announcement"
}

type AnnouncementType int32

const (
	ANNOUNCEMENT_TYPE_UNKNOWN AnnouncementType = iota
	ANNOUNCEMENT_TYPE_ALL
	ANNOUNCEMENT_TYPE_GROUP
	ANNOUNCEMENT_TYPE_INDIVIDUAL
)

func (s AnnouncementType) String() string {
	return announcementTypeToString[s]
}

var announcementTypeToString = map[AnnouncementType]string{
	ANNOUNCEMENT_TYPE_UNKNOWN:  "",
	ANNOUNCEMENT_TYPE_ALL:  "all",
	ANNOUNCEMENT_TYPE_GROUP: "group",
	ANNOUNCEMENT_TYPE_INDIVIDUAL: "individual",
}

var announcementTypeToID = map[string]AnnouncementType{
	"":  ANNOUNCEMENT_TYPE_UNKNOWN,
	"all":  ANNOUNCEMENT_TYPE_ALL,
	"group": ANNOUNCEMENT_TYPE_GROUP,
	"individual": ANNOUNCEMENT_TYPE_INDIVIDUAL,
}

func (s AnnouncementType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(announcementTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *AnnouncementType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = announcementTypeToID[strings.ToLower(j)]
	return nil
}




type AnnouncementMessageType int32

const (
	ANNOUNCEMENT_MESSAGE_TYPE_UNKNOWN AnnouncementMessageType = iota
	ANNOUNCEMENT_MESSAGE_TYPE_INFORMATION
	ANNOUNCEMENT_MESSAGE_TYPE_PROMOTION
	ANNOUNCEMENT_MESSAGE_TYPE_ALARM
)

func (s AnnouncementMessageType) String() string {
	return announcementMessageTypeToString[s]
}

var announcementMessageTypeToString = map[AnnouncementMessageType]string{
	ANNOUNCEMENT_MESSAGE_TYPE_UNKNOWN:  "",
	ANNOUNCEMENT_MESSAGE_TYPE_INFORMATION:  "information",
	ANNOUNCEMENT_MESSAGE_TYPE_PROMOTION: "promotion",
	ANNOUNCEMENT_MESSAGE_TYPE_ALARM: "alarm",
}

var announcementMessageTypeToID = map[string]AnnouncementMessageType{
	"":  ANNOUNCEMENT_MESSAGE_TYPE_UNKNOWN,
	"information":  ANNOUNCEMENT_MESSAGE_TYPE_INFORMATION,
	"promotion": ANNOUNCEMENT_MESSAGE_TYPE_PROMOTION,
	"alarm": ANNOUNCEMENT_MESSAGE_TYPE_ALARM,
}

func (s AnnouncementMessageType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(announcementMessageTypeToString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (s *AnnouncementMessageType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = announcementMessageTypeToID[strings.ToLower(j)]
	return nil
}