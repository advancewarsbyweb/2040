package models

import (
	"testing"

	unitnames "github.com/awbw/2040/types/units/names"
)

func TestIndirectFire(t *testing.T) {
	a := CreateUnitHelper(unitnames.Artillery).SetPos(1, 1)
	d := CreateUnitHelper(unitnames.Mech).SetPos(2, 1)
	a.Fire(a, d)
	if d.Hp() != 10 {
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
	if d.Hp() == 10 {
		t.Fatalf("Indirect unit fired but dealt no damage")
	}
}

func TestIndirectFireHasMoved(t *testing.T) {
	a := CreateUnitHelper(unitnames.Artillery).SetMoved(1).SetPos(1, 1)
	d := CreateUnitHelper(unitnames.Mech).SetPos(2, 2)
	a.Fire(a, d)
	if d.Hp() < 10 {
		t.Fatalf("Indirect unit fired after moving")
	}
}

func TestIndirectFireWithAmmo(t *testing.T) {
	a := CreateUnitHelper(unitnames.Artillery).SetPos(1, 1)
	d := CreateUnitHelper(unitnames.Mech).SetPos(2, 2)
	want := a.Ammo() - 1
	a.Fire(a, d)
	if a.Ammo() != want {
		t.Fatalf("Wrong ammo count. Got (%d), want (%d)", a.Ammo(), want)
	}
}

func TestIndirectFireNoAmmo(t *testing.T) {
	a := CreateUnitHelper(unitnames.Artillery).SetAmmo(0).SetPos(1, 1)
	d := CreateUnitHelper(unitnames.Mech).SetPos(2, 2)
	err := a.Fire(a, d)
	if err == nil || d.Hp() < 10 {
		t.Fatalf("Indirect unit fired with no ammo")
	}
}
