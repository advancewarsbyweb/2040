package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
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
