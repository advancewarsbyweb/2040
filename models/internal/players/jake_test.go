package models

import (
	"fmt"
	"testing"

	conames "github.com/awbw/2040/models/players/coNames"
	terrainmodels "github.com/awbw/2040/models/terrains"
	unitmodels "github.com/awbw/2040/models/units"
	unitnames "github.com/awbw/2040/models/units/names"
	"github.com/stretchr/testify/assert"
)

var (
	jakeTest   Co
	jakePlayer Player
)

func init() {
	jakeTest = CoMaker[conames.Jake]()
	jakePlayer = Player{}
}

func TestDamageBoost_Jake(t *testing.T) {
	assert := assert.New(t)

	// Plain & Pipe Rubble
	plainTiles := []int{1, 115}

	for _, tileId := range plainTiles {
		tile := terrainmodels.TestMaps.Tile(tileId)
		u := unitmodels.CreateUnitHelper(unitnames.Infantry).SetPlayer(&jakePlayer).SetTile(&tile)

		jakePlayer.CoPowerOn = "N"
		assert.Equal(jakeTest.DamageBoost(u), 10, fmt.Sprintf("Jake's %s damage boost is wrong (No COP)", tile.Name))

		jakePlayer.CoPowerOn = "Y"
		assert.Equal(jakeTest.DamageBoost(u), 20, fmt.Sprintf("Jake's %s damage boost is wrong (COP)", tile.Name))

		jakePlayer.CoPowerOn = "S"
		assert.Equal(jakeTest.DamageBoost(u), 40, fmt.Sprintf("Jake's %s damage boost is wrong (SCOP)", tile.Name))
	}
}

func TestMovementBoost_Jake(t *testing.T) {
	assert := assert.New(t)

	for _, unitName := range unitmodels.Vehicles {
		u := unitmodels.CreateUnitHelper(unitName).SetPlayer(&jakePlayer)
		jakePlayer.CoPowerOn = "S"
		assert.Equal(jakeTest.MovementBoost(u), 2, fmt.Sprintf("Jake's %s movement boost is wrong (SCOP)", u.GetName()))
	}
}

func TestRangeBoost_Jake(t *testing.T) {
	assert := assert.New(t)

	for _, unitName := range unitmodels.IndirectUnits {
		u := unitmodels.CreateUnitHelper(unitName).SetPlayer(&jakePlayer)

		jakePlayer.CoPowerOn = "N"
		assert.Equal(jakeTest.RangeBoost(u), 0, fmt.Sprintf("Jake's %s (Indirect Unit) range boost is wrong (No COP)", u.GetName()))

		jakePlayer.CoPowerOn = "Y"
		assert.Equal(jakeTest.RangeBoost(u), 1, fmt.Sprintf("Jake's %s (Indirect Unit) range boost is wrong (COP)", u.GetName()))

		jakePlayer.CoPowerOn = "S"
		assert.Equal(jakeTest.RangeBoost(u), 1, fmt.Sprintf("Jake's %s (Indirect Unit) range boost is wrong (SCOP)", u.GetName()))
	}
}
