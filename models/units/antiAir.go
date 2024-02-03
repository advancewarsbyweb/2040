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
		Name:           unitnames.AntiAir,
		MovementType:   movementtypes.T,
		MovementPoints: 6,
		Vision:         2,
		Fuel:           60,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           9,
		Cost:           8000,
	}
}
