package baseunits

import (
	"github.com/awbw/2040/models"
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type artillery struct {
	indirectUnit
}

func NewArtillery() Unit {
	return &artillery{
		indirectUnit{
			unit{
				BaseUnit: models.BaseUnit{
					Name:           unitnames.Artillery,
					MovementType:   movementtypes.T,
					MovementPoints: 5,
					Vision:         1,
					Fuel:           99,
					FuelPerTurn:    0,
					ShortRange:     2,
					LongRange:      3,
					Ammo:           6,
					Cost:           6000,
				},
			},
		},
	}
}
