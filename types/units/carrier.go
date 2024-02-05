package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/types/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type carrier struct {
	indirectUnit
}

func NewCarrier(m *models.Unit) models.IUnit {
	u := &carrier{
		indirectUnit{
			Carrier(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Carrier() models.Unit {
	return models.Unit{
		Name:           unitnames.Carrier,
		MovementType:   movementtypes.S,
		MovementPoints: 5,
		Vision:         4,
		Fuel:           99,
		FuelPerTurn:    1,
		ShortRange:     3,
		LongRange:      8,
		Ammo:           9,
		Cost:           30000,
	}
}
