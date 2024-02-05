package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/types/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type neotank struct {
	directUnit
}

func NewNeotank(m *models.Unit) models.IUnit {
	u := &neotank{
		directUnit{
			Neotank(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Neotank() models.Unit {
	return models.Unit{
		Name:           unitnames.Neotank,
		MovementType:   movementtypes.T,
		MovementPoints: 6,
		Vision:         1,
		Fuel:           99,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           9,
		Cost:           22000,
	}
}
