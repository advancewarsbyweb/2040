package cos

import (
	unitmodels "github.com/awbw/2040/models/units"
	"golang.org/x/exp/slices"
)

type jake struct {
	co
}

func (co *jake) DamageBoost(u unitmodels.Unit) int {
	return 0
}

func (co *jake) MovementBoost(u unitmodels.Unit) int {
	return 0
}

func (co *jake) RangeBoost(u unitmodels.Unit) int {
	if !slices.Contains(unitmodels.IndirectUnits, u.Name()) {
		return 0
	}
	return PowerBoost(u.Player().CoPowerOn, 0, 1, 1)
}
