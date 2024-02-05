package calculator

import (
	"testing"

	"github.com/awbw/2040/models"
	conames "github.com/awbw/2040/types/cos/names"
	terraintypes "github.com/awbw/2040/types/terrains"
	unitmodels "github.com/awbw/2040/types/units"
	unitnames "github.com/awbw/2040/types/units/names"
	"github.com/stretchr/testify/assert"
)

var (
	plain models.Tile
	road  models.Tile

	p models.Player
)

func init() {
	plain = terraintypes.TestMaps.Tile(1)
	road = terraintypes.TestMaps.Tile(15)

	p = models.Player{
		Co: &models.Co{
			Name: conames.Andy,
		},
	}
}

func TestCalculatePreviewResultSecondary(t *testing.T) {
	assert := assert.New(t)

	a := unitmodels.CreateUnitHelper(unitnames.Infantry).SetPlayer(&p).SetTile(&road)
	d := unitmodels.CreateUnitHelper(unitnames.Infantry).SetPlayer(&p).SetTile(&road)

	c := NewCalculator(a, d)

	c.CalcPreviewResult()

	assert.Equal(c.Attacker.Preview[0].Min, 55)
	assert.Equal(c.Attacker.Preview[0].Max, 64)

	assert.Equal(len(c.Defender.Preview), 2)
	assert.Equal(c.Defender.Preview[0].Min, 22)
	assert.Equal(c.Defender.Preview[0].Max, 25)

	assert.Equal(c.Defender.Preview[1].Min, 27)
	assert.Equal(c.Defender.Preview[1].Max, 32)
}

func TestCalculatePreviewResultPrimary(t *testing.T) {
	assert := assert.New(t)

	a := unitmodels.CreateUnitHelper(unitnames.Neotank).SetPlayer(&p).SetTile(&road)
	d := unitmodels.CreateUnitHelper(unitnames.MDTank).SetPlayer(&p).SetTile(&road)

	c := NewCalculator(a, d)

	c.CalcPreviewResult()

	assert.Equal(c.Attacker.Preview[0].Min, 75)
	assert.Equal(c.Attacker.Preview[0].Max, 84)

	assert.Equal(len(c.Defender.Preview), 2)
	assert.Equal(c.Defender.Preview[0].Min, 9)
	assert.Equal(c.Defender.Preview[0].Max, 10)

	assert.Equal(c.Defender.Preview[1].Min, 13)
	assert.Equal(c.Defender.Preview[1].Max, 16)
}
