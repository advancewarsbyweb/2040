package calculator

import (
	"testing"

	"github.com/awbw/2040/models"
	conames "github.com/awbw/2040/types/cos/names"
	terraintypes "github.com/awbw/2040/types/terrains"
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
		CoPowerOn: "N",
		Co: &models.Co{
			Name: conames.Andy,
			ICo:  co(conames.Andy),
		},
	}
}

func unit(unitName unitnames.UnitName) models.IUnit {
	return &models.Unit{Hp: 10, Ammo: 10, Name: unitName, Player: &p, Tile: &road}
}

func co(coName conames.CoName) models.ICo {
	return &models.Co{Name: coName}
}

func TestCalculatePreviewResultSecondary(t *testing.T) {
	assert := assert.New(t)
	a := unit(unitnames.Infantry)
	d := unit(unitnames.Infantry)

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
	a := unit(unitnames.Neotank)
	d := unit(unitnames.MDTank)

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

func TestCalculatePreviewLuck(t *testing.T) {
	assert := assert.New(t)

	a := unit(unitnames.Infantry)
	a.SetHp(8)
	d := unit(unitnames.Infantry)

	c := NewCalculator(a, d)

	c.CalcPreviewResult()

	// The Luck Value should be 7 instead of 9 for the Attacker (8 Hp)
	assert.Equal(c.Attacker.Preview[0].Max-c.Attacker.Preview[0].Min, 7)
}

/*
func TestCalculatePreviewIndirectUnit(t *testing.T) {
	assert := assert.New(t)
	a := unitmodels.CreateUnitHelper(unitnames.Artillery).SetPlayer(&p).SetTile(&road)
	d := unitmodels.CreateUnitHelper(unitnames.Tank).SetPlayer(&p).SetTile(&road)

	c := NewCalculator(a, d)
	c.CalcPreviewResult()

	assert.Equal(c.Attacker.Preview[0].Min, 70)
	assert.Equal(c.Attacker.Preview[0].Max, 79)

	assert.Equal(c.Defender.Preview[0].Min, 0)
	assert.Equal(c.Defender.Preview[0].Max, 0)
}
*/

func TestCalculatePreviewDefenderDead(t *testing.T) {
	assert := assert.New(t)
	a := unit(unitnames.Neotank)
	d := unit(unitnames.Tank)

	c := NewCalculator(a, d)
	c.CalcPreviewResult()

	assert.Equal(c.Attacker.Preview[0].Min, 105)
	assert.Equal(c.Attacker.Preview[0].Max, 114)

	assert.Equal(c.Defender.Preview[0].Min, 0)
	assert.Equal(c.Defender.Preview[0].Max, 0)
}
