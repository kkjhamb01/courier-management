package db

import (
	"database/sql"

	"github.com/kkjhamb01/courier-management/common/logger"
)

type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type TxFn func(Transaction) error

func Transactional(db *sql.DB, fn TxFn) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			logger.Debug("rollback transaction ")
			tx.Rollback()
		} else if err != nil {
			logger.Error("rollback transaction", err)
			tx.Rollback()
		} else {
			logger.Debug("commit transaction")
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
