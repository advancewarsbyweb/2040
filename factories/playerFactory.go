package factories

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types"
	"gopkg.in/guregu/null.v4"
)

type PlayerFactory struct {
	Player types.Player
}

var Player PlayerFactory

func NewPlayerFactory() PlayerFactory {
	return PlayerFactory{}
}

func init() {
	Player = NewPlayerFactory()
}

func (f *PlayerFactory) Create() *PlayerFactory {
	playerId := rand.Intn(100)
	playerType := types.NewPlayer(models.Player{
		ID:           playerId,
		GameID:       rand.Intn(100),
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
	f.Player = playerType
	return f
}

func (f *PlayerFactory) CreateRelations() *PlayerFactory {

	u := User.Create().BuildInsert()
	g := Game.Create().BuildInsert()

	p := Player.Create()
	p.Player.GameID = g.ID
	p.Player.UserID = u.ID

	p.Player.User = &u
	p.Player.Game = &g
	f.Player = p.Player
	return f
}

func (f *PlayerFactory) SetGame(g *types.Game) *PlayerFactory {
	f.Player.GameID = g.ID
	f.Player.Game = g
	return f
}

func (f *PlayerFactory) SetUser(u *types.User) *PlayerFactory {
	f.Player.UserID = u.ID
	f.Player.User = u
	return f
}

func (f *PlayerFactory) Build() types.Player {
	return f.Player
}

func (f *PlayerFactory) BuildInsert() types.Player {
	pID, _ := db.PlayerRepo.CreatePlayer(f.Player)
	f.Player.ID = pID
	return f.Player
}
