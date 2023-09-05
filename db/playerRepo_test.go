package db

import (
	"testing"

	"github.com/awbw/2040/factories"
)

func init() {
	NewTestDatabase()
}

func TestCreatePlayer(t *testing.T) {
	p := factories.Player.Create()
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
	u := factories.User.Create()
	uID, _ := UserRepo.CreateUser(u)

	p := factories.Player.Create()
	p.UserID = uID
	playerId, _ := PlayerRepo.CreatePlayer(p)

	r, err := PlayerRepo.FindPlayerUser(playerId)
	if err != nil {
		t.Fatalf("Could not find PlayerUser (%d): %s", playerId, err.Error())
	}

	if r.User.ID != uID || r.ID != playerId {
		t.Fatalf("IDs of player found do not match given User. Got (Player: %d, User: %d), want (Player: %d, User: %d)", r.ID, r.User.ID, playerId, uID)
	}
}
