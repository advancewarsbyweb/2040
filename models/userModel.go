package models

import "database/sql"

type User struct {
	ID        int
	Email     string
	Password  string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
