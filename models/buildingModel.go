package models

import "database/sql"

type Building struct {
	ID        int
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
