package cos

import (
	"testing"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/types"
	conames "github.com/awbw/2040/types/cos/names"
	unittypes "github.com/awbw/2040/types/units/baseUnits"
	unitnames "github.com/awbw/2040/types/units/names"
)

var maxTest Co
var maxPlayer types.Player

func init() {
	maxTest = NewCo(conames.Max)
	maxPlayer = types.NewPlayer(db.PlayerFactory.Create().Build())
}

func TestRangeBoost(t *testing.T) {
	u := unittypes.CreateUnit(unitnames.Artillery)

	boost := maxTest.RangeBoost(u)
	want := -1

	if boost != want {
		t.Fatalf("Max's indirect unit range boost should be -1. Got (%d), want (%d)", boost, want)
	}
}

func TestIndirectUnitDamage(t *testing.T) {
	u := unittypes.CreateUnit(unitnames.Artillery)

	boost := maxTest.DamageBoost(u)
	want := -10

	if boost != want {
		t.Fatalf("Max's indirect unit damage boost should be -10. Got (%d), want (%d)", boost, want)
	}
}

func TestDirectUnitDamage(t *testing.T) {
	u := unittypes.CreateUnit(unitnames.Tank).SetPlayer(&maxPlayer)

	boost := maxTest.DamageBoost(u)
	want := 20

	if boost != want {
		t.Fatalf("Max's direct unit damage boost should be 20. Got (%d), want (%d)", boost, want)
	}
}
