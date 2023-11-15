package types

import (
	"github.com/awbw/2040/models"
)

type Unit struct {
	models.Unit
	Player *Player
	Game   *Game
	models.BaseUnit
}

func NewUnit(u models.Unit) Unit {
	ut := Unit{
		Unit:   u,
		Player: nil,
		Game:   nil,
	}
	if u.Player != nil {
		p := NewPlayer(*u.Player)
		ut.Player = &p
	}
	if u.Game != nil {
		g := NewGame(*u.Game)
		ut.Game = &g
	}

	return ut
}
