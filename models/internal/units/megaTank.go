package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type megaTank struct {
	directUnit
}

func NewMegaTank(m *unit) Unit {
	u := &megaTank{
		directUnit{
			MegaTank(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func MegaTank() unit {
	return unit{
		Name:           unitnames.MegaTank,
		MovementType:   movementtypes.T,
		MovementPoints: 4,
		Vision:         1,
		Fuel:           50,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           3,
		Cost:           28000,
	}
}
