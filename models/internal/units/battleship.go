package models

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type battleship struct {
	indirectUnit
}

func NewBattleship(m *unit) Unit {
	u := &battleship{
		indirectUnit{
			Battleship(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Battleship() unit {
	return unit{
		Name:           unitnames.Battleship,
		MovementType:   movementtypes.S,
		MovementPoints: 5,
		Vision:         2,
		Fuel:           99,
		FuelPerTurn:    1,
		ShortRange:     2,
		LongRange:      6,
		Ammo:           9,
		Cost:           28000,
	}
}
