package db

import (
	"testing"
)

func init() {
	NewTestDatabase()
}

func TestCreatePlayer(t *testing.T) {
	p := Player.Create().Build()
	playerId, err := PlayerRepo.CreatePlayer(p)

	if err != nil {
		t.Fatalf("Could not create Player: %s", err.Error())
	}

	newPlayer, err := PlayerRepo.FindPlayer(playerId)
	if err != nil {
		t.Fatalf("Could not find Player (%d): %s", playerId, err.Error())
	}

	if newPlayer.ID != playerId {
		t.Fatalf("ID of user created does not match given User. Got (%d), want (%d)", newPlayer.ID, playerId)
	}
}

func TestFindPlayerUser(t *testing.T) {
	p := Player.Create().CreateUser().BuildInsert()

	r, err := PlayerRepo.FindPlayerUser(p.ID)
	if err != nil {
		t.Fatalf("Could not find PlayerUser (%d): %s", p.ID, err.Error())
	}

	if r.User.ID != p.User.ID || r.ID != p.ID {
		t.Fatalf("IDs of player found do not match given User. Got (Player: %d, User: %d), want (Player: %d, User: %d)", r.ID, r.User.ID, p.ID, p.User.ID)
	}
}
