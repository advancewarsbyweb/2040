package models

import (
	terrainmodels "github.com/awbw/2040/models/internal/terrains"
	unitnames "github.com/awbw/2040/models/internal/units/names"
	movementtypes "github.com/awbw/2040/types/movements"
)

type Unit struct {
	ID             int                        `db:"units_id" json:"id"`
	PlayerID       int                        `db:"units_players_id" json:"playerId"`
	GameID         int                        `db:"units_games_id" json:"gameId"`
	X              int                        `db:"units_x" json:"x"`
	Y              int                        `db:"units_y" json:"y"`
	SubDive        string                     `db:"units_sub_dive" json:"subDive"`
	Moved          int                        `db:"units_moved" json:"moved"`
	Fired          int                        `db:"units_fired" json:"fired"`
	Hp             float64                    `db:"units_hp" json:"hp"`
	Cargo1ID       int                        `db:"units_cargo1_units_id" json:"cargo1Id"`
	Cargo2ID       int                        `db:"units_cargo2_units_id" json:"cargo2Id"`
	Carried        string                     `db:"units_carried" json:"carried"`
	Game           *Game                      `db:""`
	Player         *Player                    `db:""`
	Tile           *terrainmodels.Tile        `db:""`
	Name           unitnames.UnitName         `db:"units_name" json:"name"`
	MovementPoints int                        `db:"units_movement_points" json:"movementPoints"`
	MovementType   movementtypes.MovementType `db:"units_movement_type" json:"movementType"`
	Vision         int                        `db:"units_vision" json:"vision"`
	Fuel           int                        `db:"units_fuel" json:"fuel"`
	FuelPerTurn    int                        `db:"units_fuel_per_turn" json:"fuelPerTurn"`
	ShortRange     int                        `db:"units_short_range" json:"shortRange"`
	LongRange      int                        `db:"units_long_range" json:"longRange"`
	Cost           int                        `db:"units_cost" json:"cost"`
	Ammo           int                        `db:"units_ammo" json:"ammo"`
	// Reference to the created Unit type (e.g: Infantry, Mech, etc)
	// This is to be able to return the proper struct and chain methods when using Setters
	IUnit IUnit
}

/*
type UnitFunction func(m *Unit) IUnit

var UnitMaker map[unitnames.UnitName]UnitFunction
*/

func init() {
	/*
		UnitMaker = map[unitnames.UnitName]UnitFunction{
			unitnames.Infantry:   NewInfantry,
			unitnames.Mech:       NewMech,
			unitnames.Recon:      NewRecon,
			unitnames.APC:        NewAPC,
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
	*/
}

// Create a Unit from a model retrieved from the database
/*
func CreateUnit(m *Unit) IUnit {
	return UnitMaker[m.Name](m)
}

// Create a Unit from a name
func CreateUnitHelper(name unitnames.UnitName) IUnit {
	u := &Unit{
		Hp:   10,
		Name: name,
	}
	return CreateUnit(u)
}
*/

// Set the Unit's properties with the model received from the database
// This is used after creating a Unit type struct to have the appropriate Unit methods
// For example in NewInfantry
func (u *Unit) SetUnitProperties(m *Unit) {
	u.ID = m.ID
	u.PlayerID = m.PlayerID
	u.GameID = m.GameID
	u.X = m.X
	u.Y = m.Y
	u.SubDive = m.SubDive
	u.Moved = m.Moved
	u.Fired = m.Fired
	u.Hp = m.Hp
	u.Cargo1ID = m.Cargo1ID
	u.Cargo2ID = m.Cargo2ID
	u.Carried = m.Carried
	u.Game = m.Game
	u.Player = m.Player
	u.Fuel = m.Fuel
}
