package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type apc struct {
	transportUnit
}

func NewAPC(m *unit) Unit {
	u := &apc{
		transportUnit{
			APC(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func APC() unit {
	return unit{
		Name:         unitnames.APC,
		MovementType: movementtypes.T,
		Vision:       1,
		Fuel:         99,
		FuelPerTurn:  0,
		ShortRange:   0,
		LongRange:    0,
		Ammo:         0,
		Cost:         5000,
	}
}
