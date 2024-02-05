package unitmodels

import (
	"testing"

	unitnames "github.com/awbw/2040/types/units/names"
)

func TestDirectFire(t *testing.T) {
	a := CreateUnitHelper(unitnames.Mech).SetPos(1, 1)
	d := CreateUnitHelper(unitnames.Artillery).SetPos(2, 2)
	a.Fire(a, d)
	if d.GetHp() != 10.0 {
		t.Fatalf("Direct unit fired at defender out of range")
	}
}

func TestDirectFireDamage(t *testing.T) {
	a := CreateUnitHelper(unitnames.Mech).SetPos(1, 1)
	d := CreateUnitHelper(unitnames.Artillery).SetPos(2, 1)
	err := a.Fire(a, d)
	if err == nil && d.GetHp() == 10 {
		t.Fatalf("Direct unit fired but dealt no damage")
	}
}

func TestFireWithAmmo(t *testing.T) {
	a := CreateUnitHelper(unitnames.Mech).SetPos(1, 1)
	d := CreateUnitHelper(unitnames.Artillery).SetPos(2, 1)
	want := a.GetAmmo() - 1
	err := a.Fire(a, d)
	if err != nil {
		t.Fatalf("Error happened: %s", err.Error())
	}
	if want != a.GetAmmo() {
		t.Fatalf("Wrong ammo count. Got (%d), want (%d)", a.GetAmmo(), want)
	}
}

func TestFireNoAmmo(t *testing.T) {
	a := CreateUnitHelper(unitnames.Mech).SetAmmo(0).SetPos(1, 1)
	d := CreateUnitHelper(unitnames.Artillery).SetPos(2, 1)
	a.Fire(a, d)
	if a.GetAmmo() < 0 {
		t.Fatalf("Wrong ammo count. Got (%d), want (%d)", a.GetAmmo(), 0)
	}
}
