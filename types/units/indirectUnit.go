package unitmodels

import (
	"errors"
	"math"

	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types/calculator"
	unitnames "github.com/awbw/2040/types/units/names"
)

type indirectUnit struct {
	models.Unit
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

func (u *indirectUnit) Fire(a models.IUnit, d models.IUnit) error {
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

	c := calculator.NewCalculator(a, d)
	_ = c.CalcResult()

	defHp := (c.Defender.Unit.GetHp()*10 - float64(c.Attacker.Damage)) / 10
	d.SetHp(math.Ceil(defHp))
	if c.Defender.UseAmmo == true {
		d.SetAmmo(ammo - 1)
	}
	return nil
}
