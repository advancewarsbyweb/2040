package baseunits

import "testing"

func TestFireWithAmmo(t *testing.T) {
	att := NewMech()
	def := NewArtillery()
	att.Fire(def)
	if att.GetAmmo() != 2 {
		t.Fatalf("Wrong ammo count. Got (%d), want (%d)", att.GetAmmo(), 2)
	}
}

func TestFireWithoutAmmo(t *testing.T) {
	att := NewMech()
	def := NewArtillery()
	att.SetAmmo(0)
	att.Fire(def)
	if att.GetAmmo() < 0 {
		t.Fatalf("Wrong ammo count. Got (%d), want (%d)", att.GetAmmo(), 0)
	}
}

func TestDirectFire(t *testing.T) {
	att := NewMech()
	def := NewInfantry()
	att.SetPos(1, 1)
	def.SetPos(2, 1)
}
