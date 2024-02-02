package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Press struct {
	ID         int         `db:"press_id,omitempty"`
	text       null.String `db:"press_text" json:"text" validate:"required"`
	PlayerID   int         `db:"press_players_id" json:"playerId" validate:"required,numeric"`
	Date       time.Time   `db:"press_date" json:"date" validate:"omitempty"`
	subject    null.String `db:"press_subject" json:"subject" validate:"required"`
	Type       string      `db:"press_type"`
	recipients null.String `db:"recipients"`
}

func (p *Press) Text() string {
	return p.text.String
}

func (p *Press) Subject() string {
	return p.subject.String
}

func (p *Press) Recipients() string {
	return p.recipients.String
}
