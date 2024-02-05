package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/types/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type battleship struct {
	indirectUnit
}

func NewBattleship(m *models.Unit) models.IUnit {
	u := &battleship{
		indirectUnit{
			Battleship(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Battleship() models.Unit {
	return models.Unit{
		Name:           unitnames.Battleship,
		MovementType:   movementtypes.S,
		MovementPoints: 5,
		Vision:         2,
		Fuel:           99,
		FuelPerTurn:    1,
		ShortRange:     2,
		LongRange:      6,
		Ammo:           9,
		Cost:           28000,
	}
}
