package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/awbw/2040/db"
	"github.com/gin-gonic/gin"
)

func init() {
	db.NewTestDatabase()
}

func TestValidatePlayer(t *testing.T) {
	u := db.User.Create().BuildInsert()
	p := db.Player.Create().SetUser(&u).BuildInsert()

	data, _ := json.Marshal(map[string]interface{}{
		"playerId": p.ID,
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

	_, exists := c.Get("PlayerUser")

	if !exists {
		t.Fatalf("PlayerUser was not added to context")
	}
}
