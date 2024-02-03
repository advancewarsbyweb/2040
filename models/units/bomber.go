package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type bomber struct {
	directUnit
}

func NewBomber(m *unit) Unit {
	u := &bomber{
		directUnit{
			Bomber(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Bomber() unit {
	return unit{
		Name:           unitnames.Bomber,
		MovementType:   movementtypes.A,
		MovementPoints: 8,
		Vision:         2,
		Fuel:           99,
		FuelPerTurn:    5,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           9,
		Cost:           22000,
	}
}
