package domain

type DocumentData struct {
	ObjectId           	string			`gorm:"type:string;size:64;not null;unique;column:object_id"`
	Data    			[]byte			`gorm:"type:blob;not null;column:data"`
}

func (DocumentData) TableName() string {
	return "document_storage"
}