package unittypes

import (
	"github.com/awbw/2040/models"
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type artillery struct {
	indirectUnit
}

func NewArtillery() BaseUnit {
	return &artillery{
		indirectUnit{},
	}
}
