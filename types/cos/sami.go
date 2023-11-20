package cos

import (
	baseunits "github.com/awbw/2040/types/units/baseUnits"
	unittypes "github.com/awbw/2040/types/units/baseUnits"
	unitnames "github.com/awbw/2040/types/units/names"
	"golang.org/x/exp/slices"
)

type sami struct {
	co
}

func (co *sami) DamageBoost(u unittypes.Unit) int {
	if u.GetName() != unitnames.Infantry && u.GetName() != unitnames.Mech {
		if slices.Contains(baseunits.IndirectUnits, u.GetName()) {
			return 0
		}
		return -10
	}
	return PowerBoost(u.GetPlayer().CoPowerOn, 30, 50, 70)
}

func (co *sami) MovementBoost(u unittypes.Unit) int {
	if slices.Contains(baseunits.TransportUnits, u.GetName()) {
		return 1
	}
	if u.GetName() != unitnames.Infantry && u.GetName() != unitnames.Mech {
		return 0
	}
	return PowerBoost(u.GetPlayer().CoPowerOn, 0, 1, 2)
}

func (co *sami) CaptureBoost(u unittypes.Unit) int {
	if u.GetName() != unitnames.Infantry && u.GetName() != unitnames.Artillery {
		return 0
	}
	capt := int(u.GetHp() * 1.5)
	return PowerBoost(u.GetPlayer().CoPowerOn, capt, capt, 20)
}

func NewSami() Co {
	return &sami{}
}
