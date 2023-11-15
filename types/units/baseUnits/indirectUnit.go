package baseunits

import (
	"errors"
	"math"

	"github.com/awbw/2040/types"
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

func (att *indirectUnit) Fire(def *types.Unit) error {
	if att.Moved == 1 {
		return errors.New(AttackerAlreadyMoved)
	}
	if att.Ammo == 0 {
		return errors.New(AttackerHasNoAmmo)
	}
	distanceAway := int(math.Abs(float64(def.X)-float64(att.X)) + math.Abs(float64(def.Y)-float64(att.Y)))
	if distanceAway > att.LongRange || distanceAway < att.ShortRange {
		return errors.New(DefenderOutsideOfRange)
	}
	att.Ammo = att.Ammo - 1
	return nil
}
