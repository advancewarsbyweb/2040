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
		Press      types.Press
		Recipients []int
	}

	var body PressBody

	c.Bind(&body)

	u := auth.GetAuthenticatedUser(c)

	userType, err := db.UserRepo.FindUserByPlayerId(body.Press.PlayerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if userType.ID != u.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "User sending press does not match logged in User",
		})
		return
	}
	// Make sure each recipient is valid

}
