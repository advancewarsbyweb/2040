package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/types/units/names"
)

type UnitFunction func(m *models.Unit) models.IUnit

var UnitMaker map[unitnames.UnitName]UnitFunction

func init() {
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
}

// Create a Unit from a model retrieved from the database
func CreateUnit(m *models.Unit) models.IUnit {
	return UnitMaker[m.Name](m)
}

// Create a Unit from a name
func CreateUnitHelper(name unitnames.UnitName) models.IUnit {
	u := &models.Unit{
		Hp:   10,
		Name: name,
	}
	return CreateUnit(u)
}
