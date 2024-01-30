package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type neotank struct {
	directUnit
}

func NewNeotank(m *unit) Unit {
	u := &neotank{
		directUnit{
			Neotank(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Neotank() unit {
	return unit{
		name:           unitnames.Neotank,
		movementType:   movementtypes.T,
		movementPoints: 6,
		vision:         1,
		fuel:           99,
		fuelPerTurn:    0,
		shortRange:     1,
		longRange:      1,
		ammo:           9,
		cost:           22000,
	}
}
