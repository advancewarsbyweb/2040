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
	DefenseModifier() int
}

type Co struct {
	Name       conames.CoName
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

func (co *Co) LuckRange() (int, int) {
	return 0, 9
}

func (co *Co) CostModifier() float64 {
	return 1.0
}

func (co *Co) DefenseModifier() int {
	return 0
}
