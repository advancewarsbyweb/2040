package models

import (
	terrainnames "github.com/awbw/2040/types/terrains/names"
	unitnames "github.com/awbw/2040/types/units/names"
	"gopkg.in/guregu/null.v4"
)

type IBuilding interface{}

type Building struct {
	ID          int       `db:"buildings_id"`
	GameID      int       `db:"buildings_games_id"`
	TerrainID   int       `db:"buildings_terrain_id"`
	X           int       `db:"buildings_x"`
	Y           int       `db:"buildings_y"`
	Capture     int       `db:"buildings_capture"`
	LastCapture int       `db:"buildings_last_capture"`
	LastUpdated null.Time `db:"buildings_last_updated"`
	Game        *Game     `db:""`
	Tile        *Tile     `db:""`
}

var (
	ProductionBuildings map[terrainnames.TerrainName][]unitnames.UnitName
	Base                []unitnames.UnitName
	Airport             []unitnames.UnitName
	Port                []unitnames.UnitName
)

func init() {
	Base = []unitnames.UnitName{
		unitnames.Infantry, unitnames.Mech, unitnames.Recon, unitnames.Artillery, unitnames.Tank, unitnames.AntiAir, unitnames.Missile, unitnames.Rocket,
		unitnames.MDTank, unitnames.Piperunner, unitnames.Neotank, unitnames.MegaTank,
	}
	Airport = []unitnames.UnitName{unitnames.TCopter, unitnames.BCopter, unitnames.Fighter, unitnames.Bomber, unitnames.Stealth, unitnames.BlackBomb}
	Port = []unitnames.UnitName{unitnames.BlackBoat, unitnames.Lander, unitnames.Cruiser, unitnames.Sub, unitnames.Battleship, unitnames.Carrier}
}

func (b *Building) Build() bool {
	return true
}
