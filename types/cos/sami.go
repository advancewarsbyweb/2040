package cos

import (
	"math"

	unitmodels "github.com/awbw/2040/models/units"
	unitnames "github.com/awbw/2040/types/units/names"
	"golang.org/x/exp/slices"
)

type sami struct {
	co
}

func (co *sami) DamageBoost(u unitmodels.Unit) int {
	if u.Name() != unitnames.Infantry && u.Name() != unitnames.Mech {
		if slices.Contains(unitmodels.IndirectUnits, u.Name()) {
			return 0
		}
		return -10
	}
	return PowerBoost(u.Player().CoPowerOn, 30, 50, 70)
}

func (co *sami) MovementBoost(u unitmodels.Unit) int {
	if slices.Contains(unitmodels.TransportUnits, u.Name()) {
		return 1
	}
	if u.Name() != unitnames.Infantry && u.Name() != unitnames.Mech {
		return 0
	}
	return PowerBoost(u.Player().CoPowerOn, 0, 1, 2)
}

func (co *sami) CaptureBoost(u unitmodels.Unit) int {
	if u.Name() != unitnames.Infantry && u.Name() != unitnames.Mech {
		return 0
	}
	capt := int(math.Floor(u.Hp()*1.5)) - int(u.Hp())
	return PowerBoost(u.Player().CoPowerOn, capt, capt, 20)
}

func NewSami() Co {
	return &sami{}
}
