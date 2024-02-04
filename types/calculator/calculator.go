package calculator

import (
	"github.com/awbw/2040/models"
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
	Unit    models.IUnit
	Tile    *models.Tile
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

func NewCalculator(a models.IUnit, d models.IUnit) Calculator {
	return Calculator{
		Attacker: NewCalculatorInput(a),
		Defender: NewCalculatorInput(d),
	}
}

func NewCalculatorInput(u models.IUnit) *CalculatorInput {
	return &CalculatorInput{
		Damage:  0,
		Defense: 0,
		Unit:    u,
		Tile:    u.GetTile(),
		Player:  u.GetPlayer(),
		Co:      u.GetPlayer().Co,
	}
}
