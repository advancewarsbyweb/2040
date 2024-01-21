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
}

type co struct {
	Power      string
	SuperPower string
}

type Boost string

var (
	DamageBoost   Boost = "DamageBoost"
	MovementBoost Boost = "MovementBoost"
	RangeBoost    Boost = "RangeBoost"
	CaptureBoost  Boost = "CaptureBoost"
)

var CoMaker map[conames.CoName]func() Co

func init() {
	CoMaker = map[conames.CoName]func() Co{
		conames.Max:  NewMax,
		conames.Sami: NewSami,
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
