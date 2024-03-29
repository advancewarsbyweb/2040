package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/types/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type megaTank struct {
	directUnit
}

func NewMegaTank(m *models.Unit) models.IUnit {
	u := &megaTank{
		directUnit{
			MegaTank(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func MegaTank() models.Unit {
	return models.Unit{
		Name:           unitnames.MegaTank,
		MovementType:   movementtypes.T,
		MovementPoints: 4,
		Vision:         1,
		Fuel:           50,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           3,
		Cost:           28000,
	}
}
