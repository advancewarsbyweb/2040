package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Game struct {
	ID            int         `db:"games_id" json:"id"`
	Name          string      `db:"games_name" validate:"required"`
	Password      string      `db:"games_password" json:"-"`
	StartDate     null.Time   `db:"games_start_date"`
	EndDate       null.Time   `db:"games_end_date"`
	CreatorID     int         `db:"games_creator"`
	ActivityDate  null.Time   `db:"games_activity_date"`
	MapID         int         `db:"games_maps_id" validate:"required"`
	WeatherType   string      `db:"games_weather_type"`
	WeatherStart  null.Int    `db:"games_weather_start"`
	WeatherCode   string      `db:"games_weather_code" validate:"eq=N|eq=R|eq=S"`
	Private       null.String `db:"games_private"`
	TurnID        int         `db:"games_turn"`
	Day           int         `db:"games_day"`
	Active        string      `db:"games_active"`
	Funds         int         `db:"games_funds" validate:"gte=0"`
	CaptureLimit  int         `db:"games_capture_win" validate:"gte=0"`
	DayLimit      int         `db:"games_days" validate:"gte=0"`
	Fog           string      `db:"games_fog"`
	Comment       string      `db:"games_comment"`
	Type          string      `db:"games_type"`
	BootInterval  int         `db:"games_boot_interval"`
	StartingFunds int         `db:"games_starting_funds"`
	Official      string      `db:"games_official"`
	MinRating     int         `db:"games_min_rating"`
	UnitLimit     int         `db:"games_unit_limit"`
	League        string      `db:"games_league"`
	Team          string      `db:"games_team"`
	AetInterval   null.Int    `db:"games_aet_interval"`
	AetDate       null.Int    `db:"games_aet_date"`
	UsePowers     null.String `db:"games_use_powers"`
}

// Getters

func (g *Game) GetActivityDate() time.Time { return g.ActivityDate.Time }

func (g *Game) GetStartDate() time.Time { return g.StartDate.Time }

func (g *Game) GetEndDate() time.Time { return g.EndDate.Time }

func (g *Game) GetWeatherStart() int { return int(g.WeatherStart.Int64) }

func (g *Game) GetPrivate() string { return g.Private.String }

func (g *Game) GetAETInterval() int { return int(g.AetInterval.Int64) }

func (g *Game) GetAETDate() int { return int(g.AetDate.Int64) }

func (g *Game) GetUsePowers() string { return g.UsePowers.String }

// Setters - return the model to be able to chain methods

func (g *Game) SetActivityDate(t time.Time) *Game {
	g.ActivityDate = null.TimeFrom(t)
	return g
}

func (g *Game) SetStartDate(t time.Time) *Game {
	g.StartDate = null.TimeFrom(t)
	return g
}

func (g *Game) SetEndDate(t time.Time) *Game {
	g.EndDate = null.TimeFrom(t)
	return g
}

func (g *Game) SetWeatherStart(i int64) *Game {
	g.WeatherStart = null.IntFrom(i)
	return g
}

func (g *Game) SetPrivate(s string) *Game {
	g.Private = null.StringFrom(s)
	return g
}

func (g *Game) SetAETInterval(i int64) *Game {
	g.AetInterval = null.IntFrom(i)
	return g
}

func (g *Game) SetAETDate(i int64) *Game {
	g.AetDate = null.IntFrom(i)
	return g
}

func (g *Game) SetUsePowers(s string) *Game {
	g.UsePowers = null.StringFrom(s)
	return g
}
