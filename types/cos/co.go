package cos

import (
	"math/rand"

	conames "github.com/awbw/2040/types/cos/names"
	unittypes "github.com/awbw/2040/types/units/baseUnits"
)

type Co interface {
	DamageBoost(u unittypes.Unit) int
	MovementBoost(u unittypes.Unit) int
	RangeBoost(u unittypes.Unit) int
	CaptureBoost(u unittypes.Unit) int
	LuckRange() int
}

type co struct {
	Power      string
	SuperPower string
}

var CoMakers map[conames.CoName]func() Co

func init() {
	CoMakers = map[conames.CoName]func() Co{
		conames.Max:  NewMax,
		conames.Sami: NewSami,
	}
}

func NewCo(name conames.CoName) Co {
	return CoMakers[name]()
}

func (co *co) DamageBoost(u unittypes.Unit) int {
	return 0
}

func (co *co) MovementBoost(u unittypes.Unit) int {
	return 0
}

func (co *co) RangeBoost(u unittypes.Unit) int {
	return 0
}

func (co *co) CaptureBoost(u unittypes.Unit) int {
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
