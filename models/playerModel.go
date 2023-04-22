package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	GameID int
	Game   *Game `gorm:"foreignKey: GameID"`
	UserID int
	User   *User `gorm:"foreignKey: UserID"`
}
