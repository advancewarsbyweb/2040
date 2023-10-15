package baseunits

import (
	"github.com/awbw/2040/models"
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type Infantry interface {
	Move(x int, y int) Infantry
	Fire(x int, y int) Infantry
	Capture(x int, y int) Infantry
	Load(x int, y int) Infantry
}

type infantry struct {
	models.BaseUnit
}

func (u *infantry) Move(x int, y int) Infantry {
	return u
}

func (u *infantry) Fire(x int, y int) Infantry {
	return u
}

func (u *infantry) Capture(x int, y int) Infantry {
	return u
}

func (u *infantry) Load(x int, y int) Infantry {
	return u
}

func NewInfantry() Infantry {
	return &infantry{
		BaseUnit: models.BaseUnit{
			Name:           unitnames.Infantry,
			MovementType:   movementtypes.F,
			MovementPoints: 3,
			Vision:         2,
			Fuel:           99,
			FuelPerTurn:    0,
			ShortRange:     1,
			LongRange:      1,
			Ammo:           -1,
			Cost:           1000,
		},
	}
}
