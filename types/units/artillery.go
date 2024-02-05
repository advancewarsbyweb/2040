package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/models/internal/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type artillery struct {
	indirectUnit
}

func NewArtillery(m *models.Unit) models.IUnit {
	u := &artillery{
		indirectUnit{
			Artillery(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Artillery() models.Unit {
	return models.Unit{
		Name:           unitnames.Artillery,
		MovementType:   movementtypes.T,
		MovementPoints: 5,
		Vision:         1,
		Fuel:           99,
		FuelPerTurn:    0,
		ShortRange:     2,
		LongRange:      3,
		Ammo:           6,
		Cost:           6000,
	}
}
