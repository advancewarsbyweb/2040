package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type sub struct {
	directUnit
}

func NewSub(m *unit) Unit {
	u := &sub{
		directUnit{
			Sub(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Sub() unit {
	return unit{
		name:           unitnames.Sub,
		movementType:   movementtypes.S,
		movementPoints: 5,
		vision:         5,
		fuel:           60,
		fuelPerTurn:    1,
		shortRange:     1,
		longRange:      1,
		ammo:           6,
		cost:           20000,
	}
}
