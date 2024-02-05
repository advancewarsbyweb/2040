package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/types/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type mech struct {
	directUnit
}

func (u *mech) Load() {
}

func NewMech(m *models.Unit) models.IUnit {
	u := &mech{
		directUnit{
			Mech(),
		},
	}
	u.SetUnitProperties(m)
	u.IUnit = u
	return u
}

func Mech() models.Unit {
	return models.Unit{
		Name:           unitnames.Mech,
		MovementType:   movementtypes.B,
		MovementPoints: 2,
		Vision:         2,
		Fuel:           99,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           3,
		Cost:           3000,
	}
}
