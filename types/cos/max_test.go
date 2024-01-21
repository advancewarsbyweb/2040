package cos

import (
	"testing"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/models"
	unitmodels "github.com/awbw/2040/models/units"
	conames "github.com/awbw/2040/types/cos/names"
	unitnames "github.com/awbw/2040/types/units/names"
)

var maxTest Co
var maxPlayer models.Player

func init() {
	maxTest = NewCo(conames.Max)
	maxPlayer = db.PlayerFactory.Create().Build()
}

func TestRangeBoost(t *testing.T) {
	u := unitmodels.CreateUnitHelper(unitnames.Artillery).SetPlayer(&maxPlayer)

	boost := maxTest.RangeBoost(u)
	want := -1

	if boost != want {
		t.Fatalf("Max's indirect unit range boost should be -1. Got (%d), want (%d)", boost, want)
	}
}

func TestIndirectUnitDamage(t *testing.T) {
	u := unitmodels.CreateUnitHelper(unitnames.Artillery).SetPlayer(&maxPlayer)

	boost := maxTest.DamageBoost(u)
	want := -10

	if boost != want {
		t.Fatalf("Max's indirect unit damage boost should be -10. Got (%d), want (%d)", boost, want)
	}
}

func TestDirectUnitDamage(t *testing.T) {
	u := unitmodels.CreateUnitHelper(unitnames.Tank).SetPlayer(&maxPlayer)

	boost := maxTest.DamageBoost(u)
	want := 20

	if boost != want {
		t.Fatalf("Max's direct unit damage boost should be 20. Got (%d), want (%d)", boost, want)
	}
}
