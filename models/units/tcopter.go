package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type tcopter struct {
	transportUnit
}

func NewTCopter(m *unit) Unit {
	u := &tcopter{
		transportUnit{
			TCopter(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func TCopter() unit {
	return unit{
		name:         unitnames.TCopter,
		movementType: movementtypes.A,
		vision:       1,
		fuel:         99,
		fuelPerTurn:  2,
		shortRange:   0,
		longRange:    0,
		ammo:         0,
		cost:         5000,
	}
}
