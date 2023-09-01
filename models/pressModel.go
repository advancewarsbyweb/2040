package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Press struct {
	ID         int         `db:"press_id,omitempty"`
	Text       null.String `db:"press_text"`
	PlayerID   int         `db:"press_players_id"`
	Date       time.Time   `db:"press_date"`
	Subject    null.String `db:"press_subject"`
	Type       string      `db:"press_type"`
	Recipients null.String `db:"recipients"`
}
