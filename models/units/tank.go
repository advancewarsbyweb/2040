package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type tank struct {
	directUnit
}

func NewTank(m *unit) Unit {
	u := &tank{
		directUnit{
			Tank(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Tank() unit {
	return unit{
		name:           unitnames.Tank,
		movementType:   movementtypes.T,
		movementPoints: 6,
		vision:         3,
		fuel:           70,
		fuelPerTurn:    0,
		shortRange:     1,
		longRange:      1,
		ammo:           9,
		cost:           7000,
	}
}
