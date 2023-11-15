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
	switch u.Player.CoPowerOn {
	case "N":
		return 20
	case "Y":
		return 40
	case "S":
		return 60
	default:
		return 20
	}
}

func NewMax() Co {
	return &max{}
}
