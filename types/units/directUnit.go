package unitmodels

import (
	"errors"
	"math"

	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/types/units/names"
)

// This might need to have baseUnit embedded in it directly for the inheritance to work
type directUnit struct {
	models.Unit
}

var DirectUnits []unitnames.UnitName

func init() {
	DirectUnits = []unitnames.UnitName{
		unitnames.Recon,
		unitnames.Tank,
		unitnames.MDTank,
		unitnames.Neotank,
		unitnames.MegaTank,
		unitnames.AntiAir,
		unitnames.BCopter,
		unitnames.Fighter,
		unitnames.Bomber,
		unitnames.Stealth,
		unitnames.Sub,
		unitnames.Cruiser,
	}
}

func (u *directUnit) Fire(a models.IUnit, d models.IUnit) error {
	distanceAway := int(math.Abs(float64(d.GetX())-float64(a.GetX())) + math.Abs(float64(d.GetY())-float64(a.GetY())))
	if distanceAway > a.GetShortRange() || distanceAway < a.GetLongRange() {
		return errors.New(DefenderOutsideOfRange)
	}
	ammo := a.GetAmmo()
	if ammo >= 0 {
		a.SetAmmo(ammo - 1)
	}
	d.SetHp(9)
	return nil
}
