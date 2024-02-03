package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	db.NewTestDatabase()
}

func TestCreatePress(t *testing.T) {
	assert := assert.New(t)

	g := db.Game.Create().BuildInsert()

	p1 := db.PlayerFactory.Create().SetGame(&g).BuildInsert()
	p2 := db.PlayerFactory.Create().SetGame(&g).BuildInsert()
	p3 := db.PlayerFactory.Create().CreateUser().SetGame(&g).BuildInsert()

	press := db.Press.Create().SetPlayer(&p3).Build()

	data, _ := json.Marshal(map[string]interface{}{
		"playerId":   press.PlayerID,
		"subject":    press.GetSubject(),
		"text":       press.GetText(),
		"recipients": []int{p1.ID, p2.ID},
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Method: "POST",
		Body:   io.NopCloser(bytes.NewBuffer(data)),
		Header: http.Header{},
	}
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("PlayerUser", p3)

	Press.Create(c)

	if w.Result().StatusCode != http.StatusOK {
		t.Fatalf("Wrong status code returned. Got (%d), want (%d)", w.Result().StatusCode, http.StatusOK)
	}

	var res models.Press
	json.Unmarshal(w.Body.Bytes(), &res)

	assert.Equal(press.GetSubject(), res.GetSubject(), "Press subject does not match given one. Got (%s), want (%s)")
}
