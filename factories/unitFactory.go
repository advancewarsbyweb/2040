package factories

import (
	"math/rand"

	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types"
	baseunits "github.com/awbw/2040/types/units/baseUnits"
)

type UnitFactory struct {
	Unit types.Unit
}

var Unit UnitFactory

func NewUnitFactory() UnitFactory {
	return UnitFactory{}
}

func init() {
	Unit = NewUnitFactory()
}

func (f *UnitFactory) Create() *UnitFactory {
	u := types.NewUnit(models.Unit{
		ID:       rand.Intn(100),
		GameID:   1,
		PlayerID: 1,
		X:        1,
		Y:        1,
	})
	f.Unit = u
	return f
}

func (f *UnitFactory) CreateRelations() *UnitFactory {
	p := Player.Create().BuildInsert()
	g := Game.Create().BuildInsert()

	f.Unit.GameID = g.ID
	f.Unit.PlayerID = p.ID
	f.Unit.Game = &g
	f.Unit.Player = &p
	return f
}

func (f *UnitFactory) SetGame(g *types.Game) *UnitFactory {
	f.Unit.GameID = g.ID
	f.Unit.Game = g
	return f
}

func (f *UnitFactory) SetPlayer(p *types.Player) *UnitFactory {
	f.Unit.PlayerID = p.ID
	f.Unit.Player = p
	return f
}

func (f *UnitFactory) Build() types.Unit {
	return f.Unit
}

func (f *UnitFactory) BuildInsert() types.Unit {
	// Add UnitRepo create function here
	return f.Unit
}

func (f *UnitFactory) Infantry() *UnitFactory {
	f.Unit.BaseUnit = baseunits.Infantry()
	return f
}

func (f *UnitFactory) Artillery() *UnitFactory {
	f.Unit.BaseUnit = baseunits.Artillery()
	return f
}
