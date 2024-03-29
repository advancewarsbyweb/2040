package db

import (
	"github.com/awbw/2040/models"
	unitmodels "github.com/awbw/2040/types/units"
	unitnames "github.com/awbw/2040/types/units/names"
)

type unitFactory struct {
	Unit models.IUnit
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
	u.SetGame(GameFactory.Create().Game).
		SetPlayer(PlayerFactory.Create().Player).
		SetPos(1, 1).
		SetHp(10).
		SetMoved(0)
	f.Unit = u
	return f
}

func (f *unitFactory) CreateRelations() *unitFactory {
	p := PlayerFactory.Create().BuildInsert()
	g := GameFactory.Create().BuildInsert()

	f.Unit.SetGame(g)
	f.Unit.SetPlayer(p)

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

func (f *unitFactory) Build() models.IUnit {
	return f.Unit
}

func (f *unitFactory) BuildInsert() models.IUnit {
	// Add UnitRepo create function here
	return f.Unit
}
