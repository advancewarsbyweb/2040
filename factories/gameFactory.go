package factories

import (
	"time"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types"
	"github.com/bxcodec/faker/v4"
	"gopkg.in/guregu/null.v4"
)

type GameFactory struct {
	Game types.Game
}

var Game GameFactory

func NewGameFactory() GameFactory {
	return GameFactory{}
}

func init() {
	Game = NewGameFactory()
}

func (f *GameFactory) Create() *GameFactory {
	f.Game = types.NewGame(models.Game{
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
	})
	return f
}

func (f *GameFactory) Build() types.Game {
	return f.Game
}

func (f *GameFactory) BuildInsert() types.Game {
	gID, _ := db.GameRepo.CreateGame(f.Game)
	f.Game.ID = gID
	return f.Game
}
