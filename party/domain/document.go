package domain

import (
	"database/sql"
	"time"
)

type Document struct {
	UserID           	string					`gorm:"type:string;size:64;primaryKey;not null;column:user_id"`
	ObjectId           	string					`gorm:"primaryKey;type:string;size:64;not null;unique;column:object_id"`
	DocumentInfoType    int						`gorm:"type:string;size:64;not null;column:document_info_type"`
	DocumentType    	int						`gorm:"type:string;size:64;not null;column:document_type"`
	FileType    		sql.NullString			`gorm:"type:string;size:64;column:file_type"`
	Data				DocumentData			`gorm:"foreignKey:ObjectId;references:ObjectId"`
	CreationTime        time.Time		        `gorm:"column:creation_time"`
}

func (Document) TableName() string {
	return "document"
}