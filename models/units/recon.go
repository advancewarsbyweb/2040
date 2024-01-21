package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type recon struct {
	directUnit
}

func NewRecon(m *unit) Unit {
	u := &recon{
		directUnit{
			Recon(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Recon() unit {
	return unit{
		name:           unitnames.Recon,
		movementType:   movementtypes.W,
		movementPoints: 8,
		vision:         5,
		fuel:           99,
		fuelPerTurn:    0,
		shortRange:     1,
		longRange:      1,
		ammo:           -1,
		cost:           4000,
	}
}
