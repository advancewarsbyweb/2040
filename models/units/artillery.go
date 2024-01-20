package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
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
		name:           unitnames.Artillery,
		movementType:   movementtypes.T,
		movementPoints: 5,
		vision:         1,
		fuel:           99,
		fuelPerTurn:    0,
		shortRange:     2,
		longRange:      3,
		ammo:           6,
		cost:           6000,
	}
}
