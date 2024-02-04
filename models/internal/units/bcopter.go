package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type bcopter struct {
	directUnit
}

func NewBCopter(m *unit) Unit {
	u := &bcopter{
		directUnit{
			BCopter(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func BCopter() unit {
	return unit{
		Name:           unitnames.Bomber,
		MovementType:   movementtypes.A,
		MovementPoints: 6,
		Vision:         3,
		Fuel:           99,
		FuelPerTurn:    2,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           6,
		Cost:           9000,
	}
}
