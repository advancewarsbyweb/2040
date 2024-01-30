package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type cruiser struct {
	directUnit
}

func NewCruiser(m *unit) Unit {
	u := &cruiser{
		directUnit{
			Cruiser(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Cruiser() unit {
	return unit{
		name:           unitnames.Cruiser,
		movementType:   movementtypes.S,
		movementPoints: 6,
		vision:         3,
		fuel:           99,
		fuelPerTurn:    1,
		shortRange:     1,
		longRange:      1,
		ammo:           9,
		cost:           18000,
	}
}
