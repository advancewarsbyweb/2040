package types

import "github.com/awbw/2040/models"

type PressTo struct {
	models.PressTo
}

func NewPressTo(p models.PressTo) PressTo {
	return PressTo{
		PressTo: p,
	}
}
