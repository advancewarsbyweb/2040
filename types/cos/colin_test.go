package models

import (
	"testing"

	"github.com/awbw/2040/models"
	unitmodels "github.com/awbw/2040/models/internal/units"
	unitnames "github.com/awbw/2040/models/internal/units/names"
)

var (
	colinTest   models.ICo
	colinPlayer models.Player
)

func init() {
	colinTest = NewColin()
	colinPlayer = models.Player{}
}

func TestDamageBoost_Colin(t *testing.T) {
	colinPlayer.CoPowerOn = "N"
	u := unitmodels.CreateUnitHelper(unitnames.Infantry).SetPlayer(&colinPlayer)
	boost := colinTest.DamageBoost(u)
	want := -10

	if boost != want {
		t.Fatalf("Colin's damage boost is wrong. Got (%d), want (%d)", boost, want)
	}
}

func TestDamageBoostScop_Colin(t *testing.T) {
	colinPlayer.Funds = 10500
	colinPlayer.CoPowerOn = "S"
	u := unitmodels.CreateUnitHelper(unitnames.Infantry).SetPlayer(&colinPlayer)
	boost := colinTest.DamageBoost(u)
	want := 30

	if boost != want {
		t.Fatalf("Colin's damage boost during SCOP is wrong. Got (%d), want (%d)", boost, want)
	}
}

func TestCostModifier_Colin(t *testing.T) {
	u := unitmodels.CreateUnitHelper(unitnames.Infantry).SetPlayer(&colinPlayer)
	cost := int(float64(u.GetCost()) * colinTest.CostModifier())
	want := 800

	if cost != want {
		t.Fatalf("Colin's cost modifier for an Infantry is wrong. Got (%d), want (%d)", cost, want)
	}
}
