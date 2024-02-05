package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/models/internal/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type recon struct {
	directUnit
}

func NewRecon(m *models.Unit) models.IUnit {
	u := &recon{
		directUnit{
			Recon(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Recon() models.Unit {
	return models.Unit{
		Name:           unitnames.Recon,
		MovementType:   movementtypes.W,
		MovementPoints: 8,
		Vision:         5,
		Fuel:           99,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           -1,
		Cost:           4000,
	}
}
