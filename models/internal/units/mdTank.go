package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/models/internal/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type mdTank struct {
	directUnit
}

func NewMDTank(m *models.Unit) models.IUnit {
	u := &mdTank{
		directUnit{
			MDTank(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func MDTank() models.Unit {
	return models.Unit{
		Name:           unitnames.MDTank,
		MovementType:   movementtypes.T,
		MovementPoints: 5,
		Vision:         1,
		Fuel:           50,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           8,
		Cost:           16000,
	}
}
