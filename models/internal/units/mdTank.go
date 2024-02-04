package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
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
		Name:           unitnames.MDTank,
		MovementType:   movementtypes.T,
		MovementPoints: 5,
		Vision:         1,
		Fuel:           50,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           8,
		Cost:           16000,
	}
}
