package cos

import (
	"github.com/awbw/2040/types"
	baseunits "github.com/awbw/2040/types/units/baseUnits"
	unitnames "github.com/awbw/2040/types/units/names"
	"golang.org/x/exp/slices"
)

type sami struct {
	co
}

func (co *sami) DamageBoost(u *types.Unit) int {
	if u.Name != unitnames.Infantry && u.Name != unitnames.Mech {
		if slices.Contains(baseunits.IndirectUnits, u.Name) {
			return 0
		}
		return -10
	}
	return PowerBoost(u.Player.CoPowerOn, 30, 50, 70)
}

func (co *sami) MovementBoost(u *types.Unit) int {
	if slices.Contains(baseunits.TransportUnits, u.Name) {
		return 1
	}
	if u.Name != unitnames.Infantry && u.Name != unitnames.Mech {
		return 0
	}
	return PowerBoost(u.Player.CoPowerOn, 0, 1, 2)
}

func (co *sami) CaptureBoost(u *types.Unit) int {
	if u.Name != unitnames.Infantry && u.Name != unitnames.Artillery {
		return 0
	}
	capt := int(u.HP * 1.5)
	return PowerBoost(u.Player.CoPowerOn, capt, capt, 20)
}

func NewSami() Co {
	return &sami{}
}
