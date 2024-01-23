package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type piperunner struct {
	indirectUnit
}

func NewPiperunner(m *unit) Unit {
	u := &piperunner{
		indirectUnit{
			Piperunner(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Piperunner() unit {
	return unit{
		name:           unitnames.Piperunner,
		movementType:   movementtypes.P,
		movementPoints: 9,
		vision:         4,
		fuel:           99,
		fuelPerTurn:    0,
		shortRange:     2,
		longRange:      5,
		ammo:           6,
		cost:           20000,
	}
}
