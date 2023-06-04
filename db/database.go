package db

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func NewDatabase() {
	DB = ConnectDatabase(os.Getenv("DSN"))
}

func ConnectDatabase(dsn string) *sqlx.DB {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
	return db
}
