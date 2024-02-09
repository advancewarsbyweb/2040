package middleware

import (
	"net/http"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/utils/auth"
	"github.com/gin-gonic/gin"
)

func ValidatePlayer(c *gin.Context) {
	u := auth.GetAuthenticatedUser(c)

	type Body struct {
		ID int `json:"playerId"`
	}

	var playerID Body
	c.Bind(&playerID)

	playerUser, err := db.PlayerRepo.FindPlayerRelations(playerID.ID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	if playerUser.UserID != u.ID {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Player does not match logged in user",
		})
	}

	c.Set("PlayerUser", playerUser)
	c.Next()
}
