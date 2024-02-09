package models

import (
	conames "github.com/awbw/2040/types/cos/names"
)

type ICo interface {
	DamageBoost(u IUnit) int
	MovementBoost(u IUnit) int
	RangeBoost(u IUnit) int
	CaptureBoost(u IUnit) int
	LuckRange() (int, int)
	CostModifier() float64
	DefenseBoost(u IUnit) int
	TerrainStarsBoost(u IUnit) int
}

type Co struct {
	Name       conames.CoName
	Power      string
	SuperPower string
	ICo        ICo
}

func (co *Co) DamageBoost(u IUnit) int {
	if u.GetPlayer().CoPowerOn != "N" {
		return 10
	}
	return 0
}

func (co *Co) MovementBoost(u IUnit) int {
	return 0
}

func (co *Co) RangeBoost(u IUnit) int {
	return 0
}

func (co *Co) CaptureBoost(u IUnit) int {
	return 0
}

func (co *Co) LuckRange() (int, int) {
	return 0, 9
}

func (co *Co) CostModifier() float64 {
	return 1.0
}

func (co *Co) DefenseBoost(u IUnit) int {
	if u.GetPlayer().CoPowerOn != "N" {
		return 10
	}
	return 0
}

func (co *Co) TerrainStarsBoost(u IUnit) int {
	return 0
}
