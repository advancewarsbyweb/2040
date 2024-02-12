package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/models"
	gamecolumns "github.com/awbw/2040/models/columnNames/game"
	"github.com/awbw/2040/utils"
	"github.com/awbw/2040/ws"
	eventdata "github.com/awbw/2040/ws/events/data"
	eventtypes "github.com/awbw/2040/ws/events/types"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type GameController struct{}

var Game GameController

var validate *validator.Validate

func NewGameController() GameController {
	return GameController{}
}

func init() {
	Game = NewGameController()

	validate = validator.New(validator.WithRequiredStructEnabled())
}

func (gc *GameController) Get(c *gin.Context) {
	gameId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid Game ID provided"})
		return
	}
	game, err := db.GameRepo.FindGame(gameId)

	c.JSON(http.StatusOK, game)
}

func (gc *GameController) Create(c *gin.Context) {

	var body models.Game

	c.Bind(&body)

	err := validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed for some fields",
			"errors":  utils.ValidatorErrors(err),
		})
		return
	}

	gameId, err := db.GameRepo.CreateGame(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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
		ID            int                    `json:"id" validate:"gt=0"`
		UpdatedFields map[string]interface{} `json:"updatedFields" validate:"required,dive,keys,eq=name|eq=password|eq=maps_id|eq=private|eq=funds|eq=capture_win|eq=day|eq=comment|eq=boot_interval|eq=unit_limit|eq=team|eq=aet_interval|eq=use_powers"`
	}
	c.Bind(&body)

	err := validate.Struct(body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	g, err := db.GameRepo.UpdateGame(body.ID, body.UpdatedFields)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
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
	game, err := db.GameRepo.FindGame(gameId)

	if game.GetStartDate().IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Game is already started"})
		return
	}

	// This should be a middleware
	user := utils.GetAuthenticatedUser(c)
	if game.CreatorID != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Only the creator can start the game"})
		return
	}

	updatedFields := map[string]interface{}{gamecolumns.StartDate: time.Now()}
	_, err = db.GameRepo.UpdateGame(gameId, updatedFields)

	c.JSON(http.StatusOK, gin.H{
		"message": "Game started",
		"game":    game,
	})

	// Send Notification to users in Lobby of Game
	users, err := db.UserRepo.FindUsersByGame(gameId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	notif := ws.Event{
		Type:      eventtypes.NotificationResponse,
		Timestamp: time.Now(),
		Users:     users,
		Data: eventdata.NotificationResponse{
			Type:        "GameStart",
			ContextType: "Game",
			ContextID:   game.ID,
		},
	}

	ws.ClientManager.Publish(notif)
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
