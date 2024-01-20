package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type tank struct {
	directUnit
}

func NewTank(m *unit) Unit {
	u := &tank{
		directUnit{
			unit{
				baseUnit: Tank(),
			},
		},
	}
	u.SetUnitProperties(m)
	return u
}

func Tank() baseUnit {
	return baseUnit{
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
