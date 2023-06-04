package db

import (
	"testing"

	"github.com/awbw/2040/factories"
)

func init() {
	NewTestDatabase()
}

func TestCreateGame(t *testing.T) {
	g := factories.Game.Create()
	gameId, err := GameRepo.CreateGame(g)

	if err != nil {
		t.Errorf("Could not create Game (%d): %s", g.ID, err.Error())
	}

	newGame, err := GameRepo.FindGame(gameId)

	if err != nil {
		t.Errorf("Could not find Game (%d): %s", gameId, err.Error())
		return
	}

	if newGame.ID != gameId {
		t.Errorf("ID of Game created does not match given Game. Got %d, want %d", newGame.ID, gameId)
		return
	}
	t.Logf("Created Game %s", newGame.Name)
}
