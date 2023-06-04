package models

import (
	"database/sql"
)

type Player struct {
	ID           int
	GameID       int
	UserID       int
	CountryID    int
	CoID         int
	Funds        int
	OldInterface string
	Eliminated   string
	LastRead     sql.NullTime
	TurnStart    sql.NullTime
	TurnClock    int
	PauseTimer   bool
	CoPower      int
	CoPowerOn    string
	Order        int
	AcceptDraw   string
	CoMaxPower   int
	CoMaxSpower  int
	CoImage      string
	Team         string //events.Team
	AETCount     int
	TurnFlag     bool
}
