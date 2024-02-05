package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/types/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type tcopter struct {
	transportUnit
}

func NewTCopter(m *models.Unit) models.IUnit {
	u := &tcopter{
		transportUnit{
			TCopter(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func TCopter() models.Unit {
	return models.Unit{
		Name:         unitnames.TCopter,
		MovementType: movementtypes.A,
		Vision:       1,
		Fuel:         99,
		FuelPerTurn:  2,
		ShortRange:   0,
		LongRange:    0,
		Ammo:         0,
		Cost:         5000,
	}
}
