package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/types/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type apc struct {
	transportUnit
}

func NewAPC(m *models.Unit) models.IUnit {
	u := &apc{
		transportUnit{
			APC(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func APC() models.Unit {
	return models.Unit{
		Name:         unitnames.APC,
		MovementType: movementtypes.T,
		Vision:       1,
		Fuel:         99,
		FuelPerTurn:  0,
		ShortRange:   0,
		LongRange:    0,
		Ammo:         0,
		Cost:         5000,
	}
}
