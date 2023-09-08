package middleware

import (
	"net/http"

	"github.com/awbw/2040/utils/auth"
	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	loggedUser, err := auth.RequireAuth(c)

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
	}

	c.Set("User", loggedUser)
	c.Next()
}
