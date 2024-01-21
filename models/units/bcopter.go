package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
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
		name:           unitnames.Bomber,
		movementType:   movementtypes.A,
		movementPoints: 6,
		vision:         3,
		fuel:           99,
		fuelPerTurn:    2,
		shortRange:     1,
		longRange:      1,
		ammo:           6,
		cost:           9000,
	}
}
