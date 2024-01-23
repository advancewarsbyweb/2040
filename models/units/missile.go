package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type missile struct {
	indirectUnit
}

func NewMissile(m *unit) Unit {
	u := &missile{
		indirectUnit{
			Missile(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Missile() unit {
	return unit{
		name:           unitnames.Missile,
		movementType:   movementtypes.W,
		movementPoints: 4,
		vision:         5,
		fuel:           50,
		fuelPerTurn:    0,
		shortRange:     3,
		longRange:      5,
		ammo:           6,
		cost:           12000,
	}
}
