package domain

type AnnouncementUser struct {
	AnnouncementId             	int64                		 `gorm:"primaryKey;not null;column:announcement_id"`
	UserId                		string                		 `gorm:"primaryKey;not null;column:user_id"`
	Announcement  	         	Announcement
}

func (AnnouncementUser) TableName() string {
	return "announcement_user"
}