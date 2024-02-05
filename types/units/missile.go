package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/models/internal/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type missile struct {
	indirectUnit
}

func NewMissile(m *models.Unit) models.IUnit {
	u := &missile{
		indirectUnit{
			Missile(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Missile() models.Unit {
	return models.Unit{
		Name:           unitnames.Missile,
		MovementType:   movementtypes.W,
		MovementPoints: 4,
		Vision:         5,
		Fuel:           50,
		FuelPerTurn:    0,
		ShortRange:     3,
		LongRange:      5,
		Ammo:           6,
		Cost:           12000,
	}
}
