package models

import (
	"math"

	"github.com/awbw/2040/models"
)

type colin struct {
	models.Co
}

func (co *colin) DamageBoost(u models.IUnit) int {
	if u.GetPlayer().CoPowerOn != "S" {
		return -10
	}
	boost := math.Floor(float64(u.GetPlayer().Funds)/1000) * 3
	return int(boost)
}

func (co *colin) CostModifier() float64 {
	return 0.8
}

func NewColin() models.ICo {
	return &colin{}
}
