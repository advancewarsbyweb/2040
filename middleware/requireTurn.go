package middleware

import (
	"net/http"
	"strconv"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/utils"
	"github.com/gin-gonic/gin"
)

func RequireTurn(c *gin.Context) {

	gameId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid Game ID provided",
		})
	}

	user := utils.GetAuthenticatedUser(c)
	gameModel, err := db.GameRepo.FindGame(gameId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if gameModel.TurnID != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Can not update turn, currently not the user's turn!",
		})
		return
	}

	c.Next()
}
