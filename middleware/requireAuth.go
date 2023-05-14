package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/awbw/2040/utils/auth"
	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := auth.GetTokenCookie(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New(fmt.Sprintf("Could not get authorization cookie: %s", err.Error())))
	}

	loggedUser, err := auth.RequireAuth(tokenString)

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
	}

	c.Set("user", loggedUser)
	c.Next()
}
