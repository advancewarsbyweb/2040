package models

import (
	"testing"

	"github.com/awbw/2040/models"
	conames "github.com/awbw/2040/models/internal/cos/names"
	unitmodels "github.com/awbw/2040/models/internal/units"
	unitnames "github.com/awbw/2040/models/internal/units/names"
)

var (
	samiTest   models.ICo
	samiPlayer models.Player

	Footsoldiers []unitnames.UnitName
)

func init() {
	samiTest = NewCo(conames.Sami)
	samiPlayer = models.Player{}
	Footsoldiers = []unitnames.UnitName{unitnames.Infantry, unitnames.Mech}
}

func TestDamageBoostFootsoldiers_Sami(t *testing.T) {
	for _, unitName := range Footsoldiers {
		u := unitmodels.CreateUnitHelper(unitName).SetPlayer(&samiPlayer)
		boost := samiTest.DamageBoost(u)
		want := 30

		if boost != want {
			t.Fatalf("Sami's %s damage boost is wrong. Got (%d), want (%d)", unitName, boost, want)
		}
	}
}

func TestDamageBoostDirectUnits_Sami(t *testing.T) {
	for _, unitName := range unitmodels.DirectUnits {
		u := unitmodels.CreateUnitHelper(unitName).SetPlayer(&samiPlayer)
		boost := samiTest.DamageBoost(u)
		want := -10

		if boost != want {
			t.Fatalf("Sami's %s damage boost is wrong. Got (%d), want (%d)", unitName, boost, want)
		}
	}
}

func TestDamageBoostIndirectUnits_Sami(t *testing.T) {
	for _, unitName := range unitmodels.IndirectUnits {
		u := unitmodels.CreateUnitHelper(unitName).SetPlayer(&samiPlayer)
		boost := samiTest.DamageBoost(u)
		want := 0

		if boost != want {
			t.Fatalf("Sami's %s damage boost is wrong. Got (%d), want (%d)", unitName, boost, want)
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
			t.Fatalf("Sami's %s capture boost is wrong at %d hp . Got (%d), want (%d)", unitName, hp, boost, want)
		}
	}
}

// Number with decimals. e.g: 3hp * 1.5 = 4.5, floored is 4 (+1 boost)
func TestCaptureBoostDecimalTotal_Sami(t *testing.T) {
	for _, unitName := range Footsoldiers {
		hp := 3
		u := unitmodels.CreateUnitHelper(unitName).SetPlayer(&samiPlayer).SetHp(float64(hp))
		boost := samiTest.CaptureBoost(u)
		want := 1

		if boost != want {
			t.Fatalf("Sami's %s capture boost is wrong at %d hp. Got (%d), want (%d)", unitName, hp, boost, want)
		}
	}
}

// Make sure any HP can capture a full HP building
func TestCaptureBoostScop_Sami(t *testing.T) {
	hp := 10
	for _, unitName := range Footsoldiers {
		samiPlayer.CoPowerOn = "S"
		u := unitmodels.CreateUnitHelper(unitName).SetPlayer(&samiPlayer).SetHp(float64(hp))
		boost := samiTest.CaptureBoost(u)
		want := 20

		if boost < want {
			t.Fatalf("Sami's %s capture boost is wrong during SCOP at %d hp. Got (%d), want (%d)", unitName, hp, boost, want)
		}

		hp -= 6
	}
}
