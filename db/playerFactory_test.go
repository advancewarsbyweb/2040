package db

import (
	"testing"
)

func TestSetGame(t *testing.T) {
	g := GameFactory.Build()
	p := PlayerFactory.Create().SetGame(&g).Build()

	if p.Game.Name != g.Name {
		t.Fatalf("Game was not set to the correct game. Got name (%s), want (%s)", p.Game.Name, g.Name)
	}
}

func TestSetUser(t *testing.T) {
	u := UserFactory.Build()
	p := PlayerFactory.Create().SetUser(&u).Build()

	if p.User.Username != u.Username {
		t.Fatalf("User was not set to the correct game. Got username (%s), want (%s)", p.User.Username, u.Username)
	}
}
