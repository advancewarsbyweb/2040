package controllers

import (
	"net/http"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/models"
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

func (pc *PressController) Get(c *gin.Context) {

}

func (pc *PressController) Create(c *gin.Context) {

	var body PressBody

	err := c.Bind(&body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	p, _ := c.Get("PlayerUser")
	playerUser := p.(models.Player)
	playersInGame, err := db.PlayerRepo.FindPlayersByGame(playerUser.GameID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
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
		}
	}

	pressId, err := db.PressRepo.CreatePress(body.Press, body.Recipients)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	press, err := db.PressRepo.FindPress(pressId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, press)
}
