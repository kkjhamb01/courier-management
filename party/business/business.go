package business

import (
	_ "database/sql"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/party/db"
	"gitlab.artin.ai/backend/courier-management/uaa/security"
	"gorm.io/gorm"
)

type Service struct {
	config                config.PartyData
	db					  *gorm.DB
	jwtUtils      		  *security.JWTUtils
	promotionApi      	  PromotionAPI
}

func NewService(config config.Data, jwtConfig config.JwtData) *Service {
	db,err := db.NewOrm(config.Party.Database)
	if err != nil{
		logger.Fatalf("cannot connect to database", err)
	}
	jwtUtils, err := security.NewJWTUtils(jwtConfig)
	if err != nil{
		logger.Fatalf("cannot create jwtutils ", err)
	}
	return &Service{
		config: config.Party,
		db: db,
		jwtUtils: &jwtUtils,
		promotionApi: NewPromotionAPI(config.Promotion),
	}
}