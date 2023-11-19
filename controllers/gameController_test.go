package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/types"
	"github.com/gin-gonic/gin"
)

func init() {
	db.NewTestDatabase()
}

func TestCreateGame(t *testing.T) {
	g := db.Game.Create().Build()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	jsonBody, _ := json.Marshal(g)

	c.Request = &http.Request{
		Method: "POST",
		Body:   io.NopCloser(bytes.NewBuffer(jsonBody)),
		Header: http.Header{},
	}
	c.Request.Header.Set("Content-Type", "application/json")

	Game.Create(c)

	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("Wrong status code returned. Got (%d), want (%d)", w.Result().StatusCode, http.StatusOK)
	}

	var res types.Game
	json.Unmarshal(w.Body.Bytes(), &res)

	if g.Name != res.Name {
		t.Fatalf("Game created does not have given name. Got (%s), want (%s)", res.Name, g.Name)
	}
}

func TestGetGame(t *testing.T) {
	g := db.Game.Create().BuildInsert()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{
		{
			Key:   "id",
			Value: strconv.Itoa(g.ID),
		},
	}

	Game.Get(c)

	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("Wrong status code returned. Got (%d), want (%d)", w.Result().StatusCode, http.StatusOK)
	}

	var res types.Game
	json.Unmarshal(w.Body.Bytes(), &res)

	if g.ID != res.ID {
		t.Fatalf("Wrong game returned from controller. Got (%d), want (%d)", res.ID, g.ID)
	}
}
