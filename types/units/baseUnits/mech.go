package unittypes

import (
	"github.com/awbw/2040/models"
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type mech struct {
	directUnit
}

func (u *mech) Load() {
}

func NewMech() BaseUnit {
	u := &mech{
		directUnit{
			baseUnit{
				models.BaseUnit{
					Name:           unitnames.Mech,
					MovementType:   movementtypes.B,
					MovementPoints: 2,
					Vision:         2,
					Fuel:           99,
					FuelPerTurn:    0,
					ShortRange:     1,
					LongRange:      1,
					Ammo:           3,
					Cost:           3000,
				},
			},
		},
	}
	return u
}
