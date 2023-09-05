package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/awbw/2040/db"
	gamecolumns "github.com/awbw/2040/models/columnNames/game"
	"github.com/awbw/2040/types"
	"github.com/awbw/2040/utils"
	"github.com/awbw/2040/ws"
	"github.com/awbw/2040/ws/events"
	"github.com/gin-gonic/gin"
)

type GameController struct{}

var Game GameController

func NewGameController() GameController {
	return GameController{}
}

func init() {
	Game = NewGameController()
}

func (gc *GameController) Get(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid Game ID provided",
		})
		return
	}
	game, err := db.GameRepo.FindGame(gameId)

	c.JSON(http.StatusOK, game)
}

func (gc *GameController) Create(c *gin.Context) {

	var body types.Game

	c.Bind(&body)

	// Validation here

	gameId, err := db.GameRepo.CreateGame(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

    game, err := db.GameRepo.FindGame(gameId)
    if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
    }

	c.JSON(http.StatusOK, game)
}

// Update Game with all of the given properties body
func (gc *GameController) Update(c *gin.Context) {

	var body struct {
		ID            int                    `json:"id"`
		UpdatedFields map[string]interface{} `json:"updatedFields"`
	}
	c.Bind(&body)

	g, err := db.GameRepo.UpdateGame(body.ID, body.UpdatedFields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"game": g,
	})
}

func (gc *GameController) Delete(c *gin.Context) {

	// Make sure request can delete given game

	c.Status(http.StatusOK)
}

func (gc *GameController) Start(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("id"))
	gameModel, err := db.GameRepo.FindGame(gameId)

	if gameModel.StartDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Game is already started",
		})
		return
	}

	// This should be a middleware
	user := utils.GetAuthenticatedUser(c)
	if gameModel.CreatorID != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Only the creator can start the game",
		})
		return
	}

	updatedFields := map[string]interface{}{gamecolumns.StartDate: time.Now()}
	_, err = db.GameRepo.UpdateGame(gameId, updatedFields)

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

func (gc *GameController) Join(c *gin.Context) {
	// RequireAuth middleware executed before

	//_, err := db.PlayerRepo.CreatePlayer(models.Player{})
}

func UpdateTurn(c *gin.Context) {

	// Probably a Websocket Handler function instead

	// Increment turn count and update player

	// Function to update Turn here

}
