package db

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/awbw/2040/models"
	"gopkg.in/guregu/null.v4"
)

type playerFactory struct {
	Player *models.Player
}

var PlayerFactory playerFactory

func NewPlayerFactory() playerFactory {
	return playerFactory{}
}

func init() {
	PlayerFactory = NewPlayerFactory()
}

func (f *playerFactory) Create() *playerFactory {
	playerId := rand.Intn(100)
	playerType := models.Player{
		ID:           1,
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
	}
	f.Player = &playerType
	return f
}

func (f *playerFactory) CreateRelations() *playerFactory {

	u := UserFactory.Create().BuildInsert()
	g := GameFactory.Create().BuildInsert()

	p := PlayerFactory.Create().SetGame(g).SetUser(&u)
	f.Player = p.Player
	return f
}

func (f *playerFactory) CreateUser() *playerFactory {
	u := UserFactory.Create().BuildInsert()
	f.Player.UserID = u.ID
	f.Player.User = &u
	return f
}

func (f *playerFactory) SetGame(g *models.Game) *playerFactory {
	f.Player.GameID = g.ID
	f.Player.Game = g
	return f
}

func (f *playerFactory) SetUser(u *models.User) *playerFactory {
	f.Player.UserID = u.ID
	f.Player.User = u
	return f
}

func (f *playerFactory) Build() *models.Player {
	return f.Player
}

func (f *playerFactory) BuildInsert() *models.Player {
	pID, _ := PlayerRepo.CreatePlayer(*f.Player)
	f.Player.ID = pID
	return f.Player
}
