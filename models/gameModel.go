package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	StartDate    sql.NullTime
	EndDate      sql.NullTime
	CreatorID    int
	Creator      User `gorm:"foreignKey: CreatorID"`
	Players      []Player
	ActivityDate sql.NullTime
	MapID        int
	Map          Map `gorm:"foreignKey: MapID"`
	WeatherType  string
	WeatherStart int
	WeatherCode  string
	Private      bool
	TurnID       int
	Turn         Player `gorm:"foreignKey: TurnID"`
	Day          int
	Active       bool
	Funds        int
	CaptureLimit int
	DayLimit     int
	Fog          bool
	Comment      string
	Type         string
	Official     bool
	MinRating    int
	UnitLimit    int
	League       bool
	Team         bool
}
