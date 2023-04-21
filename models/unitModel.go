package models

import "gorm.io/gorm"

type Unit struct {
	gorm.Model
	PlayerID int
	Player   Player `gorm:"foreignKey: PlayerID"`
	GameID   int
	Game     Game `gorm:"foreignKey: GameID"`
}
