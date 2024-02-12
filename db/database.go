package db

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var DB *sqlx.DB

func NewTestDatabase() {
	godotenv.Load("../.env")
	DB = ConnectDatabase(os.Getenv("TEST_DSN"))
}

func NewDatabase(dsnEnv string) {
	DB = ConnectDatabase(os.Getenv(dsnEnv))
}

func ConnectDatabase(dsn string) *sqlx.DB {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(time.Minute * 4) // <-- this
	return db
}

func FormatCreateQuery(table string, columns []string) string {
	var (
		insertCols []string
		values     []string
	)

	// Format columns to insert with the values from the list above to make sure they always match
	for _, col := range columns {
		insertCols = append(insertCols, col)
		values = append(values, fmt.Sprintf(":%s", col))
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(insertCols, ","), strings.Join(values, ","))
}
