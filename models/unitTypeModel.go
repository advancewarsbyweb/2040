package models

import "gorm.io/gorm"

type UnitType struct {
	gorm.Model
	Name           string `gorm:"column_name:units_name"`
	MovementPoints int    `gorm:"column_name:units_movement_points"`
	MovementType   string `gorm:"column_name:units_movement_type"`
	Vision         int    `gorm:"column_name:units_vision"`
	Fuel           int    `gorm:"column_name:units_fuel"`
	FuelPerTurn    int    `gorm:"column_name:units_fuel_per_turn"`
	ShortRange     int    `gorm:"column_name:units_short_range"`
	LongRange      int    `gorm:"column_name:units_long_range"`
	Cost           int    `gorm:"column_name:units_cost"`
	Ammo           int    `gorm:"column_name:units_ammo"`
}
