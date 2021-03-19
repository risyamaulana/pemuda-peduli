package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type ConnectTo struct {
	DBRead *sqlx.DB
	DBExec *sqlx.DB
}

// NewPostgresqlDBConnection ...
func newPostgresqlDBConnection() *ConnectTo {

	// env Load
	err := godotenv.Load(".sample.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	db := dbReadConn()
	db2 := dbExecConn()

	dbConn := ConnectTo{
		DBRead: db,
		DBExec: db2,
	}

	return &dbConn
}

func dbReadConn() *sqlx.DB {
	// Connect to database read
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", mustGetenv("POSTGRES_HOST"), mustGetenv("POSTGRES_PORT"), mustGetenv("POSTGRES_USER"), mustGetenv("POSTGRES_PASSWORD"), mustGetenv("POSTGRES_DBNAME"))
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln("error connection: ", err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database is unreachable: %s", err)
	}

	return db
}

func dbExecConn() *sqlx.DB {
	// Connect to database exect
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", mustGetenv("POSTGRES_HOST"), mustGetenv("POSTGRES_PORT"), mustGetenv("POSTGRES_USER"), mustGetenv("POSTGRES_PASSWORD"), mustGetenv("POSTGRES_DBNAME"))

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln("error connection: ", err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database is unreachable: %s", err)
	}

	return db
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}
