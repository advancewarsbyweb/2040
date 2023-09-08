package auth

import (
	"github.com/awbw/2040/types"
	"github.com/gin-gonic/gin"
)

func GetAuthenticatedUser(c *gin.Context) types.User {
	loggedUser, _ := c.Get("User")
	u := loggedUser.(types.User)

	return u
}
