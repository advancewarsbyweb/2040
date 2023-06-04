package factories

import (
	"database/sql"
	"math/rand"
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
		ID:           rand.Int(),
		Name:         faker.Word(),
		StartDate:    sql.NullTime{Time: time.Now()},
		EndDate:      sql.NullTime{},
		ActivityDate: sql.NullTime{Time: time.Now()},
		MapID:        1,
		WeatherType:  "Clear",
		WeatherCode:  "C",
		WeatherStart: 0,
		Private:      false,
		TurnID:       0,
		Day:          1,
		Active:       true,
		Funds:        1000,
		CaptureLimit: 1000,
		DayLimit:     0,
		Fog:          "N",
		Comment:      faker.Paragraph(),
		Type:         "N",
		Official:     "N",
		MinRating:    0,
		UnitLimit:    50,
		League:       "N",
		Team:         "N",
		CreatedAt:    sql.NullTime{Time: time.Now()},
		UpdatedAt:    sql.NullTime{Time: time.Now()},
		DeletedAt:    sql.NullTime{},
	}
}
