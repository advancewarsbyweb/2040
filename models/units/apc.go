package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type apc struct {
	transportUnit
}

func NewAPC(m *unit) Unit {
	u := &apc{
		transportUnit{
			APC(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func APC() unit {
	return unit{
		name:         unitnames.APC,
		movementType: movementtypes.T,
		vision:       1,
		fuel:         99,
		fuelPerTurn:  0,
		shortRange:   0,
		longRange:    0,
		ammo:         0,
		cost:         5000,
	}
}
