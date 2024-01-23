package cos

import (
	"testing"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/models"
	unitmodels "github.com/awbw/2040/models/units"
	conames "github.com/awbw/2040/types/cos/names"
)

var (
	maxTest   Co
	maxPlayer models.Player
)

func init() {
	maxTest = NewCo(conames.Max)
	maxPlayer = db.PlayerFactory.Create().Build()
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
