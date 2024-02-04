package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/models/internal/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type piperunner struct {
	indirectUnit
}

func NewPiperunner(m *models.Unit) models.IUnit {
	u := &piperunner{
		indirectUnit{
			Piperunner(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Piperunner() models.Unit {
	return models.Unit{
		Name:           unitnames.Piperunner,
		MovementType:   movementtypes.P,
		MovementPoints: 9,
		Vision:         4,
		Fuel:           99,
		FuelPerTurn:    0,
		ShortRange:     2,
		LongRange:      5,
		Ammo:           6,
		Cost:           20000,
	}
}
