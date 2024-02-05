package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/models/internal/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type fighter struct {
	directUnit
}

func NewFighter(m *models.Unit) models.IUnit {
	u := &fighter{
		directUnit{
			Fighter(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Fighter() models.Unit {
	return models.Unit{
		Name:           unitnames.Fighter,
		MovementType:   movementtypes.A,
		MovementPoints: 9,
		Vision:         2,
		Fuel:           99,
		FuelPerTurn:    5,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           9,
		Cost:           20000,
	}
}
