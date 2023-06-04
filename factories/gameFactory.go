package factories

import (
	"database/sql"
	"time"

	"github.com/awbw/2040/models"
	"github.com/bxcodec/faker/v4"
)

type GameFactory struct{}

var Game GameFactory

func NewGameFactory() GameFactory {
	return GameFactory{}
}

func init() {
	Game = NewGameFactory()
}

func (f *GameFactory) Create() models.Game {
	return models.Game{
		Name:          faker.Word(),
		Password:      "",
		StartDate:     sql.NullTime{Time: time.Now()},
		EndDate:       sql.NullTime{},
		CreatorID:     1,
		ActivityDate:  sql.NullTime{Time: time.Now()},
		MapID:         1,
		WeatherType:   "Clear",
		WeatherCode:   "C",
		WeatherStart:  sql.NullInt16{},
		Private:       sql.NullString{String: "N"},
		TurnID:        0,
		Day:           1,
		Active:        "N",
		Funds:         1000,
		CaptureLimit:  1000,
		DayLimit:      0,
		Fog:           "N",
		Comment:       faker.Paragraph(),
		Type:          "N",
		BootInterval:  -1,
		StartingFunds: 0,
		Official:      "N",
		MinRating:     0,
		UnitLimit:     50,
		League:        "N",
		Team:          "N",
		AETInterval:   sql.NullInt16{},
		AETDate:       sql.NullTime{},
		UsePowers:     sql.NullString{},
	}
}
