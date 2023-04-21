package utils

import (
	"github.com/awbw/2040/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) models.User {
	loggedUser, _ := c.Get("User")

	return loggedUser.(models.User)
}
