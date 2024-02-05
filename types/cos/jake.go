package cos

import (
	"github.com/awbw/2040/models"
	terrainnames "github.com/awbw/2040/types/terrains/names"
	unitmodels "github.com/awbw/2040/types/units"
	"golang.org/x/exp/slices"
)

type jake struct {
	models.Co
}

func (co *jake) DamageBoost(u models.IUnit) int {
	if u.GetTile().Name == terrainnames.Plain || terrainnames.Rubble.Match(u.GetTile().Name) {
		return PowerBoost(u.GetPlayer().CoPowerOn, 10, 20, 40)
	}
	return 0
}

func (co *jake) MovementBoost(u models.IUnit) int {
	if !slices.Contains(unitmodels.Vehicles, u.GetName()) {
		return 0
	}
	return PowerBoost(u.GetPlayer().CoPowerOn, 0, 0, 2)
}

func (co *jake) RangeBoost(u models.IUnit) int {
	if !slices.Contains(unitmodels.IndirectUnits, u.GetName()) {
		return 0
	}
	return PowerBoost(u.GetPlayer().CoPowerOn, 0, 1, 1)
}

func NewJake() models.ICo {
	return &jake{}
}
