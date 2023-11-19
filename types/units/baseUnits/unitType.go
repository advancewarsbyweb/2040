package unittypes

import (
	"github.com/awbw/2040/db"
	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types"
	unitnames "github.com/awbw/2040/types/units/names"
)

type Unit interface {
	BaseUnit
	SetGame(g *types.Game) Unit
	GetGame() *types.Game
	SetPlayer(p *types.Player) Unit
	GetPlayer() *types.Player
	GetName() unitnames.UnitName
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
	SetBaseUnit(name unitnames.UnitName) Unit
}

type unit struct {
	models.Unit
	Player *types.Player
	Game   *types.Game
	// This is embedded in to have the BaseUnit's specific methods like Move, Fire, etc
	BaseUnit
}

func NewUnit(u models.Unit) Unit {
	ut := unit{
		Unit:     u,
		Player:   nil,
		Game:     nil,
		BaseUnit: MakeUnit[u.Name](),
	}
	if u.Player != nil {
		p := types.NewPlayer(*u.Player)
		ut.Player = &p
	}
	if u.Game != nil {
		g := types.NewGame(*u.Game)
		ut.Game = &g
	}

	return &ut
}

func CreateUnit(name unitnames.UnitName) Unit {
	u := db.UnitFactory.Create(name).Build()
	return NewUnit(u)
}

func (u *unit) SetGame(g *types.Game) Unit {
	u.Game = g
	u.PlayerID = g.ID
	return u
}

func (u *unit) GetGame() *types.Game {
	return u.Game
}

func (u *unit) SetPlayer(p *types.Player) Unit {
	u.Player = p
	u.PlayerID = p.ID
	return u
}

func (u *unit) GetPlayer() *types.Player {
	return u.Player
}

func (u *unit) GetName() unitnames.UnitName {
	return u.Name
}

func (u *unit) GetAmmo() int {
	return u.Ammo
}

func (u *unit) SetAmmo(ammo int) Unit {
	if ammo > 0 {
		u.Ammo = ammo
	}
	return u
}

func (u *unit) GetFuel() int {
	return u.Fuel
}

func (u *unit) SetFuel(fuel int) Unit {
	u.Fuel = fuel
	return u
}

func (u *unit) GetHp() float64 {
	return u.HP
}

func (u *unit) SetHp(hp float64) Unit {
	u.HP = hp
	return u
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
	return u
}

func (u *unit) GetMoved() int {
	return u.Moved
}

func (u *unit) SetMoved(moved int) Unit {
	u.Moved = moved
	return u
}

func (u *unit) GetShortRange() int {
	return u.ShortRange
}

func (u *unit) GetLongRange() int {
	return u.LongRange
}

func (u *unit) SetBaseUnit(name unitnames.UnitName) Unit {
	bu := MakeUnit[name]()
	u.Name = name
	u.BaseUnit = bu
	return u
}
