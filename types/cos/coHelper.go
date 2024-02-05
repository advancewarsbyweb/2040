package cos

import (
	"github.com/awbw/2040/models"
	conames "github.com/awbw/2040/types/cos/names"
)

var CoMaker map[conames.CoName]func() models.ICo

func init() {
	CoMaker = map[conames.CoName]func() models.ICo{
		conames.Andy:  NewAndy,
		conames.Max:   NewMax,
		conames.Sami:  NewSami,
		conames.Colin: NewColin,
		conames.Jake:  NewJake,
	}
}

func NewCo(name conames.CoName) models.ICo {
	return CoMaker[name]()
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
