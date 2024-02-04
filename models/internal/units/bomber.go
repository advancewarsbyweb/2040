package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/models/internal/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type bomber struct {
	directUnit
}

func NewBomber(m *models.Unit) models.IUnit {
	u := &bomber{
		directUnit{
			Bomber(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Bomber() models.Unit {
	return models.Unit{
		Name:           unitnames.Bomber,
		MovementType:   movementtypes.A,
		MovementPoints: 8,
		Vision:         2,
		Fuel:           99,
		FuelPerTurn:    5,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           9,
		Cost:           22000,
	}
}
