package unittypes

import (
	unitnames "github.com/awbw/2040/types/units/names"
)

type UnitFunction func() BaseUnit

var MakeUnit map[unitnames.UnitName]UnitFunction

type BaseUnit interface {
	Move()
	Fire(a Unit, d Unit) error
	Load()
}

// Type to implement methods on only
type baseUnit struct{}

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
