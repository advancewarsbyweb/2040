package baseunits

import (
	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types"
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type Unit interface {
	Move()
	Fire(def Unit) error
	Load()

	// Getters / Setters
	GetName() unitnames.UnitName
	GetAmmo() int
	SetAmmo(ammo int)
	GetHp() float64
	SetHp(hp float64)
	GetX() int
	GetY() int
	SetPos(x int, y int)
	GetMoved() int
	SetMoved(moved int)
}

type unit types.Unit

func (u *unit) Move() {

}

func (u *unit) Load() {
}

func (u *unit) GetName() unitnames.UnitName {
	return u.Name
}

func (u *unit) GetAmmo() int {
	return u.Ammo
}

func (u *unit) SetAmmo(ammo int) {
	u.Ammo = ammo
}

func (u *unit) GetHp() float64 {
	return u.HP
}

func (u *unit) SetHp(hp float64) {
	u.HP = hp
}

func (u *unit) GetX() int {
	return u.X
}

func (u *unit) GetY() int {
	return u.Y
}

func (u *unit) SetPos(x int, y int) {
	u.X = x
	u.Y = y
}

func (u *unit) GetMoved() int {
	return u.Moved
}

func (u *unit) SetMoved(moved int) {
	u.Moved = moved
}

func Recon() models.BaseUnit {
	return models.BaseUnit{
		Name:           unitnames.Recon,
		MovementType:   movementtypes.W,
		MovementPoints: 8,
		Vision:         5,
		Fuel:           99,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           -1,
		Cost:           4000,
	}
}

func TCopter() models.BaseUnit {
	return models.BaseUnit{
		Name:         unitnames.TCopter,
		MovementType: movementtypes.A,
		Vision:       1,
		Fuel:         99,
		FuelPerTurn:  2,
		ShortRange:   0,
		LongRange:    0,
		Ammo:         0,
		Cost:         5000,
	}
}

func APC() models.BaseUnit {
	return models.BaseUnit{
		Name:         unitnames.APC,
		MovementType: movementtypes.T,
		Vision:       1,
		Fuel:         99,
		FuelPerTurn:  0,
		ShortRange:   0,
		LongRange:    0,
		Ammo:         0,
		Cost:         5000,
	}
}

func Artillery() models.BaseUnit {
	return models.BaseUnit{
		Name:           unitnames.Artillery,
		MovementType:   movementtypes.T,
		MovementPoints: 5,
		Vision:         1,
		Fuel:           99,
		FuelPerTurn:    0,
		ShortRange:     2,
		LongRange:      3,
		Ammo:           6,
		Cost:           6000,
	}
}
