package factories

import (
	"math/rand"

	"github.com/awbw/2040/models"
	baseunits "github.com/awbw/2040/types/units/baseUnits"
)

type UnitFactory struct{}

var Unit UnitFactory

func NewUnitFactory() UnitFactory {
	return UnitFactory{}
}

func init() {
	Unit = NewUnitFactory()
}

func (f *UnitFactory) Create(gameId int, playerId int) models.Unit {

	gameModel := Game.Create()
	playerModel := Player.Create()
	playerModel.GameID = gameModel.ID

	return models.Unit{
		ID:       rand.Intn(100),
		GameID:   gameId,
		PlayerID: playerId,
		X:        1,
		Y:        1,
	}
}

func (f *UnitFactory) Infantry(gameId int, playerId int) models.Unit {
	u := Unit.Create(gameId, playerId)
	u.BaseUnit = baseunits.Infantry()

	return u
}

func (f *UnitFactory) Artillery(gameId, int, playerId int) models.Unit {
	u := Unit.Create(gameId, playerId)
	u.BaseUnit = baseunits.Artillery()

	return u
}
