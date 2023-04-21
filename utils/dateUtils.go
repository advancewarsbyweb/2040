package utils

import (
	"database/sql"
	"time"
)

func NullTime() sql.NullTime {
	return sql.NullTime{Time: time.Now()}
}
