package models

import (
	"database/sql"
)

type Game struct {
	ID            int            `db:"games_id"`
	Name          string         `db:"games_name"`
	Password      string         `db:"games_password"`
	StartDate     sql.NullTime   `db:"games_start_date"`
	EndDate       sql.NullTime   `db:"games_end_date"`
	CreatorID     int            `db:"games_creator"`
	ActivityDate  sql.NullTime   `db:"games_activity_date"`
	MapID         int            `db:"games_maps_id"`
	WeatherType   string         `db:"games_weather_type"`
	WeatherStart  sql.NullInt16  `db:"games_weather_start"`
	WeatherCode   string         `db:"games_weather_code"`
	Private       sql.NullString `db:"games_private"`
	TurnID        int            `db:"games_turn"`
	Day           int            `db:"games_day"`
	Active        string         `db:"games_active"`
	Funds         int            `db:"games_funds"`
	CaptureLimit  int            `db:"games_capture_win"`
	DayLimit      int            `db:"games_days"`
	Fog           string         `db:"games_fog"`
	Comment       string         `db:"games_comment"`
	Type          string         `db:"games_type"`
	BootInterval  int            `db:"games_boot_interval"`
	StartingFunds int            `db:"games_starting_funds"`
	Official      string         `db:"games_official"`
	MinRating     int            `db:"games_min_rating"`
	UnitLimit     int            `db:"games_unit_limit"`
	League        string         `db:"games_league"`
	Team          string         `db:"games_team"`
	AETInterval   sql.NullInt16  `db:"games_aet_interval"`
	AETDate       sql.NullTime   `db:"games_aet_date"`
	UsePowers     sql.NullString `db:"games_use_powers"`
}
