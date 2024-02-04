package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/models/internal/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type cruiser struct {
	directUnit
}

func NewCruiser(m *models.Unit) models.IUnit {
	u := &cruiser{
		directUnit{
			Cruiser(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Cruiser() models.Unit {
	return models.Unit{
		Name:           unitnames.Cruiser,
		MovementType:   movementtypes.S,
		MovementPoints: 6,
		Vision:         3,
		Fuel:           99,
		FuelPerTurn:    1,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           9,
		Cost:           18000,
	}
}
