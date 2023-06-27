package factories

import (
	"time"

	"github.com/awbw/2040/models"
	"github.com/bxcodec/faker/v4"
	"gopkg.in/guregu/null.v4"
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
		StartDate:     null.TimeFrom(time.Now()),
		EndDate:       null.Time{},
		CreatorID:     1,
		ActivityDate:  null.TimeFrom(time.Now()),
		MapID:         1,
		WeatherType:   "Clear",
		WeatherCode:   "C",
		WeatherStart:  null.Int{},
		Private:       "N",
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
		AETInterval:   null.Int{},
		AETDate:       null.Int{},
		UsePowers:     "Y",
	}
}
