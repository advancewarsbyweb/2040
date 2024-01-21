package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type fighter struct {
	directUnit
}

func NewFighter(m *unit) Unit {
	u := &fighter{
		directUnit{
			Fighter(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Fighter() unit {
	return unit{
		name:           unitnames.Fighter,
		movementType:   movementtypes.A,
		movementPoints: 9,
		vision:         2,
		fuel:           99,
		fuelPerTurn:    5,
		shortRange:     1,
		longRange:      1,
		ammo:           9,
		cost:           20000,
	}
}
