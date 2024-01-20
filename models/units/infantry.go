package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type infantry struct {
	directUnit
}

func (u *infantry) Load() {
}

// Create an Infantry from a Unit model
func NewInfantry(m *unit) Unit {
	u := &infantry{
		directUnit{
			Infantry(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Infantry() unit {
	return unit{
		name:           unitnames.Infantry,
		movementType:   movementtypes.F,
		movementPoints: 3,
		vision:         2,
		fuel:           99,
		fuelPerTurn:    0,
		shortRange:     1,
		longRange:      1,
		ammo:           -1,
		cost:           1000,
	}
}
