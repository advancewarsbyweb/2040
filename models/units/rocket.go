package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
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
		name:           unitnames.Rocket,
		movementType:   movementtypes.W,
		movementPoints: 5,
		vision:         1,
		fuel:           50,
		fuelPerTurn:    0,
		shortRange:     3,
		longRange:      5,
		ammo:           6,
		cost:           15000,
	}
}
