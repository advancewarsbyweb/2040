package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/factories"
	"github.com/awbw/2040/types"
	"github.com/gin-gonic/gin"
)

func init() {
	db.NewTestDatabase()
}

func TestCreatePress(t *testing.T) {

	g := factories.Game.Create().BuildInsert()

	p1 := factories.Player.Create().SetGame(&g).BuildInsert()
	p2 := factories.Player.Create().SetGame(&g).BuildInsert()
	p3 := factories.Player.Create().CreateUser().SetGame(&g).BuildInsert()

	press := factories.Press.Create().SetPlayer(&p3).Build()

	data, _ := json.Marshal(map[string]interface{}{
		"playerId":   press.PlayerID,
		"subject":    press.Subject,
		"text":       press.Text,
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

	var res types.Press
	json.Unmarshal(w.Body.Bytes(), &res)

	if press.Subject != res.Subject {
		t.Fatalf("Press subject does not match given one. Got (%s), want (%s)", press.Subject, res.Subject)
	}
}
