package unitmodels

import unitnames "github.com/awbw/2040/types/units/names"

var Vehicles []unitnames.UnitName

func init() {
	Vehicles = []unitnames.UnitName{
		unitnames.Recon,
		unitnames.APC,
		unitnames.Artillery,
		unitnames.Tank,
		unitnames.AntiAir,
		unitnames.Missile,
		unitnames.Rocket,
		unitnames.MDTank,
		unitnames.Piperunner,
		unitnames.Neotank,
		unitnames.MegaTank,
	}
}
