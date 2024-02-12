package db

import (
	"testing"

	gamecolumns "github.com/awbw/2040/models/columnNames/game"
)

func init() {
	NewTestDatabase()
}

func TestCreateGame(t *testing.T) {
	g := Game.Create().Build()
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
}

func TestUpdateGame(t *testing.T) {
	g := Game.Create().Build()
	gameId, _ := GameRepo.CreateGame(g)

	newName := "Test Update Game"
	newFunds := 5000
	updatedFields := map[string]interface{}{
		gamecolumns.Name:  newName,
		gamecolumns.Funds: newFunds,
	}
	updatedFields, err := GameRepo.UpdateGame(gameId, updatedFields)
	if err != nil {
		t.Fatalf("Could not update Game (%d): %s", gameId, err.Error())
	}

	updatedGame, err := GameRepo.FindGame(gameId)
	if err != nil {
		t.Fatalf("Could not find Game (%d): %s", gameId, err.Error())
	}

	if updatedGame.Name != newName || updatedGame.Funds != newFunds {
		t.Fatalf("Game not updated properly. Got (%s, %d), want (%s, %d)", updatedGame.Name, updatedGame.Funds, newName, newFunds)
	}
}

func TestDeleteGame(t *testing.T) {
	g := Game.Create().Build()
	gameId, _ := GameRepo.CreateGame(g)

	err := GameRepo.DeleteGame(gameId)
	if err != nil {
		t.Fatalf("Could not delete Game: %s", err.Error())
	}
}
