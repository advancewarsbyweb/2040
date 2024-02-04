package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type artillery struct {
	indirectUnit
}

func NewArtillery(m *unit) Unit {
	u := &artillery{
		indirectUnit{
			Artillery(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Artillery() unit {
	return unit{
		Name:           unitnames.Artillery,
		MovementType:   movementtypes.T,
		MovementPoints: 5,
		Vision:         1,
		Fuel:           99,
		FuelPerTurn:    0,
		ShortRange:     2,
		LongRange:      3,
		Ammo:           6,
		Cost:           6000,
	}
}
