package baseunits

import (
	"github.com/awbw/2040/models"
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type BaseUnit models.BaseUnit

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

func Mech() models.BaseUnit {
	return models.BaseUnit{
		Name:           unitnames.Mech,
		MovementType:   movementtypes.B,
		MovementPoints: 2,
		Vision:         2,
		Fuel:           99,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           3,
		Cost:           3000,
	}
}

func Recon() models.BaseUnit {
	return models.BaseUnit{
		Name:           unitnames.Recon,
		MovementType:   movementtypes.W,
		MovementPoints: 8,
		Vision:         5,
		Fuel:           99,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           -1,
		Cost:           4000,
	}
}

func TCopter() models.BaseUnit {
	return models.BaseUnit{
		Name:         unitnames.TCopter,
		MovementType: movementtypes.A,
		Vision:       1,
		Fuel:         99,
		FuelPerTurn:  2,
		ShortRange:   0,
		LongRange:    0,
		Ammo:         0,
		Cost:         5000,
	}
}

func APC() models.BaseUnit {
	return models.BaseUnit{
		Name:         unitnames.APC,
		MovementType: movementtypes.T,
		Vision:       1,
		Fuel:         99,
		FuelPerTurn:  0,
		ShortRange:   0,
		LongRange:    0,
		Ammo:         0,
		Cost:         5000,
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
