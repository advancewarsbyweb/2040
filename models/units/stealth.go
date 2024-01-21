package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type stealth struct {
	directUnit
}

func (u *stealth) FuelPerTurn() int {
	if u.subDive == "Y" {
		return u.fuelPerTurn + 3
	}
	// Eagle's Buff
	return u.fuelPerTurn
}

func NewStealth(m *unit) Unit {
	u := &stealth{
		directUnit{
			Stealth(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Stealth() unit {
	return unit{
		name:           unitnames.Stealth,
		movementType:   movementtypes.A,
		movementPoints: 6,
		vision:         4,
		fuel:           60,
		fuelPerTurn:    5,
		shortRange:     1,
		longRange:      1,
		ammo:           6,
		cost:           25000,
	}
}
