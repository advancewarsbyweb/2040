package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
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
		name:           unitnames.Carrier,
		movementType:   movementtypes.S,
		movementPoints: 5,
		vision:         4,
		fuel:           99,
		fuelPerTurn:    1,
		shortRange:     3,
		longRange:      8,
		ammo:           9,
		cost:           30000,
	}
}
