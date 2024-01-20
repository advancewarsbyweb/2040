package db

import (
	"github.com/awbw/2040/models"
	unitmodels "github.com/awbw/2040/models/units"
	unitnames "github.com/awbw/2040/types/units/names"
)

type unitFactory struct {
	Unit unitmodels.Unit
}

var (
	UnitFactory unitFactory
)

func NewUnitFactory() unitFactory {
	return unitFactory{}
}

func init() {
	UnitFactory = NewUnitFactory()
}

func (f *unitFactory) Create(name unitnames.UnitName) *unitFactory {
	u := unitmodels.CreateUnitHelper(name)
	u.SetGame(&Game.Create().Game).
		SetPlayer(&PlayerFactory.Create().Player).
		SetPos(1, 1).
		SetHp(10).
		SetMoved(0)
	f.Unit = u
	return f
}

func (f *unitFactory) CreateRelations() *unitFactory {
	p := PlayerFactory.Create().BuildInsert()
	g := Game.Create().BuildInsert()

	f.Unit.SetGame(&g)
	f.Unit.SetPlayer(&p)

	return f
}

func (f *unitFactory) SetGame(g *models.Game) *unitFactory {
	f.Unit.SetGame(g)
	return f
}

func (f *unitFactory) SetPlayer(p *models.Player) *unitFactory {
	f.Unit.SetPlayer(p)
	return f
}

func (f *unitFactory) Build() unitmodels.Unit {
	return f.Unit
}

func (f *unitFactory) BuildInsert() unitmodels.Unit {
	// Add UnitRepo create function here
	return f.Unit
}

/*
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
*/
