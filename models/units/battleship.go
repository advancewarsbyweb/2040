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
		name:           unitnames.Battleship,
		movementType:   movementtypes.S,
		movementPoints: 5,
		vision:         2,
		fuel:           99,
		fuelPerTurn:    1,
		shortRange:     2,
		longRange:      6,
		ammo:           9,
		cost:           28000,
	}
}
