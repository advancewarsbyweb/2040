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
	switch u.Player.CoPowerOn {
	case "N":
		return 30
	case "Y":
		return 50
	case "S":
		return 70
	}
	return 0
}

func (co *sami) MovementBoost(u *types.Unit) int {
	if slices.Contains(baseunits.TransportUnits, u.Name) {
		return 1
	}
	if u.Name != unitnames.Infantry && u.Name != unitnames.Mech {
		return 0
	}
	switch u.Player.CoPowerOn {
	case "N":
		return 0
	case "Y":
		return 1
	case "S":
		return 2
	}
	return 0
}

func NewSami() Co {
	return &sami{}
}
