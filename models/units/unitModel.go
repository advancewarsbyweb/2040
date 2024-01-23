package models

import (
	"github.com/awbw/2040/models"
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type unit struct {
	id             int
	playerID       int
	gameID         int
	x              int
	y              int
	subDive        string
	moved          int
	fired          int
	hp             float64
	cargo1ID       int
	cargo2ID       int
	carried        string
	game           *models.Game               `db:""`
	player         *models.Player             `db:""`
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
	// Reference to the created Unit type (e.g: Infantry, Mech, etc)
	// This is to be able to return the proper struct and chain methods when using Setters
	IUnit Unit
}

type UnitFunction func(m *unit) Unit

var UnitMaker map[unitnames.UnitName]UnitFunction

func init() {
	UnitMaker = map[unitnames.UnitName]UnitFunction{
		unitnames.Infantry:   NewInfantry,
		unitnames.Mech:       NewMech,
		unitnames.Recon:      NewRecon,
		unitnames.Artillery:  NewArtillery,
		unitnames.Tank:       NewTank,
		unitnames.AntiAir:    NewAntiAir,
		unitnames.BCopter:    NewBCopter,
		unitnames.Missile:    NewMissile,
		unitnames.Rocket:     NewRocket,
		unitnames.MDTank:     NewMDTank,
		unitnames.Cruiser:    NewCruiser,
		unitnames.Sub:        NewSub,
		unitnames.Fighter:    NewFighter,
		unitnames.Piperunner: NewPiperunner,
		unitnames.Neotank:    NewNeotank,
		unitnames.Bomber:     NewBomber,
		unitnames.Stealth:    NewStealth,
		unitnames.Battleship: NewBattleship,
		unitnames.MegaTank:   NewMegaTank,
		unitnames.Carrier:    NewCarrier,
	}
}

// Create a Unit from a model retrieved from the database
func CreateUnit(m *unit) Unit {
	return UnitMaker[m.name](m)
}

// Create a Unit from a name
func CreateUnitHelper(name unitnames.UnitName) Unit {
	u := &unit{
		hp:   10,
		name: name,
	}
	return CreateUnit(u)
}

// Set the Unit's properties with the model received from the database
// This is used after creating a Unit type struct to have the appropriate Unit methods
// For example in NewInfantry
func (u *unit) SetUnitProperties(m *unit) {
	u.id = m.id
	u.playerID = m.playerID
	u.gameID = m.gameID
	u.x = m.x
	u.y = m.y
	u.subDive = m.subDive
	u.moved = m.moved
	u.fired = m.fired
	u.hp = m.hp
	u.cargo1ID = m.cargo1ID
	u.cargo2ID = m.cargo2ID
	u.carried = m.carried
	u.game = m.game
	u.player = m.player
	u.fuel = m.fuel
}
