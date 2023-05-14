package models

type Player struct {
	ID     int
	GameID int
	Game   *Game `gorm:"foreignKey: GameID"`
	UserID int
	User   *User `gorm:"foreignKey: UserID"`
}
