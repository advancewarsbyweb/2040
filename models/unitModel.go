package models

import (
	"database/sql"
)

type Unit struct {
	ID       int
	PlayerID int
	Player   Player
	GameID   int
	Game     Game
	X        int
	Y        int
	UnitType
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
