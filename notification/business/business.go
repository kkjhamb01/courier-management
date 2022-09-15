package business

import (
	_ "database/sql"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/notification/db"
	"gorm.io/gorm"
)

type Service struct {
	config config.NotificationData
	db     *gorm.DB
}

func NewService(config config.NotificationData) *Service {
	db, err := db.NewOrm(config.Database)
	if err != nil {
		logger.Fatalf("cannot connect to database", err)
	}
	return &Service{
		config: config,
		db:     db,
	}
}
