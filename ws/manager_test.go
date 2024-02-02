package ws

import (
	"log"
	"net/http"
	"testing"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/utils/auth"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../.env")
	db.NewTestDatabase()
}

func TestServeWS(t *testing.T) {

	u := db.UserFactory.Create().BuildInsert()

	w := http.Hijacker{}
	c, _ := gin.CreateTestContext(w)
	c.Request = &http.Request{
		Method: "GET",
		Header: http.Header{},
	}
	c.Request.Header.Add("Connection", "Upgrade")
	c.Request.Header.Add("Upgrade", "Websocket")
	c.Request.Header.Add("Sec-Websocket-Version", "13")
	c.Request.Header.Add("Sec-Websocket-Key", "Secret")
	tokenString, _ := auth.CreateTokenString(u)
	c.Request.AddCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		MaxAge:   60 * 60,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: false,
	})
	ClientManager.ServeWS(c)
	if c.IsWebsocket() {
		log.Fatalf("Websocket")
	}
}
