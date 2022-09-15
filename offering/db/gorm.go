package db

import (
	"fmt"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/offering/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var mariaDb *gorm.DB
var mariaDbSetupOnce sync.Once

func SetupMariaDbClient() {
	mariaDbSetupOnce.Do(func() {
		var err error
		mariaDb, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                     mariaDbDsn(),
			DontSupportRenameIndex:  true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		}), &gorm.Config{})
		if err != nil {
			logger.Fatal("failed to connect to mariadb", tag.Err("err", err))
			panic("failed to connect to the database")
		}

		logger.Info("created connection to mariadb (Not guaranteed as pinging is disabled)")
	})
}

func MariaDbClient() *gorm.DB {
	return mariaDb
}

func MigrateGormIntoMariaDb() {
	err := MariaDbClient().AutoMigrate(&model.CourierStatusLog{})
	panicIfError(err, model.CourierStatusLog{})

	err = MariaDbClient().AutoMigrate(&model.AcceptedOffer{})
	panicIfError(err, model.AcceptedOffer{})
}

func panicIfError(err error, model interface{}) {
	if err != nil {
		logger.Fatal("failed to migrate", tag.Obj("model", fmt.Sprintf("%T", model)))
		panic(err)
	}
}

func mariaDbDsn() string {
	var dsn string

	dsn = config.MariaDb().Username
	dsn += ":" + config.MariaDb().Password
	dsn += "@" + config.MariaDb().Protocol
	dsn += "(" + config.MariaDb().Address + ")"
	dsn += "/" + config.Offering().DbName
	dsn += "?parseTime=true"

	return dsn
}
