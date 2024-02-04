package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type rocket struct {
	indirectUnit
}

func NewRocket(m *unit) Unit {
	u := &rocket{
		indirectUnit{
			Rocket(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Rocket() unit {
	return unit{
		Name:           unitnames.Rocket,
		MovementType:   movementtypes.W,
		MovementPoints: 5,
		Vision:         1,
		Fuel:           50,
		FuelPerTurn:    0,
		ShortRange:     3,
		LongRange:      5,
		Ammo:           6,
		Cost:           15000,
	}
}
