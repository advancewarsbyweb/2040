package models

import (
	unitnames "github.com/awbw/2040/types/units/names"
)

// We return the IUnit in the Unit Model to be able to chain methods
type IUnit interface {
	SetGame(g *Game) IUnit
	GetGame() *Game
	SetPlayer(p *Player) IUnit
	GetPlayer() *Player
	GetName() unitnames.UnitName
	GetCost() int
	GetAmmo() int
	SetAmmo(ammo int) IUnit
	GetFuel() int
	SetFuel(fuel int) IUnit
	GetHp() float64
	SetHp(hp float64) IUnit
	GetX() int
	GetY() int
	SetPos(x int, y int) IUnit
	GetMoved() int
	SetMoved(moved int) IUnit
	GetShortRange() int
	GetLongRange() int
	GetTile() *Tile
	SetTile(t *Tile) IUnit
	// Return the Unit in order to be able to return the json object to the client
	GetUnit() *Unit

	Move()
	Fire(a IUnit, d IUnit) error
	Load()
}

func (u *Unit) Move() {
}

func (u *Unit) Fire(a IUnit, b IUnit) error {
	return nil
}

func (u *Unit) Load() {
}

func (u *Unit) SetGame(g *Game) IUnit {
	u.Game = g
	u.GameID = g.ID
	return u.IUnit
}

func (u *Unit) GetGame() *Game {
	return u.Game
}

func (u *Unit) SetPlayer(p *Player) IUnit {
	u.Player = p
	u.PlayerID = p.ID
	return u.IUnit
}

func (u *Unit) GetPlayer() *Player {
	return u.Player
}

func (u *Unit) GetName() unitnames.UnitName {
	return u.Name
}

func (u *Unit) GetCost() int {
	return u.Cost
}

func (u *Unit) GetAmmo() int {
	return u.Ammo
}

func (u *Unit) SetAmmo(ammo int) IUnit {
	if ammo >= 0 {
		u.Ammo = ammo
	}
	return u.IUnit
}

func (u *Unit) GetFuel() int {
	return u.Fuel
}

func (u *Unit) SetFuel(fuel int) IUnit {
	u.Fuel = fuel
	return u.IUnit
}

func (u *Unit) GetHp() float64 {
	return u.Hp
}

func (u *Unit) SetHp(hp float64) IUnit {
	u.Hp = hp
	return u.IUnit
}

func (u *Unit) GetX() int {
	return u.X
}

func (u *Unit) GetY() int {
	return u.Y
}

func (u *Unit) SetPos(x int, y int) IUnit {
	u.X = x
	u.Y = y
	return u.IUnit
}

func (u *Unit) GetMoved() int {
	return u.Moved
}

func (u *Unit) SetMoved(moved int) IUnit {
	u.Moved = moved
	return u.IUnit
}

func (u *Unit) GetShortRange() int {
	return u.ShortRange
}

func (u *Unit) GetLongRange() int {
	return u.LongRange
}

func (u *Unit) GetTile() *Tile {
	return u.Tile
}

func (u *Unit) SetTile(t *Tile) IUnit {
	u.Tile = t
	return u.IUnit
}

func (u *Unit) GetUnit() *Unit {
	return u
}
