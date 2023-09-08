package models

import "gopkg.in/guregu/null.v4"

type Player struct {
	ID           int         `db:"players_id"`
	GameID       int         `db:"players_games_id"`
	UserID       int         `db:"players_users_id"`
	CountryID    int         `db:"players_countries_id"`
	CoID         int         `db:"players_co_id"`
	Funds        int         `db:"players_funds"`
	Turn         null.Int    `db:"players_turn"`
	OldInterface string      `db:"players_old_interface"`
	UniqID       null.String `db:"players_uniq_id"`
	Eliminated   string      `db:"players_eliminated"`
	LastRead     null.Time   `db:"players_last_read"`
	TurnStart    null.Time   `db:"players_turn_start"`
	TurnClock    int         `db:"players_turn_clock"`
	PauseTimer   bool        `db:"players_pause_timer"`
	CoPower      int         `db:"players_co_power"`
	CoPowerOn    string      `db:"players_co_power_on"`
	Order        int         `db:"players_order"`
	AcceptDraw   string      `db:"players_accept_draw"`
	CoMaxPower   int         `db:"players_co_max_power"`
	CoMaxSpower  int         `db:"players_co_max_spower"`
	CoImage      string      `db:"players_co_image"`
	Team         string      `db:"players_team"`
	AETCount     int         `db:"players_aet_count"`
	TurnFlag     bool        `db:"players_turn_flag"`
	User         *User       `db:""`
	Game         *Game       `db:""`
}
