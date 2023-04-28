package controllers

import (
	"fmt"
	"net/http"

	"github.com/awbw/2040/conf"
	"github.com/awbw/2040/models"
	"github.com/awbw/2040/utils"
	"github.com/awbw/2040/ws"
	"github.com/awbw/2040/ws/events"
	"github.com/gin-gonic/gin"
)

func GetGame(c *gin.Context) {
	game := utils.GetGame(c.Param("id"))

	c.JSON(http.StatusOK, gin.H{
		"game": game,
	})
}

func CreateGame(c *gin.Context) {

	var body models.Game

	c.Bind(&body)

	// Validation here

	result := conf.DB.Create(&body)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"game": result,
	})
}

// Update Game with all of the given properties body
func UpdateGame(c *gin.Context) {

	var body struct {
		Name      string `json:"name,omitempty"`
		StartDate string `json:"startDate,omitempty"`
		EndDate   string `json:"endDate,omitempty"`
	}

	c.Bind(&body)

	game := utils.GetGame(c.Param("id"))
	// Validation here

	conf.DB.Model(&body).Updates(body)

	c.JSON(http.StatusOK, gin.H{
		"game": game,
	})
}

func DeleteGame(c *gin.Context) {

	// Make sure request can delete given game
	conf.DB.Delete(&models.Game{}, c.Param("id"))

	c.Status(http.StatusOK)
}

func StartGame(c *gin.Context) {
	game := utils.GetGame(c.Param("id"))

	if !game.StartDate.Valid {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Game is already started",
		})
		return
	}

	user := utils.GetUser(c)
	if game.Creator.ID != user.ID {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Only the creator can start the game",
		})
		return
	}

	conf.DB.Model(&game).Updates(models.Game{StartDate: utils.NullTime()})

	c.JSON(http.StatusOK, gin.H{
		"message": "Game started",
		"game":    game,
	})

	// Send notification to users in lobby of game

	var users []models.User
	err := conf.DB.Model(&models.User{}).Preload("Player").Where("players_games_id = ?", game.ID).Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not find users to send notifications",
		})
	}

	newNotification := events.NewNotification{
		Message: fmt.Sprintf("Game %s has started!", game.Name),
		Url:     fmt.Sprintf("/play/%d", game.ID),
	}

	ws.SendNotificationHandler(newNotification, users)
}

func UpdateTurn(c *gin.Context) {

	// Increment turn count and update player

	user := utils.GetUser(c)
	game := utils.GetGame(c.Param("id"))
	if game.Turn.ID != user.ID {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not the user's turn",
		})
		return
	}

}
