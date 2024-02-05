package unitmodels

import (
	"github.com/awbw/2040/models"
	unitnames "github.com/awbw/2040/types/units/names"
)

var TransportUnits []unitnames.UnitName

type transportUnit struct {
	models.Unit
}

func init() {
	TransportUnits = []unitnames.UnitName{
		unitnames.APC,
		unitnames.TCopter,
		unitnames.Lander,
		unitnames.BlackBoat,
	}
}
