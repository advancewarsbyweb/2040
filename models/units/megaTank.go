package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type megaTank struct {
	directUnit
}

func NewMegaTank(m *unit) Unit {
	u := &megaTank{
		directUnit{
			MegaTank(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func MegaTank() unit {
	return unit{
		name:           unitnames.MegaTank,
		movementType:   movementtypes.T,
		movementPoints: 4,
		vision:         1,
		fuel:           50,
		fuelPerTurn:    0,
		shortRange:     1,
		longRange:      1,
		ammo:           3,
		cost:           28000,
	}
}
