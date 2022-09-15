package db

import (
	"database/sql"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
	"os"

	"gitlab.artin.ai/backend/courier-management/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)


func NewOrm(config config.Database) (*gorm.DB, error) {
	loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.DatabaseHost + ":" + config.DatabasePort + ")/" + config.DatabaseName + "?parseTime=true"
	conn, err := sql.Open(config.DatabaseDriver, dsn)
	conn = sqldblogger.OpenDriver(dsn, conn.Driver(), loggerAdapter)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: conn,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt: true,
	})

	dbDB, err := db.DB()
	dbDB.SetMaxIdleConns(config.MaxIdleConnections)
	dbDB.SetMaxOpenConns(config.MaxOpenConnections)
	dbDB.SetConnMaxIdleTime(time.Duration(config.MaxIdleTimeInMinutes) * time.Minute)
	dbDB.SetConnMaxLifetime(time.Duration(config.MaxLifetimeInMinutes) * time.Minute)

	return db,err
}