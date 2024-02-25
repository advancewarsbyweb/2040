package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/awbw/2040/db"
	httputils "github.com/awbw/2040/utils/http"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	db.NewTestDatabase()
}

func TestValidatePlayerSkipValidation(t *testing.T) {
	assert := assert.New(t)
	f := db.GameFactory.Create()
	f.Game.SetEndDate(time.Now())

	g := f.BuildInsert()
	u := db.UserFactory.Create().BuildInsert()
	p := db.PlayerFactory.Create().SetGame(g).SetUser(&u).BuildInsert()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Method: "GET",
		Header: http.Header{},
	}
	c.Params = []gin.Param{
		{
			Key:   "playerId",
			Value: strconv.Itoa(p.ID),
		},
	}
	ValidatePlayer(c)

	_, exists := c.Get("ValidationSkip")
	assert.Equal(exists, true, "Skip validation is not set")
}

func TestValidatePlayerQueryParam(t *testing.T) {
	assert := assert.New(t)
	u := db.UserFactory.Create().BuildInsert()
	p := db.PlayerFactory.CreateRelations().SetUser(&u).BuildInsert()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Method: "GET",
		Header: http.Header{},
	}
	c.Params = []gin.Param{
		{
			Key:   "playerId",
			Value: strconv.Itoa(p.ID),
		},
	}
	c.Set("User", u)
	ValidatePlayer(c)

	if !httputils.StatusOK(w) {
		t.Fatalf("Wrong status code returned. Got (%d), want (%d). %v", w.Result().StatusCode, http.StatusOK, httputils.FormatError(w))
	}

	_, exists := c.Get("PlayerUser")
	assert.Equal(exists, true, "PlayerUser is not set properly")
}

func TestValidatePlayerWithBody(t *testing.T) {
	u := db.UserFactory.Create().BuildInsert()
	p := db.PlayerFactory.CreateRelations().SetUser(&u).BuildInsert()

	data, _ := json.Marshal(map[string]interface{}{
		"playerId":   p.ID,
		"extrafield": "extratest",
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Method: "POST",
		Body:   io.NopCloser(bytes.NewBuffer(data)),
		Header: http.Header{},
	}
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("User", u)
	ValidatePlayer(c)

	if !httputils.StatusOK(w) {
		t.Fatalf("Wrong status code returned. Got (%d), want (%d). %v", w.Result().StatusCode, http.StatusOK, httputils.FormatError(w))
	}

	_, exists := c.Get("PlayerUser")

	if !exists {
		t.Fatalf("PlayerUser was not added to context")
	}
}
