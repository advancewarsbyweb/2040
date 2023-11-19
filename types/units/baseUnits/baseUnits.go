package unittypes

import (
	"github.com/awbw/2040/models"
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type UnitFunction func() BaseUnit

var MakeUnit map[unitnames.UnitName]UnitFunction

type BaseUnit interface {
	Move()
	Fire(a Unit, d Unit) error
	Load()

	// Required to obtain base fuel for units and set it to the unit type
	GetAmmo() int
	GetFuel() int
	GetFuelPerTurn() int
	GetMovementType() movementtypes.MovementType
	GetMovementPoint() int
	GetVision() int
	GetShortRange() int
	GetLongRange() int
	GetCost() int
}

type baseUnit struct {
	models.BaseUnit
}

func init() {
	MakeUnit = map[unitnames.UnitName]UnitFunction{
		unitnames.Infantry:  NewInfantry,
		unitnames.Mech:      NewMech,
		unitnames.Artillery: NewArtillery,
	}
}

func (u *baseUnit) Move() {
}

func (u *baseUnit) Fire(a Unit, b Unit) error {
	return nil
}

func (u *baseUnit) Load() {
}

func (u *baseUnit) GetAmmo() int {
	return u.Ammo
}

func (u *baseUnit) GetFuel() int {
	return u.Fuel
}

func (u *baseUnit) GetMovementType() movementtypes.MovementType {
	return u.MovementType
}

func (u *baseUnit) GetMovementPoint() int {
	return u.MovementPoints
}

func (u *baseUnit) GetVision() int {
	return u.Vision
}

func (u *baseUnit) GetFuelPerTurn() int {
	return u.FuelPerTurn
}

func (u *baseUnit) GetShortRange() int {
	return u.ShortRange
}

func (u *baseUnit) GetLongRange() int {
	return u.LongRange
}

func (u *baseUnit) GetCost() int {
	return u.Cost
}
