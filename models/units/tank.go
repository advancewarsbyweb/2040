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
		Name:           unitnames.Tank,
		MovementType:   movementtypes.T,
		MovementPoints: 6,
		Vision:         3,
		Fuel:           70,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           9,
		Cost:           7000,
	}
}
