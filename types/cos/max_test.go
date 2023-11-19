package cos

import (
	"testing"

	conames "github.com/awbw/2040/types/cos/names"
	unittypes "github.com/awbw/2040/types/units/baseUnits"
	unitnames "github.com/awbw/2040/types/units/names"
)

func TestRangeBoost(t *testing.T) {
	u := unittypes.CreateUnit(unitnames.Artillery)

	max := NewCo(conames.Max)

	boost := max.RangeBoost(&u)
	want := -1

	if boost != want {
		t.Fatalf("Max's indirect unit boost should be -1. Got (%d), want (%d)", boost, want)
	}
}
