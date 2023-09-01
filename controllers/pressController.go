package controllers

import (
	"github.com/awbw/2040/types"
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
	type body struct {
		Press      types.Press
		Recipients []int
	}
}
