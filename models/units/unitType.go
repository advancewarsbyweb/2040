package models

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/types/units/names"
)

// We return the IUnit in the Unit Model to be able to chain methods
type Unit interface {
	SetGame(g *models.Game) Unit
	Game() *models.Game
	SetPlayer(p *models.Player) Unit
	Player() *models.Player
	Name() unitnames.UnitName
	Ammo() int
	SetAmmo(ammo int) Unit
	Fuel() int
	SetFuel(fuel int) Unit
	Hp() float64
	SetHp(hp float64) Unit
	X() int
	Y() int
	SetPos(x int, y int) Unit
	Moved() int
	SetMoved(moved int) Unit
	ShortRange() int
	LongRange() int

	Move()
	Fire(a Unit, d Unit) error
	Load()
}

func (u *unit) Move() {
}

func (u *unit) Fire(a Unit, b Unit) error {
	return nil
}

func (u *unit) Load() {
}

func (u *unit) SetGame(g *models.Game) Unit {
	u.game = g
	u.gameID = g.ID
	return u.IUnit
}

func (u *unit) Game() *models.Game {
	return u.game
}

func (u *unit) SetPlayer(p *models.Player) Unit {
	u.player = p
	u.playerID = p.ID
	return u.IUnit
}

func (u *unit) Player() *models.Player {
	return u.player
}

func (u *unit) Name() unitnames.UnitName {
	return u.name
}

func (u *unit) Ammo() int {
	return u.ammo
}

func (u *unit) SetAmmo(ammo int) Unit {
	if ammo >= 0 {
		u.ammo = ammo
	}
	return u.IUnit
}

func (u *unit) Fuel() int {
	return u.fuel
}

func (u *unit) SetFuel(fuel int) Unit {
	u.fuel = fuel
	return u.IUnit
}

func (u *unit) Hp() float64 {
	return u.hp
}

func (u *unit) SetHp(hp float64) Unit {
	u.hp = hp
	return u.IUnit
}

func (u *unit) X() int {
	return u.x
}

func (u *unit) Y() int {
	return u.y
}

func (u *unit) SetPos(x int, y int) Unit {
	u.x = x
	u.y = y
	return u.IUnit
}

func (u *unit) Moved() int {
	return u.moved
}

func (u *unit) SetMoved(moved int) Unit {
	u.moved = moved
	return u.IUnit
}

func (u *unit) ShortRange() int {
	return u.shortRange
}

func (u *unit) LongRange() int {
	return u.longRange
}
