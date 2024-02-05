package unitmodels

import (
	"testing"

	unitnames "github.com/awbw/2040/models/internal/units/names"
)

func TestIndirectFire(t *testing.T) {
	a := CreateUnitHelper(unitnames.Artillery).SetPos(1, 1)
	d := CreateUnitHelper(unitnames.Mech).SetPos(2, 1)
	a.Fire(a, d)
	if d.GetHp() != 10 {
		t.Fatalf("Indirect unit fired and damaged defender while adjacent to it")
	}
}

func TestIndirectFireDamage(t *testing.T) {
	a := CreateUnitHelper(unitnames.Artillery).SetPos(1, 1)
	d := CreateUnitHelper(unitnames.Mech).SetPos(2, 2)
	err := a.Fire(a, d)
	if err != nil {
		t.Fatal(err)
	}
	if d.GetHp() == 10 {
		t.Fatalf("Indirect unit fired but dealt no damage")
	}
}

func TestIndirectFireHasMoved(t *testing.T) {
	a := CreateUnitHelper(unitnames.Artillery).SetMoved(1).SetPos(1, 1)
	d := CreateUnitHelper(unitnames.Mech).SetPos(2, 2)
	a.Fire(a, d)
	if d.GetHp() < 10 {
		t.Fatalf("Indirect unit fired after moving")
	}
}

func TestIndirectFireWithAmmo(t *testing.T) {
	a := CreateUnitHelper(unitnames.Artillery).SetPos(1, 1)
	d := CreateUnitHelper(unitnames.Mech).SetPos(2, 2)
	want := a.GetAmmo() - 1
	a.Fire(a, d)
	if a.GetAmmo() != want {
		t.Fatalf("Wrong ammo count. Got (%d), want (%d)", a.GetAmmo(), want)
	}
}

func TestIndirectFireNoAmmo(t *testing.T) {
	a := CreateUnitHelper(unitnames.Artillery).SetAmmo(0).SetPos(1, 1)
	d := CreateUnitHelper(unitnames.Mech).SetPos(2, 2)
	err := a.Fire(a, d)
	if err == nil || d.GetHp() < 10 {
		t.Fatalf("Indirect unit fired with no ammo")
	}
}
