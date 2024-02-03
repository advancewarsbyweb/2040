package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type mech struct {
	directUnit
}

func (u *mech) Load() {
}

func NewMech(m *unit) Unit {
	u := &mech{
		directUnit{
			Mech(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Mech() unit {
	return unit{
		Name:           unitnames.Mech,
		MovementType:   movementtypes.B,
		MovementPoints: 2,
		Vision:         2,
		Fuel:           99,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           3,
		Cost:           3000,
	}
}
