package calculator

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	"github.com/awbw/2040/models"
)

type Calculator struct {
	Attacker *CalculatorInput
	Defender *CalculatorInput
}

type CalculatorInput struct {
	// Damage and Defense can vary based on the CO, Power, etc
	Damage    int
	MaxDamage int
	MinDamage int
	Unit      models.IUnit
	Tile      *models.Tile
	Player    *models.Player
	Co        models.ICo
}

func (c *Calculator) CalcPreview() error {
	return nil
}

func (c *Calculator) CalcResult() error {
	minLuck, maxLuck := c.Attacker.Co.LuckRange()
	luckValue := rand.Intn(maxLuck-minLuck) + minLuck
	damage, err := c.CalcDamage(c.Attacker, c.Defender, luckValue)
	if err != nil {
		return err
	}
	c.Attacker.Damage = damage

	minLuck, maxLuck = c.Defender.Co.LuckRange()
	luckValue = rand.Intn(maxLuck-minLuck) + minLuck
	damage, err = c.CalcDamage(c.Defender, c.Attacker, luckValue)
	if err != nil {
		return err
	}
	c.Defender.Damage = damage
	return nil
}

func (c *Calculator) CalcDamagePreview(a *CalculatorInput, d *CalculatorInput, minLuck int, maxLuck int) (int, int, error) {
	baseDamage := c.BaseDamage(a.Unit, d.Unit)
	if baseDamage <= 0 {
		return 0, 0, errors.New(fmt.Sprintf("%s can not fire at %s", a.Unit.GetName(), d.Unit.GetName()))
	}

	attackValue := 100 + a.Co.DamageBoost(a.Unit)
	// Avoid divisions by 0
	aMinHp := math.Ceil((a.Unit.GetHp()*10 - float64(d.MaxDamage)) / 10)
	aMaxHp := math.Ceil((a.Unit.GetHp()*10 - float64(d.MinDamage)) / 10)

	dDef := 100 + d.Co.DefenseModifier()
	terrainStars := d.Tile.Defense
	dHp := math.Ceil(d.Unit.GetHp())

	minDamage := (baseDamage*attackValue/100 + minLuck) * int(aMinHp) / 10 * (200 - (dDef + terrainStars*int(dHp))) / 100
	maxDamage := (baseDamage*attackValue/100 + maxLuck) * int(aMaxHp) / 10 * (200 - (dDef + terrainStars*int(dHp))) / 100

	return minDamage, maxDamage, nil
}

func (c *Calculator) CalcDamage(a *CalculatorInput, d *CalculatorInput, luckValue int) (int, error) {
	baseDamage := c.BaseDamage(a.Unit, d.Unit)
	if baseDamage <= 0 {
		return 0, errors.New(fmt.Sprintf("%s can not fire at %s", a.Unit.GetName(), d.Unit.GetName()))
	}

	attackValue := 100 + a.Co.DamageBoost(a.Unit)
	aHp := math.Ceil((a.Unit.GetHp()*10 - float64(d.Damage)) / 10)

	dDef := 100 + d.Co.DefenseModifier()
	terrainStars := d.Tile.Defense
	dHp := math.Ceil(d.Unit.GetHp())

	damage := (baseDamage*attackValue/100 + luckValue) * int(aHp) / 10 * (200 - (dDef + terrainStars*int(dHp))) / 100

	return damage, nil
}

func (c *Calculator) BaseDamage(a models.IUnit, d models.IUnit) int {
	baseDamage := DamageMatrix.FindPrimary(a.GetName(), d.GetName())

	if baseDamage <= 0 {
		baseDamage = DamageMatrix.FindSecondary(a.GetUnit().Name, d.GetUnit().Name)
	}
	return baseDamage
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
		Damage:    0,
		MinDamage: 0,
		MaxDamage: 0,
		Unit:      u,
		Tile:      u.GetTile(),
		Player:    u.GetPlayer(),
		Co:        u.GetPlayer().Co,
	}
}
