package unittypes

import (
	"github.com/awbw/2040/models"
)

type Unit interface {
    Move()
    Fire()
}

type unit struct {
	models.Unit
}

func NewUnit(u models.Unit) Unit {
	return Unit{
		u,
	}
}


