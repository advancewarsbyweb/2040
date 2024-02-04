package unitmodels

import (
	unitnames "github.com/awbw/2040/models/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type infantry struct {
	directUnit
}

func (u *infantry) Load() {
}

// Create an Infantry from a Unit model
func NewInfantry(m *unit) Unit {
	u := &infantry{
		directUnit{
			Infantry(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Infantry() unit {
	return unit{
		Name:           unitnames.Infantry,
		MovementType:   movementtypes.F,
		MovementPoints: 3,
		Vision:         2,
		Fuel:           99,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           -1,
		Cost:           1000,
	}
}
