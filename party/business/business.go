package business

import (
	_ "database/sql"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/kkjhamb01/courier-management/common/logger"
	"github.com/kkjhamb01/courier-management/party/db"
	"github.com/kkjhamb01/courier-management/uaa/security"
	"gorm.io/gorm"
)

type Service struct {
	config       config.PartyData
	db           *gorm.DB
	jwtUtils     *security.JWTUtils
	promotionApi PromotionAPI
}

func NewService(config config.Data, jwtConfig config.JwtData) *Service {
	db, err := db.NewOrm(config.Party.Database)
	if err != nil {
		logger.Fatalf("cannot connect to database", err)
	}
	jwtUtils, err := security.NewJWTUtils(jwtConfig)
	if err != nil {
		logger.Fatalf("cannot create jwtutils ", err)
	}
	return &Service{
		config:       config.Party,
		db:           db,
		jwtUtils:     &jwtUtils,
		promotionApi: NewPromotionAPI(config.Promotion),
	}
}
