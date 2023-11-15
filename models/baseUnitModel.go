package models

import (
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type BaseUnit struct {
	Name           unitnames.UnitName         `db:"units_name"`
	MovementPoints int                        `db:"units_movement_points"`
	MovementType   movementtypes.MovementType `db:"units_movement_type"`
	Vision         int                        `db:"units_vision"`
	Fuel           int                        `db:"units_fuel"`
	FuelPerTurn    int                        `db:"units_fuel_per_turn"`
	ShortRange     int                        `db:"units_short_range"`
	LongRange      int                        `db:"units_long_range"`
	Cost           int                        `db:"units_cost"`
	Ammo           int                        `db:"units_ammo"`
}
