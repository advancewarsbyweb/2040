package cos

import (
	"github.com/awbw/2040/types"
)

type Co interface {
	DamageBoost(u *types.Unit) int
	MovementBoost(u *types.Unit) int
	RangeBoost()
	CaptureBoost(u *types.Unit) int
	LuckRange()
}

type co struct {
	Power      string
	SuperPower string
}

func (co *co) DamageBoost(u *types.Unit) int {
	return 0
}

func (co *co) MovementBoost(u *types.Unit) int {
	return 0
}

func (co *co) RangeBoost() {
}

func (co *co) CaptureBoost(u *types.Unit) int {
	return 0
}

func (co *co) LuckRange() {
}
