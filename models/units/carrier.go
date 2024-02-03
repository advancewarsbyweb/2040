package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type carrier struct {
	indirectUnit
}

func NewCarrier(m *unit) Unit {
	u := &carrier{
		indirectUnit{
			Carrier(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Carrier() unit {
	return unit{
		Name:           unitnames.Carrier,
		MovementType:   movementtypes.S,
		MovementPoints: 5,
		Vision:         4,
		Fuel:           99,
		FuelPerTurn:    1,
		ShortRange:     3,
		LongRange:      8,
		Ammo:           9,
		Cost:           30000,
	}
}
