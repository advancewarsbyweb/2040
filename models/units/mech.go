package models

import (
	"fmt"

	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type mech struct {
	directUnit
}

func (u *mech) Load() {
}

func NewMech(m *unit) Unit {
	fmt.Printf("MECH")
	u := &mech{
		directUnit{
			unit{
				baseUnit: Mech(),
			},
		},
	}
	u.SetUnitProperties(m)
	return u
}

func Mech() baseUnit {
	return baseUnit{
		name:           unitnames.Mech,
		movementType:   movementtypes.B,
		movementPoints: 2,
		vision:         2,
		fuel:           99,
		fuelPerTurn:    0,
		shortRange:     1,
		longRange:      1,
		ammo:           3,
		cost:           3000,
	}
}
