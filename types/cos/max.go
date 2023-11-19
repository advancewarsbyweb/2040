package cos

import (
	"github.com/awbw/2040/types"
	baseunits "github.com/awbw/2040/types/units/baseUnits"
	"golang.org/x/exp/slices"
)

type max struct {
	co
}

func (co *max) DamageBoost(u *types.Unit) int {
	if !slices.Contains(baseunits.DirectUnits, u.Name) {
		if slices.Contains(baseunits.IndirectUnits, u.Name) {
			return -10
		}
		return 0
	}
	return PowerBoost(u.Player.CoPowerOn, 20, 40, 60)
}

func (co *max) MovementBoost(u *types.Unit) int {
	if !slices.Contains(baseunits.DirectUnits, u.Name) {
		return 0
	}
	return PowerBoost(u.Player.CoPowerOn, 0, 1, 2)
}

func (co *max) RangeBoost(u *types.Unit) int {
	if slices.Contains(baseunits.IndirectUnits, u.Name) {
		return -1
	}
	return 0
}

func NewMax() Co {
	return &max{}
}
