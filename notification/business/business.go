package business

import (
	_ "database/sql"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/notification/db"
	"gorm.io/gorm"
)

type Service struct {
	config                config.NotificationData
	db					  *gorm.DB
}

func NewService(config config.NotificationData) *Service {
	db,err := db.NewOrm(config.Database)
	if err != nil{
		logger.Fatalf("cannot connect to database", err)
	}
	return &Service{
		config: config,
		db: db,
	}
}