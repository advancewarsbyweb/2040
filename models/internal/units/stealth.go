package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type stealth struct {
	directUnit
}

func (u *stealth) GetFuelPerTurn() int {
	if u.SubDive == "Y" {
		return u.FuelPerTurn + 3
	}
	// Eagle's Buff
	return u.FuelPerTurn
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
		Name:           unitnames.Stealth,
		MovementType:   movementtypes.A,
		MovementPoints: 6,
		Vision:         4,
		Fuel:           60,
		FuelPerTurn:    5,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           6,
		Cost:           24000,
	}
}
