package models

import (
	"errors"
	"math"

	unitnames "github.com/awbw/2040/types/units/names"
)

type indirectUnit struct {
	unit
}

var IndirectUnits []unitnames.UnitName

func init() {
	IndirectUnits = []unitnames.UnitName{
		unitnames.Artillery,
		unitnames.Rocket,
		unitnames.Missile,
		unitnames.Piperunner,
		unitnames.Battleship,
		unitnames.Carrier,
	}
}

func (u *indirectUnit) Fire(a Unit, d Unit) error {
	if a.Moved() == 1 {
		return errors.New(AttackerAlreadyMoved)
	}
	ammo := a.Ammo()
	if ammo == 0 {
		return errors.New(AttackerHasNoAmmo)
	}
	distanceAway := int(math.Abs(float64(d.X())-float64(a.X())) + math.Abs(float64(d.Y())-float64(a.Y())))
	if distanceAway > a.LongRange() || distanceAway < a.ShortRange() {
		return errors.New(DefenderOutsideOfRange)
	}
	a.SetAmmo(ammo - 1)
	d.SetHp(9)
	return nil
}
