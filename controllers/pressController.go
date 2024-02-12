package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/models"
	"github.com/awbw/2040/ws"
	eventdata "github.com/awbw/2040/ws/events/data"
	eventtypes "github.com/awbw/2040/ws/events/types"
	"github.com/gin-gonic/gin"
)

type PressController struct{}

type PressBody struct {
	models.Press
	Recipients []int `json:"recipients"`
}

var Press PressController

func NewPressController() PressController {
	return PressController{}
}

func init() {
	Press = NewPressController()
}

func (pc *PressController) GetAll(c *gin.Context) {
	_, err := strconv.Atoi(c.Param("playerId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	// Get every Press that the playerID sent and received
}

func (pc *PressController) Create(c *gin.Context) {

	var body PressBody

	err := c.Bind(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	p, _ := c.Get("PlayerUser")
	playerUser := p.(models.Player)
	playersInGame, err := db.PlayerRepo.FindPlayersRelationsByGame(playerUser.GameID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	pIDs := make(map[int]bool)
	for _, p := range playersInGame {
		pIDs[p.ID] = true
	}

	for _, pID := range body.Recipients {
		if _, validRecipient := pIDs[pID]; !validRecipient {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "One or more recipient is invalid",
			})
			return
		}
	}

	pressId, err := db.PressRepo.CreatePress(body.Press, body.Recipients)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	press, err := db.PressRepo.FindPress(pressId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Could not find press with inserted ID (%d). %s", pressId, err.Error())})
		return
	}

	c.JSON(http.StatusOK, press)

	var users []models.User
	for _, p := range playersInGame {
		users = append(users, *p.User)
	}

	pressNotif := ws.Event{
		Type:      eventtypes.NotificationResponse,
		Timestamp: time.Now(),
		Users:     users,
		Data: eventdata.NotificationResponse{
			Type:        "NewPress",
			ContextType: "Game",
			ContextID:   playersInGame[0].GameID,
		},
	}

	ws.ClientManager.Publish(pressNotif)
}
