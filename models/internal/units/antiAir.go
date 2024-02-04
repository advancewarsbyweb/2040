package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/models/internal/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type antiAir struct {
	directUnit
}

func NewAntiAir(m *models.Unit) models.IUnit {
	u := &antiAir{
		directUnit{
			AntiAir(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func AntiAir() models.Unit {
	return models.Unit{
		Name:           unitnames.AntiAir,
		MovementType:   movementtypes.T,
		MovementPoints: 6,
		Vision:         2,
		Fuel:           60,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           9,
		Cost:           8000,
	}
}
