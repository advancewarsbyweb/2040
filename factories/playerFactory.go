package factories

import (
	"database/sql"
	"math/rand"
	"strconv"
	"time"

	"github.com/awbw/2040/models"
)

type PlayerFactory struct {
	models.Player
}

var Player PlayerFactory

func NewPlayerFactory() PlayerFactory {
	return PlayerFactory{}
}

func init() {
	Player = NewPlayerFactory()
}

func (f *PlayerFactory) Create() models.Player {
	playerId := rand.Int()
	return models.Player{
		ID:           playerId,
		GameID:       Game.Create().ID,
		UserID:       1,
		CountryID:    1,
		CoID:         1,
		Funds:        rand.Intn(100000),
		OldInterface: "N",
		Eliminated:   "N",
		LastRead:     sql.NullTime{Time: time.Now()},
		TurnStart:    sql.NullTime{Time: time.Now()},
		TurnClock:    10800,
		PauseTimer:   false,
		CoPower:      9000,
		CoPowerOn:    "N",
		Order:        1,
		AcceptDraw:   "N",
		CoMaxPower:   27000,
		CoMaxSpower:  54000,
		CoImage:      "IMAGE",
		Team:         strconv.Itoa(playerId),
		AETCount:     0,
		TurnFlag:     true,
	}
}
