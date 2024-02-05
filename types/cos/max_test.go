package models

import (
	"testing"

	"github.com/awbw/2040/models"
	conames "github.com/awbw/2040/models/internal/cos/names"
	unitmodels "github.com/awbw/2040/models/internal/units"
)

var (
	maxTest   models.ICo
	maxPlayer models.Player
)

func init() {
	maxTest = NewCo(conames.Max)
	maxPlayer = models.Player{}
}

func TestRangeBoost_Max(t *testing.T) {
	for _, unitName := range unitmodels.IndirectUnits {
		u := unitmodels.CreateUnitHelper(unitName).SetPlayer(&maxPlayer)
		boost := maxTest.RangeBoost(u)
		want := -1

		if boost != want {
			t.Fatalf("Max's %s (Indirect Unit) range boost is wrong. Got (%d), want (%d)", unitName, boost, want)
		}
	}
}

func TestIndirectUnitDamage_Max(t *testing.T) {
	for _, unitName := range unitmodels.IndirectUnits {
		u := unitmodels.CreateUnitHelper(unitName).SetPlayer(&maxPlayer)
		boost := maxTest.DamageBoost(u)
		want := -10

		if boost != want {
			t.Fatalf("Max's %s (Indirect Unit) damage boost is wrong. Got (%d), want (%d)", unitName, boost, want)
		}
	}
}

func TestDirectUnitDamage_Max(t *testing.T) {
	for _, unitName := range unitmodels.DirectUnits {
		u := unitmodels.CreateUnitHelper(unitName).SetPlayer(&maxPlayer)
		boost := maxTest.DamageBoost(u)
		want := 20

		if boost != want {
			t.Fatalf("Max's %s (Direct Unit) damage boost is wrong. Got (%d), want (%d)", unitName, boost, want)
		}
	}
}
