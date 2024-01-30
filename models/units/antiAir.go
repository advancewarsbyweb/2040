package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type antiAir struct {
	directUnit
}

func NewAntiAir(m *unit) Unit {
	u := &antiAir{
		directUnit{
			AntiAir(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func AntiAir() unit {
	return unit{
		name:           unitnames.AntiAir,
		movementType:   movementtypes.T,
		movementPoints: 6,
		vision:         2,
		fuel:           60,
		fuelPerTurn:    0,
		shortRange:     1,
		longRange:      1,
		ammo:           9,
		cost:           8000,
	}
}
