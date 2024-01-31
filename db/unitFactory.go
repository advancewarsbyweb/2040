package db

import (
	"github.com/awbw/2040/models"
	unitmodels "github.com/awbw/2040/models/units"
	unitnames "github.com/awbw/2040/models/units/names"
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
