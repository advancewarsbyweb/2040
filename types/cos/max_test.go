package cos

import (
	"testing"

	baseunits "github.com/awbw/2040/types/units/baseUnits"
)

func TestRangeBoost(t *testing.T) {
	u := baseunits.NewArtillery()

	max := NewMax()

	boost := max.RangeBoost(u)
}
