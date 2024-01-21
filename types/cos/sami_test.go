package cos

import (
	"testing"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/models"
	unitmodels "github.com/awbw/2040/models/units"
	conames "github.com/awbw/2040/types/cos/names"
	unitnames "github.com/awbw/2040/types/units/names"
)

var samiTest Co
var samiPlayer models.Player

var Footsoldiers []unitnames.UnitName

func init() {
	samiTest = NewCo(conames.Sami)
	samiPlayer = db.PlayerFactory.Create().Build()
	Footsoldiers = []unitnames.UnitName{unitnames.Infantry, unitnames.Mech}
}

func TestDamageBoostFootsoldiers_Sami(t *testing.T) {
	for _, unitName := range Footsoldiers {
		u := unitmodels.CreateUnitHelper(unitName).SetPlayer(&samiPlayer)
		boost := samiTest.DamageBoost(u)
		want := 30

		if boost != want {
			t.Fatalf("Sami's %s damage boost should be 30. Got (%d), want (%d)", unitName, boost, want)
		}
	}
}

func TestDamageBoostDirectUnits_Sami(t *testing.T) {
	for _, unitName := range unitmodels.DirectUnits {
		u := unitmodels.CreateUnitHelper(unitName).SetPlayer(&samiPlayer)
		boost := samiTest.DamageBoost(u)
		want := -10

		if boost != want {
			t.Fatalf("Sami's %s damage boost should be -10. Got (%d), want (%d)", unitName, boost, want)
		}
	}
}

func TestCaptureBoost_Sami(t *testing.T) {
	for _, unitName := range Footsoldiers {
		hp := 10
		u := unitmodels.CreateUnitHelper(unitName).SetPlayer(&samiPlayer).SetHp(float64(hp))
		boost := samiTest.CaptureBoost(u)
		want := 5

		if boost != want {
			t.Fatalf("Sami's %s capture boost should be +5 at %d hp . Got (%d), want (%d)", unitName, hp, boost, want)
		}
	}
}

// Number with decimals. e.g: 3hp * 1.5 = 4.5, floored is 4
func TestCaptureBoostDecimalTotal_Sami(t *testing.T) {
	for _, unitName := range Footsoldiers {
		hp := 3
		u := unitmodels.CreateUnitHelper(unitName).SetPlayer(&samiPlayer).SetHp(float64(hp))
		boost := samiTest.CaptureBoost(u)
		want := 1

		if boost != want {
			t.Fatalf("Sami's %s capture boost should be +1 at %d hp. Got (%d), want (%d)", unitName, hp, boost, want)
		}
	}
}
