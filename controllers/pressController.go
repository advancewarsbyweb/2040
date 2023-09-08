package controllers

import (
	"net/http"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/types"
	"github.com/awbw/2040/utils/auth"
	"github.com/gin-gonic/gin"
)

type PressController struct{}

var Press PressController

func NewPressController() PressController {
	return PressController{}
}

func init() {
	Press = NewPressController()
}

func (pc *PressController) Get(c *gin.Context) {

}

func (pc *PressController) Create(c *gin.Context) {
	type PressBody struct {
		types.Press
		Recipients []int `json:"recipients"`
	}

	var body PressBody

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// PlayerIsUser middleware
	u := auth.GetAuthenticatedUser(c)

	playerUser, err := db.PlayerRepo.FindPlayerUser(body.Press.PlayerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if playerUser.UserID != u.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "User sending press does not match logged in User",
		})
		return
	}
	playersInGame, err := db.PlayerRepo.FindPlayersByGame(playerUser.GameID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	pIDs := make(map[int]bool)
	for _, p := range playersInGame {
		pIDs[p.ID] = true
	}

	for _, pID := range body.Recipients {
		if _, validRecipient := pIDs[pID]; !validRecipient {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "One or more recipient is invalid",
			})
			return
		}
	}

	pressId, err := db.PressRepo.CreatePress(body.Press, body.Recipients)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	press, err := db.PressRepo.FindPress(pressId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, press)
}
