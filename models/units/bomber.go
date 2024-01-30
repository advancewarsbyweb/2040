package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type bomber struct {
	directUnit
}

func NewBomber(m *unit) Unit {
	u := &bomber{
		directUnit{
			Bomber(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Bomber() unit {
	return unit{
		name:           unitnames.Bomber,
		movementType:   movementtypes.A,
		movementPoints: 8,
		vision:         2,
		fuel:           99,
		fuelPerTurn:    5,
		shortRange:     1,
		longRange:      1,
		ammo:           9,
		cost:           22000,
	}
}
