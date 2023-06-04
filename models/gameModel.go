package models

import (
	"database/sql"
)

type Game struct {
	ID           int
	Name         string
	StartDate    sql.NullTime
	EndDate      sql.NullTime
	CreatorID    int
	Creator      *User `gorm:"foreignKey: CreatorID"`
	Players      []Player
	ActivityDate sql.NullTime
	MapID        int
	Map          Map `gorm:"foreignKey: MapID"`
	WeatherType  string
	WeatherStart int
	WeatherCode  string
	Private      bool
	TurnID       int
	Turn         *Player `gorm:"foreignKey: TurnID"`
	Day          int
	Active       bool
	Funds        int
	CaptureLimit int
	DayLimit     int
	Fog          string
	Comment      string
	Type         string
	Official     string
	MinRating    int
	UnitLimit    int
	League       string
	Team         string
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
	DeletedAt    sql.NullTime
}
