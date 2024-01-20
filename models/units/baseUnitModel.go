package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type baseUnit struct {
	name           unitnames.UnitName         `db:"units_name"`
	movementPoints int                        `db:"units_movement_points"`
	movementType   movementtypes.MovementType `db:"units_movement_type"`
	vision         int                        `db:"units_vision"`
	fuel           int                        `db:"units_fuel"`
	fuelPerTurn    int                        `db:"units_fuel_per_turn"`
	shortRange     int                        `db:"units_short_range"`
	longRange      int                        `db:"units_long_range"`
	cost           int                        `db:"units_cost"`
	ammo           int                        `db:"units_ammo"`
}

type UnitFunction func(u *unit) Unit

var UnitMaker map[unitnames.UnitName]UnitFunction

type BaseUnit interface {
	Move()
	Fire(a Unit, d Unit) error
	Load()
}

func init() {
	UnitMaker = map[unitnames.UnitName]UnitFunction{
		unitnames.Infantry:  NewInfantry,
		unitnames.Mech:      NewMech,
		unitnames.Artillery: NewArtillery,
		unitnames.Tank:      NewTank,
	}
}

func (u *unit) Move() {
}

func (u *unit) Fire(a Unit, b Unit) error {
	return nil
}

func (u *unit) Load() {
}
