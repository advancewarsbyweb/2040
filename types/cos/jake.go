package cos

import (
	unitmodels "github.com/awbw/2040/models/units"
	terraintypes "github.com/awbw/2040/types/terrains"
	"golang.org/x/exp/slices"
)

type jake struct {
	co
}

func (co *jake) DamageBoost(u unitmodels.Unit) int {
	if u.GetTile().Name == terraintypes.Plain || terraintypes.Rubble.Match(u.GetTile().Name) {
		return PowerBoost(u.GetPlayer().CoPowerOn, 10, 20, 40)
	}
	return 0
}

func (co *jake) MovementBoost(u unitmodels.Unit) int {
	if !slices.Contains(unitmodels.Vehicles, u.GetName()) {
		return 0
	}
	return PowerBoost(u.GetPlayer().CoPowerOn, 0, 0, 2)
}

func (co *jake) RangeBoost(u unitmodels.Unit) int {
	if !slices.Contains(unitmodels.IndirectUnits, u.GetName()) {
		return 0
	}
	return PowerBoost(u.GetPlayer().CoPowerOn, 0, 1, 1)
}

func NewJake() Co {
	return &jake{}
}
