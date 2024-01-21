package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type mdTank struct {
	directUnit
}

func NewMDTank(m *unit) Unit {
	u := &mdTank{
		directUnit{
			MDTank(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func MDTank() unit {
	return unit{
		name:           unitnames.MDTank,
		movementType:   movementtypes.T,
		movementPoints: 5,
		vision:         1,
		fuel:           50,
		fuelPerTurn:    0,
		shortRange:     1,
		longRange:      1,
		ammo:           8,
		cost:           16000,
	}
}
