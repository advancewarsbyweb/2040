package types

import (
	"github.com/awbw/2040/models"
)

type Unit struct {
	models.Unit
}

func NewUnit(u models.Unit) Unit {
	return Unit{
		u,
	}
}
