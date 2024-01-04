package db

import (
	"time"

	"github.com/awbw/2040/models"
	"github.com/bxcodec/faker/v4"
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
	}

	f.Game.SetStartDate(time.Now()).
		SetActivityDate(time.Now()).
		SetWeatherStart(0).
		SetPrivate("N").
		SetAETInterval(4).
		SetAETDate(0).
		SetUsePowers("Y")

	return f
}

func (f *gameFactory) Build() models.Game {
	return f.Game
}

func (f *gameFactory) BuildInsert() models.Game {
	gID, _ := GameRepo.CreateGame(f.Game)
	f.Game.ID = gID
	return f.Game
}
