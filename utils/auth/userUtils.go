package auth

import (
	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types"
	"github.com/gin-gonic/gin"
)

func GetAuthenticatedUser(c *gin.Context) types.User {
	loggedUser, _ := c.Get("User")

	t := types.NewUser(loggedUser.(models.User))

	return t
}
