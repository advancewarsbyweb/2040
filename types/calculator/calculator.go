package calculator

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types/cos"
	unitmodels "github.com/awbw/2040/types/units"
	"golang.org/x/exp/slices"
)

type damageRange struct {
	Hp  int
	Min int
	Max int
}

type Calculator struct {
	Attacker *CalculatorInput
	Defender *CalculatorInput
}

type CalculatorInput struct {
	// Damage and Defense can vary based on the CO, Power, etc
	Damage int
	// There may be multiple possible HP a defender can counter attack from so we need
	// to consider all possibilities
	// e.g: { 5, 27, 32 } = 5HP, min 27, max 32
	Preview   []damageRange
	MaxDamage int
	MinDamage int
	Unit      models.IUnit
	Tile      *models.Tile
	Player    *models.Player
	Co        models.ICo
}

func (c *Calculator) CalcPreviewResult() error {
	preview, err := c.CalcDamagePreview(c.Attacker, c.Defender)
	if err != nil {
		return err
	}
	c.Attacker.Preview = preview
	if slices.Contains(unitmodels.IndirectUnits, c.Attacker.Unit.GetName()) {
		return nil
	}

	preview, err = c.CalcDamagePreview(c.Defender, c.Attacker)
	if err != nil {
		return err
	}
	c.Defender.Preview = preview
	return nil
}

func (c *Calculator) CalcResult() error {
	damage, err := c.CalcDamage(c.Attacker, c.Defender)
	if err != nil {
		return err
	}
	c.Attacker.Damage = damage
	if slices.Contains(unitmodels.IndirectUnits, c.Attacker.Unit.GetName()) {
		return nil
	}

	damage, err = c.CalcDamage(c.Defender, c.Attacker)
	if err != nil {
		return err
	}
	c.Defender.Damage = damage
	return nil
}

// For the Damage Preview there can be multiple damage ranges
func (c *Calculator) CalcDamagePreview(a *CalculatorInput, d *CalculatorInput) ([]damageRange, error) {
	baseDamage := c.BaseDamage(a.Unit, d.Unit)
	if baseDamage <= 0 {
		return nil, errors.New(fmt.Sprintf("%s can not fire at %s", a.Unit.GetName(), d.Unit.GetName()))
	}

	attackValue := 100 + a.Co.DamageBoost(a.Unit)
	minHp := math.Ceil((a.Unit.GetHp()*10 - float64(d.Preview[0].Max)) / 10)
	maxHp := math.Ceil((a.Unit.GetHp()*10 - float64(d.Preview[0].Min)) / 10)
	if d.Preview[0].Min > 100 {
		maxHp = 0
	}
	if d.Preview[0].Max > 100 {
		minHp = 0
	}
	minLuck, maxLuck := c.Attacker.Co.LuckRange()

	dDef := 100 + d.Co.DefenseModifier()
	terrainStars := d.Tile.Defense
	dHp := math.Ceil(d.Unit.GetHp())

	minHpMinDamage := c.Formula(baseDamage, attackValue, minLuck, int(minHp), dDef, terrainStars, int(dHp))
	minHpMaxDamage := c.Formula(baseDamage, attackValue, maxLuck, int(minHp), dDef, terrainStars, int(dHp))

	var damageValues []damageRange
	damageValues = []damageRange{
		{Hp: int(minHp), Min: minHpMinDamage, Max: minHpMaxDamage},
	}

	if minHp != maxHp {
		maxHpMinDamage := c.Formula(baseDamage, attackValue, minLuck, int(maxHp), dDef, terrainStars, int(dHp))
		maxHpMaxDamage := c.Formula(baseDamage, attackValue, maxLuck, int(maxHp), dDef, terrainStars, int(dHp))

		damageValues = append(damageValues, damageRange{Hp: int(maxHp), Min: maxHpMinDamage, Max: maxHpMaxDamage})
	}

	return damageValues, nil
}

func (c *Calculator) CalcDamage(a *CalculatorInput, d *CalculatorInput) (int, error) {
	baseDamage := c.BaseDamage(a.Unit, d.Unit)
	if baseDamage <= 0 {
		return 0, errors.New(fmt.Sprintf("%s can not fire at %s", a.Unit.GetName(), d.Unit.GetName()))
	}

	attackValue := 100 + a.Co.DamageBoost(a.Unit)
	aHp := math.Ceil((a.Unit.GetHp()*10 - float64(d.Damage)) / 10)
	minLuck, maxLuck := c.Attacker.Co.LuckRange()
	luckValue := rand.Intn(maxLuck-minLuck) + minLuck

	dDef := 100 + d.Co.DefenseModifier()
	terrainStars := d.Tile.Defense
	dHp := math.Ceil(d.Unit.GetHp())

	damage := c.Formula(baseDamage, attackValue, luckValue, int(aHp), dDef, terrainStars, int(dHp))

	return damage, nil
}

func (c *Calculator) Formula(baseDamage int, attackValue int, luckValue int, aHp int, dDef int, terrainStars int, dHp int) int {
	damage := (baseDamage*attackValue/100 + luckValue) * aHp / 10 * (200 - (dDef + terrainStars*dHp)) / 100
	// Ceil to nearest 0.05
	ceiled := math.Ceil(float64(damage)/0.05) * 0.05
	// Floor to nearest int
	return int(math.Floor(ceiled))
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
		Damage:  0,
		Preview: []damageRange{{}, {}},
		Unit:    u,
		Tile:    u.GetTile(),
		Player:  u.GetPlayer(),
		Co:      cos.NewCo(u.GetPlayer().Co.Name),
	}
}
