package baseunits

import (
	"github.com/awbw/2040/models"
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

func Infantry() models.BaseUnit {
	return models.BaseUnit{
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
	}
}

func Artillery() models.BaseUnit {
	return models.BaseUnit{
		Name:           unitnames.Artillery,
		MovementType:   movementtypes.T,
		MovementPoints: 5,
		Vision:         1,
		Fuel:           99,
		FuelPerTurn:    0,
		ShortRange:     2,
		LongRange:      3,
		Ammo:           6,
		Cost:           6000,
	}
}
