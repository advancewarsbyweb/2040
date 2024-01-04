package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Game struct {
	ID            int         `db:"games_id"`
	Name          string      `db:"games_name" validate:"required"`
	Password      string      `db:"games_password" json:"-"`
	startDate     null.Time   `db:"games_start_date"`
	endDate       null.Time   `db:"games_end_date"`
	CreatorID     int         `db:"games_creator"`
	activityDate  null.Time   `db:"games_activity_date"`
	MapID         int         `db:"games_maps_id" validate:"required"`
	WeatherType   string      `db:"games_weather_type"`
	weatherStart  null.Int    `db:"games_weather_start"`
	WeatherCode   string      `db:"games_weather_code" validate:"eq=N|R|S"`
	private       null.String `db:"games_private"`
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
	aetInterval   null.Int    `db:"games_aet_interval"`
	aetDate       null.Int    `db:"games_aet_date"`
	usePowers     null.String `db:"games_use_powers"`
}

// Getters

func (g *Game) ActivityDate() time.Time { return g.activityDate.Time }

func (g *Game) StartDate() time.Time { return g.startDate.Time }

func (g *Game) EndDate() time.Time { return g.endDate.Time }

func (g *Game) WeatherStart() int { return int(g.weatherStart.Int64) }

func (g *Game) Private() string { return g.private.String }

func (g *Game) AETInterval() int { return int(g.aetInterval.Int64) }

func (g *Game) AETDate() int { return int(g.aetDate.Int64) }

func (g *Game) UsePowers() string { return g.usePowers.String }

// Setters - return the model to be able to chain methods

func (g *Game) SetActivityDate(t time.Time) *Game {
	g.activityDate = null.TimeFrom(t)
	return g
}

func (g *Game) SetStartDate(t time.Time) *Game {
	g.startDate = null.TimeFrom(t)
	return g
}

func (g *Game) SetEndDate(t time.Time) *Game {
	g.endDate = null.TimeFrom(t)
	return g
}

func (g *Game) SetWeatherStart(i int64) *Game {
	g.weatherStart = null.IntFrom(i)
	return g
}

func (g *Game) SetPrivate(s string) *Game {
	g.private = null.StringFrom(s)
	return g
}

func (g *Game) SetAETInterval(i int64) *Game {
	g.aetInterval = null.IntFrom(i)
	return g
}

func (g *Game) SetAETDate(i int64) *Game {
	g.aetDate = null.IntFrom(i)
	return g
}

func (g *Game) SetUsePowers(s string) *Game {
	g.usePowers = null.StringFrom(s)
	return g
}
