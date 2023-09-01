package factories

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types"
	"gopkg.in/guregu/null.v4"
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

func (f *PlayerFactory) Create() types.Player {
	playerId := rand.Intn(100)
	return types.NewPlayer(models.Player{
		ID:           playerId,
		GameID:       Game.Create().ID,
		UserID:       1,
		CountryID:    1,
		CoID:         1,
		Funds:        rand.Intn(10) * 1000,
		OldInterface: "N",
		UniqID:       null.StringFrom("N"),
		Eliminated:   "N",
		LastRead:     null.TimeFrom(time.Now()),
		TurnStart:    null.TimeFrom(time.Now()),
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
	})
}
