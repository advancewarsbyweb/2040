package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB = db
}
