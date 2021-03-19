package db

import (
	"log"

	_ "github.com/jmoiron/sqlx"
)

// Database driver choice
const (
	PostgreDriver int = iota
	MySQLDriver
	SQLiteDriver
	MSSQLDriver
)

// NewDBConnectionFactory to switch db driver at runtime
// TODO Pass DBConfig as parameter
func NewDBConnectionFactory(driver int) *ConnectTo {
	switch driver {
	case PostgreDriver:
		return newPostgresqlDBConnection()
	case MySQLDriver:
		// return newSQLiteDBConnection()
		return nil
	default:
		log.Fatal("Invalid database driver")
		return nil
	}
}
