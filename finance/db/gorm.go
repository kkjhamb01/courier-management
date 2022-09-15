package db

import (
	"fmt"
	"gitlab.artin.ai/backend/courier-management/common/config"
	"gitlab.artin.ai/backend/courier-management/common/logger"
	"gitlab.artin.ai/backend/courier-management/common/logger/tag"
	"gitlab.artin.ai/backend/courier-management/finance/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
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
	err := MariaDbClient().AutoMigrate(&model.Account{})
	panicIfError(err, model.Account{})

	err = MariaDbClient().AutoMigrate(&model.AccountRole{})
	panicIfError(err, model.AccountRole{})

	err = MariaDbClient().AutoMigrate(&model.Transaction{})
	panicIfError(err, model.Transaction{})

	err = MariaDbClient().AutoMigrate(&model.PaymentMethod{})
	panicIfError(err, model.PaymentMethod{})
}

func InitSeedData() error {
	insertedTaxAccounts := int64(0)
	err := MariaDbClient().Model(&model.Account{}).
		Where("type = ?", model.AccountTypeTax).
		Count(&insertedTaxAccounts).Error
	if err != nil {
		return err
	}

	insertedRevenueAccounts := int64(0)
	err = MariaDbClient().Model(&model.Account{}).
		Where("type = ?", model.AccountTypeRevenue).
		Count(&insertedRevenueAccounts).Error
	if err != nil {
		return err
	}

	insertedArtinAccounts := int64(0)
	err = MariaDbClient().Model(&model.Account{}).
		Where("type = ?", model.AccountTypeArtin).
		Count(&insertedArtinAccounts).Error
	if err != nil {
		return err
	}

	if insertedTaxAccounts == 0 {
		accountRoles := []model.AccountRole{
			{
				FromDate: time.Unix(0, 0),
				Status:   model.AccountRoleStatusActive,
				Type:     model.AccountRoleTypeAdmin,
			},
		}

		var taxAccount = model.Account{
			Roles:  accountRoles,
			Status: model.AccountStatusOpen,
			Type:   model.AccountTypeTax,
		}
		err = MariaDbClient().Create(&taxAccount).Error
		if err != nil {
			logger.Error("failed to save tax account", err)
			return err
		}
	}

	if insertedRevenueAccounts == 0 {
		accountRoles := []model.AccountRole{
			{
				FromDate: time.Unix(0, 0),
				Status:   model.AccountRoleStatusActive,
				Type:     model.AccountRoleTypeAdmin,
			},
		}

		var revenueAccount = model.Account{
			Roles:  accountRoles,
			Status: model.AccountStatusOpen,
			Type:   model.AccountTypeRevenue,
		}
		err = MariaDbClient().Create(&revenueAccount).Error
		if err != nil {
			logger.Error("failed to save revenue account", err)
			return err
		}
	}

	if insertedArtinAccounts == 0 {
		accountRoles := []model.AccountRole{
			{
				FromDate: time.Unix(0, 0),
				Status:   model.AccountRoleStatusActive,
				Type:     model.AccountRoleTypeAdmin,
			},
		}

		var artinAccount = model.Account{
			Roles:  accountRoles,
			Status: model.AccountStatusOpen,
			Type:   model.AccountTypeArtin,
		}
		err = MariaDbClient().Create(&artinAccount).Error
		if err != nil {
			logger.Error("failed to save artin account", err)
			return err
		}
	}

	return nil
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
	dsn += "/" + config.Finance().DbName
	dsn += "?parseTime=true"

	return dsn
}
