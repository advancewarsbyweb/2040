package models

import (
	"math/rand"
)

type ICo interface {
	DamageBoost(u IUnit) int
	MovementBoost(u IUnit) int
	RangeBoost(u IUnit) int
	CaptureBoost(u IUnit) int
	LuckRange() int
	CostModifier() float64
}

type Co struct {
	Power      string
	SuperPower string
}

func (co *Co) DamageBoost(u IUnit) int {
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

func (co *Co) LuckRange() int {
	return rand.Intn(9)
}

func (co *Co) CostModifier() float64 {
	return 1.0
}
