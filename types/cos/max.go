package cos

import (
	unitmodels "github.com/awbw/2040/models/units"
	"golang.org/x/exp/slices"
)

type max struct {
	co
}

func (co *max) DamageBoost(u unitmodels.Unit) int {
	if !slices.Contains(unitmodels.DirectUnits, u.Name()) {
		if slices.Contains(unitmodels.IndirectUnits, u.Name()) {
			return -10
		}
		return 0
	}
	return PowerBoost(u.Player().CoPowerOn, 20, 40, 60)
}

func (co *max) MovementBoost(u unitmodels.Unit) int {
	if !slices.Contains(unitmodels.DirectUnits, u.Name()) {
		return 0
	}
	return PowerBoost(u.Player().CoPowerOn, 0, 1, 2)
}

func (co *max) RangeBoost(u unitmodels.Unit) int {
	if slices.Contains(unitmodels.IndirectUnits, u.Name()) {
		return -1
	}
	return 0
}

func NewMax() Co {
	return &max{}
}
