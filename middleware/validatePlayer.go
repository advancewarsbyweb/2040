package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/utils/auth"
	"github.com/gin-gonic/gin"
)

type Body struct {
	PlayerID int `json:"playerId"`
}

func ValidatePlayer(c *gin.Context) {

	var playerId int

	if c.Param("playerId") != "" {
		parsedPID, err := strconv.Atoi(c.Param("playerId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		playerId = parsedPID
	} else {
		var body Body
		c.Bind(&body)
		if body.PlayerID != 0 {
			playerId = body.PlayerID
		}
	}

	if playerId == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "No Player ID found in query params or request body"})
		return
	}

	playerRelations, err := db.PlayerRepo.FindPlayerRelations(playerId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// EndDate means the Game is over and we do not need to check for a valid Player
	if !playerRelations.Game.GetEndDate().IsZero() {
		fmt.Printf("%v", playerRelations.Game.GetEndDate())
		c.Set("ValidationSkip", true)
		c.Next()
		return
	}
	u, exists := auth.GetAuthenticatedUser(c)
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Authenticated User could not be found"})
		return
	}

	if playerRelations.UserID != u.ID {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Player does not match logged in user",
		})
		return
	}

	c.Set("PlayerUser", playerRelations)
	c.Next()
}
