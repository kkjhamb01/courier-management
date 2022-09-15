package db

import (
	"database/sql"
	"os"
	"time"

	"github.com/kkjhamb01/courier-management/common/config"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewOrm(config config.Database) (*gorm.DB, error) {
	loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.DatabaseHost + ":" + config.DatabasePort + ")/" + config.DatabaseName

	conn, err := sql.Open(config.DatabaseDriver, dsn)
	if err != nil {
		panic(err)
	}
	conn = sqldblogger.OpenDriver(dsn, conn.Driver(), loggerAdapter)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: conn,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	dbDB, err := db.DB()
	dbDB.SetMaxIdleConns(config.MaxIdleConnections)
	dbDB.SetMaxOpenConns(config.MaxOpenConnections)
	dbDB.SetConnMaxIdleTime(time.Duration(config.MaxIdleTimeInMinutes) * time.Minute)
	dbDB.SetConnMaxLifetime(time.Duration(config.MaxLifetimeInMinutes) * time.Minute)

	return db, err
}
