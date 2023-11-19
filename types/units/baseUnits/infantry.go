package unittypes

import (
	"github.com/awbw/2040/models"
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type infantry struct {
	directUnit
}

func (u *infantry) Load() {
}

func NewInfantry() BaseUnit {
	return &infantry{
		directUnit{
			baseUnit{
				models.BaseUnit{
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
				},
			},
		},
	}
}
