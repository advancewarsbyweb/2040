package types

import (
	"time"

	"github.com/awbw/2040/models"
)

type Player struct {
	models.Player
	LastRead  time.Time
	TurnStart time.Time
	User      *User
	Game      *Game
}

func NewPlayer(p models.Player) Player {
	playerType := Player{
		Player:    p,
		LastRead:  p.LastRead.Time,
		TurnStart: p.TurnStart.Time,
		User:      nil,
		Game:      nil,
	}
	if p.User != nil {
		u := NewUser(*p.User)
		playerType.User = &u
	}
	if p.Game != nil {
		g := NewGame(*p.Game)
		playerType.Game = &g
	}
	return playerType
}
