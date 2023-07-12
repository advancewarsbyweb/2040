package types

import (
	"time"

	"github.com/awbw/2040/models"
)

type Game struct {
	models.Game
	StartDate    time.Time
	EndDate      time.Time
	ActivityDate time.Time
	WeatherStart int
	AETInterval  int
	AETDate      int
}

func NewGame(g models.Game) Game {
	return Game{
		Game:         g,
		StartDate:    g.StartDate.Time,
		EndDate:      g.EndDate.Time,
		ActivityDate: g.ActivityDate.Time,
		WeatherStart: int(g.WeatherStart.Int64),
		AETInterval:  int(g.AETInterval.Int64),
		AETDate:      int(g.AETDate.Int64),
	}
}
