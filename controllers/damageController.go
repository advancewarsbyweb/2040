package controllers

import (
	"net/http"
	"strconv"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/types/calculator"
	unitmodels "github.com/awbw/2040/types/units"
	"github.com/gin-gonic/gin"
)

type damageController struct{}

var Damage damageController

func NewDamageController() damageController {
	return damageController{}
}

func init() {
	Damage = NewDamageController()
}

func (dc *damageController) GetPreview(c *gin.Context) {
	q := c.Request.URL.Query()

	attId, err := strconv.Atoi(q["attacker"][0])
	defId, err2 := strconv.Atoi(q["defender"][0])
	if err != nil || err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid Attacker or Defender ID provided"})
		return
	}

	a, err := db.UnitRepo.FindUnit(attId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not find attacker"})
	}

	d, err := db.UnitRepo.FindUnit(defId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not find defender"})
	}

	calc := calculator.NewCalculator(unitmodels.CreateUnit(a), unitmodels.CreateUnit(d))
	_ = calc.CalcPreviewResult()

	c.JSON(http.StatusOK, gin.H{
		"attacker": calc.Attacker.Preview,
		"defender": calc.Defender.Preview,
	})
}
