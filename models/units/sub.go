package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type sub struct {
	directUnit
}

func NewSub(m *unit) Unit {
	u := &sub{
		directUnit{
			Sub(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Sub() unit {
	return unit{
		Name:           unitnames.Sub,
		MovementType:   movementtypes.S,
		MovementPoints: 5,
		Vision:         5,
		Fuel:           60,
		FuelPerTurn:    1,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           6,
		Cost:           20000,
	}
}
