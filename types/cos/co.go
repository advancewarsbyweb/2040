package cos

import (
	"math/rand"

	"github.com/awbw/2040/types"
)

type Co interface {
	DamageBoost(u *types.Unit) int
	MovementBoost(u *types.Unit) int
	RangeBoost(u *types.Unit) int
	CaptureBoost(u *types.Unit) int
	LuckRange() int
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

func (co *co) RangeBoost(u *types.Unit) int {
	return 0
}

func (co *co) CaptureBoost(u *types.Unit) int {
	return 0
}

func (co *co) LuckRange() int {
	return rand.Intn(9)
}

func PowerBoost(coPowerOn string, noCopBoost int, copBoost int, scopBoost int) int {
	switch coPowerOn {
	case "N":
		return noCopBoost
	case "Y":
		return copBoost
	case "S":
		return scopBoost
	default:
		return noCopBoost
	}
}
