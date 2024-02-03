package models

import (
	"github.com/awbw/2040/models"
	terrainmodels "github.com/awbw/2040/models/terrains"

	unitnames "github.com/awbw/2040/models/units/names"
)

// We return the IUnit in the Unit Model to be able to chain methods
type Unit interface {
	SetGame(g *models.Game) Unit
	GetGame() *models.Game
	SetPlayer(p *models.Player) Unit
	GetPlayer() *models.Player
	GetName() unitnames.UnitName
	GetCost() int
	GetAmmo() int
	SetAmmo(ammo int) Unit
	GetFuel() int
	SetFuel(fuel int) Unit
	GetHp() float64
	SetHp(hp float64) Unit
	GetX() int
	GetY() int
	SetPos(x int, y int) Unit
	GetMoved() int
	SetMoved(moved int) Unit
	GetShortRange() int
	GetLongRange() int
	GetTile() *terrainmodels.Tile
	SetTile(t *terrainmodels.Tile) Unit
	GetUnit() *unit

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
	u.Game = g
	u.GameID = g.ID
	return u.IUnit
}

func (u *unit) GetGame() *models.Game {
	return u.Game
}

func (u *unit) SetPlayer(p *models.Player) Unit {
	u.Player = p
	u.PlayerID = p.ID
	return u.IUnit
}

func (u *unit) GetPlayer() *models.Player {
	return u.Player
}

func (u *unit) GetName() unitnames.UnitName {
	return u.Name
}

func (u *unit) GetCost() int {
	return u.Cost
}

func (u *unit) GetAmmo() int {
	return u.Ammo
}

func (u *unit) SetAmmo(ammo int) Unit {
	if ammo >= 0 {
		u.Ammo = ammo
	}
	return u.IUnit
}

func (u *unit) GetFuel() int {
	return u.Fuel
}

func (u *unit) SetFuel(fuel int) Unit {
	u.Fuel = fuel
	return u.IUnit
}

func (u *unit) GetHp() float64 {
	return u.Hp
}

func (u *unit) SetHp(hp float64) Unit {
	u.Hp = hp
	return u.IUnit
}

func (u *unit) GetX() int {
	return u.X
}

func (u *unit) GetY() int {
	return u.Y
}

func (u *unit) SetPos(x int, y int) Unit {
	u.X = x
	u.Y = y
	return u.IUnit
}

func (u *unit) GetMoved() int {
	return u.Moved
}

func (u *unit) SetMoved(moved int) Unit {
	u.Moved = moved
	return u.IUnit
}

func (u *unit) GetShortRange() int {
	return u.ShortRange
}

func (u *unit) GetLongRange() int {
	return u.LongRange
}

func (u *unit) GetTile() *terrainmodels.Tile {
	return u.Tile
}

func (u *unit) SetTile(t *terrainmodels.Tile) Unit {
	u.Tile = t
	return u.IUnit
}

func (u *unit) GetUnit() *unit {
	return u
}
