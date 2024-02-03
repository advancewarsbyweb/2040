package calculator

import (
	"github.com/awbw/2040/models"
	terrainmodels "github.com/awbw/2040/models/terrains"
	unitmodels "github.com/awbw/2040/models/units"
	"github.com/awbw/2040/types/cos"
)

type Calculator struct {
	Attacker *CalculatorInput
	Defender *CalculatorInput
}

type CalculatorInput struct {
	// Damage and Defense can vary based on the CO, Power, etc
	Damage  int
	Defense int
	Unit    unitmodels.Unit
	Tile    *terrainmodels.Tile
	Player  *models.Player
	Co      cos.Co
}

func (c *Calculator) CalcResult() {
	a := c.Attacker
	d := c.Defender
	a.Damage = 75
	d.Damage = 50
}

func (c *Calculator) CalcDamage() {

}

func (ci *CalculatorInput) DamageBoost() int {
	boost := ci.Co.DamageBoost(ci.Unit)
	if ci.Player.CoPowerOn != "N" {
		boost += 10
	}
	return boost
}

func NewCalculator(a unitmodels.Unit, d unitmodels.Unit) Calculator {
	return Calculator{
		Attacker: NewCalculatorInput(a),
		Defender: NewCalculatorInput(d),
	}
}

func NewCalculatorInput(u unitmodels.Unit) *CalculatorInput {
	return &CalculatorInput{
		Damage:  0,
		Defense: 0,
		Unit:    u,
		Tile:    u.GetTile(),
		Player:  u.GetPlayer(),
		Co:      u.GetPlayer().Co,
	}
}
