package models

import (
	"math"

	unitmodels "github.com/awbw/2040/models/units"
	unitnames "github.com/awbw/2040/models/units/names"
	"golang.org/x/exp/slices"
)

type sami struct {
	co
}

func (co *sami) DamageBoost(u unitmodels.Unit) int {
	if u.GetName() != unitnames.Infantry && u.GetName() != unitnames.Mech {
		if slices.Contains(unitmodels.IndirectUnits, u.GetName()) {
			return 0
		}
		return -10
	}
	return PowerBoost(u.GetPlayer().CoPowerOn, 30, 50, 70)
}

func (co *sami) MovementBoost(u unitmodels.Unit) int {
	if slices.Contains(unitmodels.TransportUnits, u.GetName()) {
		return 1
	}
	if u.GetName() != unitnames.Infantry && u.GetName() != unitnames.Mech {
		return 0
	}
	return PowerBoost(u.GetPlayer().CoPowerOn, 0, 1, 2)
}

func (co *sami) CaptureBoost(u unitmodels.Unit) int {
	if u.GetName() != unitnames.Infantry && u.GetName() != unitnames.Mech {
		return 0
	}
	capt := int(math.Floor(u.GetHp()*1.5)) - int(u.GetHp())
	return PowerBoost(u.GetPlayer().CoPowerOn, capt, capt, 20)
}

func NewSami() Co {
	return &sami{}
}
