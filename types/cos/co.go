package cos

import (
	"math/rand"

	unitmodels "github.com/awbw/2040/models/units"
	conames "github.com/awbw/2040/types/cos/names"
)

type Co interface {
	DamageBoost(u unitmodels.Unit) int
	MovementBoost(u unitmodels.Unit) int
	RangeBoost(u unitmodels.Unit) int
	CaptureBoost(u unitmodels.Unit) int
	LuckRange() int
	CostModifier() float64
}

type co struct {
	Power      string
	SuperPower string
}

var CoMaker map[conames.CoName]func() Co

func init() {
	CoMaker = map[conames.CoName]func() Co{
		conames.Andy:  NewAndy,
		conames.Max:   NewMax,
		conames.Sami:  NewSami,
		conames.Colin: NewColin,
		conames.Jake:  NewJake,
	}
}

func NewCo(name conames.CoName) Co {
	return CoMaker[name]()
}

func (co *co) DamageBoost(u unitmodels.Unit) int {
	return 0
}

func (co *co) MovementBoost(u unitmodels.Unit) int {
	return 0
}

func (co *co) RangeBoost(u unitmodels.Unit) int {
	return 0
}

func (co *co) CaptureBoost(u unitmodels.Unit) int {
	return 0
}

func (co *co) LuckRange() int {
	return rand.Intn(9)
}

func (co *co) CostModifier() float64 {
	return 1.0
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
