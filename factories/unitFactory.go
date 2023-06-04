package factories

import (
	"math/rand"

	"github.com/awbw/2040/models"
)

type UnitFactory struct{}

var Unit UnitFactory

func NewUnitFactory() UnitFactory {
	return UnitFactory{}
}

func init() {
	Unit = NewUnitFactory()
}

func (f *UnitFactory) Create() models.Unit {

	gameModel := Game.Create()

	return models.Unit{
		ID:       rand.Int(),
		GameID:   gameModel.ID,
		PlayerID: Player.Create().ID,
		X:        1,
		Y:        1,
	}
}
