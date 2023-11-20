package db

import (
	"math/rand"

	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types"
	movementtypes "github.com/awbw/2040/types/movements"
	unitnames "github.com/awbw/2040/types/units/names"
)

type UnitFunction func() models.BaseUnit
type unitFactory struct {
	Unit models.Unit
}

var (
	BaseUnits   map[unitnames.UnitName]UnitFunction
	UnitFactory unitFactory
)

func NewUnitFactory() unitFactory {
	return unitFactory{}
}

func init() {
	BaseUnits = map[unitnames.UnitName]UnitFunction{
		unitnames.Infantry:  Infantry,
		unitnames.Mech:      Mech,
		unitnames.Recon:     Recon,
		unitnames.TCopter:   TCopter,
		unitnames.APC:       APC,
		unitnames.Artillery: Artillery,
		unitnames.Tank:      Tank,
	}
	UnitFactory = NewUnitFactory()
}

func (f *unitFactory) Create(name unitnames.UnitName) *unitFactory {
	bu := BaseUnits[name]()
	u := models.Unit{
		ID:       rand.Intn(100),
		GameID:   1,
		PlayerID: 1,
		X:        1,
		Y:        1,
		HP:       10,
		BaseUnit: bu,
	}
	f.Unit = u
	return f
}

func (f *unitFactory) CreateRelations() *unitFactory {
	p := PlayerFactory.Create().BuildInsert()
	g := Game.Create().BuildInsert()

	f.Unit.Game = &models.Game{}
	f.Unit.GameID = g.ID

	f.Unit.Player = &models.Player{}
	f.Unit.PlayerID = p.ID

	return f
}

func (f *unitFactory) SetGame(g *types.Game) *unitFactory {
	f.Unit.Game = &models.Game{}
	f.Unit.GameID = g.ID
	return f
}

func (f *unitFactory) SetPlayer(p *types.Player) *unitFactory {
	f.Unit.Player = &models.Player{}
	f.Unit.PlayerID = p.ID
	return f
}

func (f *unitFactory) Build() models.Unit {
	return f.Unit
}

func (f *unitFactory) BuildInsert() models.Unit {
	// Add UnitRepo create function here
	return f.Unit
}

func (f *unitFactory) SetBaseUnit(name unitnames.UnitName) *unitFactory {
	f.Unit.BaseUnit = BaseUnits[name]()
	return f
}

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

func Tank() models.BaseUnit {
	return models.BaseUnit{
		Name:           unitnames.Tank,
		MovementType:   movementtypes.T,
		MovementPoints: 6,
		Vision:         3,
		Fuel:           70,
		FuelPerTurn:    0,
		ShortRange:     1,
		LongRange:      1,
		Ammo:           9,
		Cost:           7000,
	}
}
