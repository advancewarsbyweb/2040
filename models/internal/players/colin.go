package models

import (
	"math"

	unitmodels "github.com/awbw/2040/models/units"
)

type colin struct {
	co
}

func (co *colin) DamageBoost(u unitmodels.Unit) int {
	if u.GetPlayer().CoPowerOn != "S" {
		return -10
	}
	boost := math.Floor(float64(u.GetPlayer().Funds)/1000) * 3
	return int(boost)
}

func (co *colin) CostModifier() float64 {
	return 0.8
}

func NewColin() Co {
	return &colin{}
}
