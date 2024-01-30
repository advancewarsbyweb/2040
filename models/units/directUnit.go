package models

import (
	"errors"
	"math"

	unitnames "github.com/awbw/2040/models/units/names"
)

// This might need to have baseUnit embedded in it directly for the inheritance to work
type directUnit struct {
	unit
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

func (u *directUnit) Fire(a Unit, d Unit) error {
	distanceAway := int(math.Abs(float64(d.X())-float64(a.X())) + math.Abs(float64(d.Y())-float64(a.Y())))
	if distanceAway > a.ShortRange() || distanceAway < a.LongRange() {
		return errors.New(DefenderOutsideOfRange)
	}
	ammo := a.Ammo()
	if ammo >= 0 {
		a.SetAmmo(ammo - 1)
	}
	d.SetHp(9)
	return nil
}
