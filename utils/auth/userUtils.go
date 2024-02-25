package auth

import (
	"github.com/awbw/2040/models"
	"github.com/gin-gonic/gin"
)

func GetAuthenticatedUser(c *gin.Context) (*models.User, bool) {
	loggedUser, exists := c.Get("User")
	if !exists {
		return nil, exists
	}
	u := loggedUser.(models.User)

	return &u, exists
}
