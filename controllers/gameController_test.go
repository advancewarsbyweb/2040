package controllers

import (
	"encoding/json"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/factories"
	"github.com/awbw/2040/types"
	"github.com/gin-gonic/gin"
)

func init() {
	db.NewTestDatabase()
}

func TestGetGame(t *testing.T) {
	g := factories.Game.Create()
	gameId, _ := db.GameRepo.CreateGame(g)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: strconv.Itoa(gameId),
		},
	}
	Game.Get(c)

	var res types.Game
	json.Unmarshal(w.Body.Bytes(), &res)

	if gameId != res.ID {
		t.Fatalf("Wrong game returned from controller. Got (%d), want (%d)", res.ID, gameId)
	}
}
