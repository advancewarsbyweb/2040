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
		Name:         unitnames.TCopter,
		MovementType: movementtypes.A,
		Vision:       1,
		Fuel:         99,
		FuelPerTurn:  2,
		ShortRange:   0,
		LongRange:    0,
		Ammo:         0,
		Cost:         5000,
	}
}
