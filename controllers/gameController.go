package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/models"
	"github.com/awbw/2040/utils"
	"github.com/awbw/2040/ws"
	"github.com/awbw/2040/ws/events"
	"github.com/gin-gonic/gin"
)

func GetGame(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid Game ID provided",
		})
	}
	game, err := db.GameRepo.FindGame(gameId)

	c.JSON(http.StatusOK, gin.H{
		"game": game,
	})

}

func CreateGame(c *gin.Context) {

	var body models.Game

	c.Bind(&body)

	// Validation here

	_, err := db.GameRepo.CreateGame(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}

// Update Game with all of the given properties body
func UpdateGame(c *gin.Context) {

	var body struct {
		Game          models.Game `json:"game,omitempty"`
		UpdatedFields []string    `json:"updatedFields"`
	}
	c.Bind(&body)

	gameModel, err := db.GameRepo.UpdateGame(body.Game, body.UpdatedFields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"game": gameModel,
	})
}

func DeleteGame(c *gin.Context) {

	// Make sure request can delete given game

	c.Status(http.StatusOK)
}

func StartGame(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("id"))
	gameModel, err := db.GameRepo.FindGame(gameId)

	if !gameModel.StartDate.Valid {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Game is already started",
		})
		return
	}

	// This should be a middleware
	user := utils.GetAuthenticatedUser(c)
	if gameModel.Creator.ID != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Only the creator can start the game",
		})
		return
	}

	updatedFields := []string{"games_start_date"}
	_, err = db.GameRepo.UpdateGame(models.Game{StartDate: utils.NullTime()}, updatedFields)

	c.JSON(http.StatusOK, gin.H{
		"message": "Game started",
		"game":    gameModel,
	})

	// Send notification to users in lobby of game

	users, err := db.UserRepo.FindUsersByGame(gameId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	newNotification := events.NotificationResponse{
		Message: fmt.Sprintf("Game %s has started!", gameModel.Name),
		Url:     fmt.Sprintf("/play/%d", gameModel.ID),
	}

	ws.NotificationHandler(newNotification, users)
}

func JoinGame(c *gin.Context) {
	// RequireAuth middleware executed before

	_, err := db.PlayerRepo.CreatePlayer(models.Player{})
}

func UpdateTurn(c *gin.Context) {

	// Probably a Websocket Handler function instead

	// Increment turn count and update player

	// Function to update Turn here

}
