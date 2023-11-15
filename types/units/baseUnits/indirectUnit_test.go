package baseunits

import (
	"testing"
)

func TestIndirectFire(t *testing.T) {
	att := NewArtillery()
	def := NewInfantry()
	att.SetPos(1, 1)
	def.SetPos(2, 3)
	att.Fire(def)
	if want := NewArtillery().GetAmmo() - 1; att.GetAmmo() != want {
		t.Fatalf("Wrong ammo count. Got (%d), want (%d)", att.GetAmmo(), want)
	}
}

func TestIndirectFireHasMoved(t *testing.T) {
	att := NewArtillery()
	def := NewInfantry()
	att.SetMoved(1)
	err := att.Fire(def)
	if err.Error() != AttackerAlreadyMoved {
		t.Fatalf("Attacker fired despite having already moved")
	}
}
