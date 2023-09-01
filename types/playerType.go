package types

import (
	"time"

	"github.com/awbw/2040/models"
)

type Player struct {
	models.Player
	LastRead  time.Time
	TurnStart time.Time
}

func NewPlayer(p models.Player) Player {
	return Player{
		Player:    p,
		LastRead:  p.LastRead.Time,
		TurnStart: p.TurnStart.Time,
	}
}
