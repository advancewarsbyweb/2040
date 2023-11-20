package db

import (
	"time"

	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types"
	"github.com/bxcodec/faker/v4"
	"gopkg.in/guregu/null.v4"
)

type gameFactory struct {
	Game models.Game
}

var Game gameFactory

func NewGameFactory() gameFactory {
	return gameFactory{}
}

func init() {
	Game = NewGameFactory()
}

func (f *gameFactory) Create() *gameFactory {
	f.Game = models.Game{
		Name:          faker.Word(),
		Password:      "",
		CreatorID:     1,
		MapID:         1,
		WeatherType:   "Clear",
		WeatherCode:   "C",
		TurnID:        0,
		Day:           1,
		StartDate:     null.TimeFrom(time.Now()),
		EndDate:       null.Time{},
		ActivityDate:  null.TimeFrom(time.Now()),
		WeatherStart:  null.Int{},
		Private:       null.String{},
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
		UsePowers:     null.String{},
	}
	return f
}

func (f *gameFactory) Build() models.Game {
	return f.Game
}

func (f *gameFactory) BuildInsert() models.Game {
	gID, _ := GameRepo.CreateGame(types.NewGame(f.Game))
	f.Game.ID = gID
	return f.Game
}
