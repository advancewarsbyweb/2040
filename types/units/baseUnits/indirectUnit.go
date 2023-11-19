package unittypes

import (
	"errors"
	"math"

	unitnames "github.com/awbw/2040/types/units/names"
)

type indirectUnit struct {
	baseUnit
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
	if a.GetMoved() == 1 {
		return errors.New(AttackerAlreadyMoved)
	}
	ammo := a.GetAmmo()
	if ammo == 0 {
		return errors.New(AttackerHasNoAmmo)
	}
	distanceAway := int(math.Abs(float64(d.GetX())-float64(a.GetX())) + math.Abs(float64(d.GetY())-float64(a.GetY())))
	if distanceAway > a.GetLongRange() || distanceAway < a.GetShortRange() {
		return errors.New(DefenderOutsideOfRange)
	}
	a.SetAmmo(ammo - 1)
	d.SetHp(9)
	return nil
}
