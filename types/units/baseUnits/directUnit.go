package baseunits

import (
	"errors"
	"math"

	unitnames "github.com/awbw/2040/types/units/names"
)

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
		unitnames.Sub,
		unitnames.Cruiser,
	}
}

func (att *directUnit) Fire(def Unit) error {
	distanceAway := int(math.Abs(float64(def.GetX())-float64(att.GetY())) + math.Abs(float64(def.GetX())-float64(att.GetY())))
	if distanceAway > att.LongRange || distanceAway < att.ShortRange {
		return errors.New(DefenderOutsideOfRange)
	}
	ammo := att.GetAmmo()
	if ammo >= 0 {
		att.SetAmmo(ammo - 1)
		return nil
	}
	return nil
}
