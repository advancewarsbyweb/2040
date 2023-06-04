package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../.env")
}

func NewTestDatabase() {
	DB = ConnectDatabase(os.Getenv("TEST_DSN"))
}

func TestConnectDatabase(t *testing.T) {
	db := ConnectDatabase(os.Getenv("TEST_DSN"))
	defer func() {
		db.Close()
	}()
	err := db.Ping()

	if err != nil {
		t.Errorf("Failed to ping database: %s", err.Error())
	}
}
