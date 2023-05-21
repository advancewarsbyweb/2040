package models

import "database/sql"

type Unit struct {
	ID        int
	PlayerID  int
	Player    Player `gorm:"foreignKey: PlayerID"`
	GameID    int
	Game      Game `gorm:"foreignKey: GameID"`
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
