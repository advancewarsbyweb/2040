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

	t.Logf("Created User (%d)", p.ID)
}
