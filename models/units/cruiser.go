package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type cruiser struct {
	directUnit
}

func NewCruiser(m *unit) Unit {
	u := &cruiser{
		directUnit{
			Cruiser(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Cruiser() unit {
	return unit{
		Name:           unitnames.Cruiser,
		MovementType:   movementtypes.S,
		MovementPoints: 6,
		Vision:         3,
		Fuel:           99,
		FuelPerTurn:    1,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           9,
		Cost:           18000,
	}
}
