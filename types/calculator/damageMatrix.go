package calculator

import (
	unitnames "github.com/awbw/2040/types/units/names"
)

type unitDamage map[unitnames.UnitName][2]int

type damageMatrix map[unitnames.UnitName]unitDamage

var (
	Units        map[unitnames.UnitName]int
	DamageMatrix damageMatrix
)

type IDamageMatrix interface {
	FindPrimary(a unitnames.UnitName, d unitnames.UnitName)
	FindSecondary(a unitnames.UnitName, d unitnames.UnitName)
}

func (m *damageMatrix) FindPrimary(a unitnames.UnitName, d unitnames.UnitName) int {
	return DamageMatrix[a][d][0]
}

func (m *damageMatrix) FindSecondary(a unitnames.UnitName, d unitnames.UnitName) int {
	return DamageMatrix[a][d][1]
}

func init() {
	/*
		// Arbitrary index so the matrix is not full of repeated names ?
		allUnits := []unitnames.UnitName{
			unitnames.AntiAir,
			unitnames.APC,
			unitnames.Artillery,
			unitnames.BCopter,
			unitnames.Battleship,
			unitnames.BlackBoat,
			unitnames.BlackBomb,
			unitnames.Bomber,
			unitnames.Carrier,
			unitnames.Cruiser,
			unitnames.Fighter,
			unitnames.Infantry,
			unitnames.Lander,
			unitnames.MDTank,
			unitnames.Mech,
			unitnames.MegaTank,
			unitnames.Missile,
			unitnames.Neotank,
			unitnames.Piperunner,
			unitnames.Recon,
			unitnames.Stealth,
			unitnames.Sub,
			unitnames.TCopter,
			unitnames.Tank,
		}
	*/

	DamageMatrix = damageMatrix{
		unitnames.AntiAir: unitDamage{
			unitnames.AntiAir:    d(45, -1),
			unitnames.APC:        d(70, -1),
			unitnames.Artillery:  d(45, -1),
			unitnames.BCopter:    d(105, -1),
			unitnames.Battleship: d(-1, -1),
			unitnames.BlackBoat:  d(-1, -1),
			unitnames.BlackBomb:  d(120, -1),
			unitnames.Bomber:     d(75, -1),
			unitnames.Carrier:    d(-1, -1),
			unitnames.Cruiser:    d(-1, -1),
			unitnames.Fighter:    d(65, -1),
			unitnames.Infantry:   d(105, -1),
			unitnames.Lander:     d(-1, -1),
			unitnames.MDTank:     d(10, -1),
			unitnames.Mech:       d(105, -1),
			unitnames.MegaTank:   d(1, -1),
			unitnames.Missile:    d(55, -1),
			unitnames.Neotank:    d(5, -1),
			unitnames.Piperunner: d(25, -1),
			unitnames.Recon:      d(60, -1),
			unitnames.Stealth:    d(75, -1),
			unitnames.Sub:        d(-1, -1),
			unitnames.TCopter:    d(105, -1),
			unitnames.Tank:       d(25, -1),
		},
		unitnames.APC: unitDamage{},
		unitnames.Artillery: unitDamage{
			unitnames.Infantry:   d(90, 0),
			unitnames.Mech:       d(85, 0),
			unitnames.Recon:      d(80, 0),
			unitnames.Tank:       d(70, 0),
			unitnames.MDTank:     d(45, 0),
			unitnames.Neotank:    d(40, 0),
			unitnames.MegaTank:   d(15, 0),
			unitnames.AntiAir:    d(75, 0),
			unitnames.Artillery:  d(75, 0),
			unitnames.Rocket:     d(80, 0),
			unitnames.Piperunner: d(70, 0),
			unitnames.Missile:    d(80, 0),
			unitnames.APC:        d(60, 20),
			unitnames.Cruiser:    d(50, 0),
			unitnames.Sub:        d(60, 0),
			unitnames.Battleship: d(40, 0),
			unitnames.Carrier:    d(45, 0),
			unitnames.Lander:     d(55, 0),
			unitnames.BlackBoat:  d(55, 0),
		},
		unitnames.BCopter: unitDamage{
			unitnames.Infantry:   d(0, 75),
			unitnames.Mech:       d(0, 75),
			unitnames.Recon:      d(55, 30),
			unitnames.Tank:       d(55, 6),
			unitnames.MDTank:     d(25, 1),
			unitnames.Neotank:    d(20, 1),
			unitnames.MegaTank:   d(10, 0),
			unitnames.AntiAir:    d(25, 6),
			unitnames.Artillery:  d(65, 25),
			unitnames.Rocket:     d(65, 35),
			unitnames.Piperunner: d(55, 6),
			unitnames.Missile:    d(65, 35),
			unitnames.APC:        d(60, 20),
			unitnames.BCopter:    d(0, 65),
			unitnames.TCopter:    d(0, 95),
			unitnames.Cruiser:    d(25, 0),
			unitnames.Sub:        d(25, 0),
			unitnames.Battleship: d(25, 0),
			unitnames.Carrier:    d(25, 0),
			unitnames.Lander:     d(25, 0),
			unitnames.BlackBoat:  d(25, 0),
		},
		unitnames.Battleship: unitDamage{
			unitnames.Infantry:   d(95, 0),
			unitnames.Mech:       d(90, 0),
			unitnames.Recon:      d(90, 0),
			unitnames.Tank:       d(80, 0),
			unitnames.MDTank:     d(55, 0),
			unitnames.Neotank:    d(50, 0),
			unitnames.MegaTank:   d(25, 0),
			unitnames.AntiAir:    d(85, 0),
			unitnames.Artillery:  d(80, 0),
			unitnames.Rocket:     d(85, 0),
			unitnames.Piperunner: d(80, 0),
			unitnames.Missile:    d(90, 0),
			unitnames.APC:        d(80, 0),
			unitnames.Cruiser:    d(95, 0),
			unitnames.Sub:        d(95, 0),
			unitnames.Battleship: d(50, 0),
			unitnames.Carrier:    d(60, 0),
			unitnames.Lander:     d(95, 0),
			unitnames.BlackBoat:  d(95, 0),
		},
		unitnames.BlackBoat: unitDamage{},
		unitnames.BlackBomb: unitDamage{},
		unitnames.Bomber: unitDamage{
			unitnames.Infantry:   d(110, 0),
			unitnames.Mech:       d(110, 0),
			unitnames.Recon:      d(105, 0),
			unitnames.Tank:       d(105, 0),
			unitnames.MDTank:     d(95, 0),
			unitnames.Neotank:    d(90, 0),
			unitnames.MegaTank:   d(35, 00),
			unitnames.AntiAir:    d(95, 0),
			unitnames.Artillery:  d(105, 0),
			unitnames.Rocket:     d(105, 0),
			unitnames.Piperunner: d(105, 0),
			unitnames.Missile:    d(105, 0),
			unitnames.APC:        d(105, 0),
			unitnames.Cruiser:    d(50, 0),
			unitnames.Sub:        d(95, 0),
			unitnames.Battleship: d(75, 0),
			unitnames.Carrier:    d(105, 0),
			unitnames.Lander:     d(95, 0),
			unitnames.BlackBoat:  d(75, 0),
		},
		unitnames.Carrier: unitDamage{
			unitnames.Fighter:   d(100, 0),
			unitnames.Bomber:    d(100, 0),
			unitnames.Stealth:   d(100, 0),
			unitnames.BCopter:   d(115, 0),
			unitnames.TCopter:   d(115, 0),
			unitnames.BlackBomb: d(120, 0),
		},
		unitnames.Cruiser: unitDamage{
			unitnames.Fighter:    d(0, 85),
			unitnames.Bomber:     d(0, 100),
			unitnames.Stealth:    d(0, 100),
			unitnames.BCopter:    d(0, 105),
			unitnames.TCopter:    d(0, 105),
			unitnames.BlackBomb:  d(0, 120),
			unitnames.Cruiser:    d(25, 0),
			unitnames.Sub:        d(90, 0),
			unitnames.Battleship: d(5, 0),
			unitnames.Carrier:    d(5, 0),
			unitnames.Lander:     d(25, 0),
			unitnames.BlackBoat:  d(25, 0),
		},
		unitnames.Fighter: unitDamage{
			unitnames.Fighter:   d(55, 0),
			unitnames.Bomber:    d(100, 0),
			unitnames.Stealth:   d(85, 0),
			unitnames.BCopter:   d(120, 0),
			unitnames.TCopter:   d(120, 0),
			unitnames.BlackBomb: d(120, 0),
		},
		unitnames.Infantry: unitDamage{
			unitnames.AntiAir:    d(0, 5),
			unitnames.APC:        d(0, 14),
			unitnames.Artillery:  d(0, 15),
			unitnames.BCopter:    d(0, 7),
			unitnames.Infantry:   d(0, 55),
			unitnames.MDTank:     d(0, 1),
			unitnames.Mech:       d(0, 45),
			unitnames.MegaTank:   d(0, 1),
			unitnames.Missile:    d(0, 26),
			unitnames.Neotank:    d(0, 1),
			unitnames.Piperunner: d(0, 5),
			unitnames.Recon:      d(0, 12),
			unitnames.Rocket:     d(0, 25),
			unitnames.TCopter:    d(0, 30),
			unitnames.Tank:       d(0, 5),
		},
		unitnames.Lander: unitDamage{},
		unitnames.MDTank: unitDamage{
			unitnames.AntiAir:    d(105, 7),
			unitnames.APC:        d(105, 45),
			unitnames.Artillery:  d(105, 45),
			unitnames.BCopter:    d(0, 12),
			unitnames.Battleship: d(10, 0),
			unitnames.BlackBoat:  d(35, 0),
			unitnames.Carrier:    d(10, 0),
			unitnames.Cruiser:    d(30, 0),
			unitnames.Infantry:   d(0, 105),
			unitnames.Lander:     d(35, 0),
			unitnames.MDTank:     d(55, 1),
			unitnames.Mech:       d(0, 95),
			unitnames.MegaTank:   d(25, 1),
			unitnames.Missile:    d(105, 35),
			unitnames.Neotank:    d(45, 1),
			unitnames.Piperunner: d(85, 8),
			unitnames.Recon:      d(105, 45),
			unitnames.Sub:        d(10, 0),
			unitnames.TCopter:    d(0, 45),
			unitnames.Tank:       d(85, 8),
		},
		unitnames.Mech: unitDamage{
			unitnames.AntiAir:    d(65, 5),
			unitnames.APC:        d(75, 20),
			unitnames.Artillery:  d(70, 32),
			unitnames.BCopter:    d(-1, 9),
			unitnames.Battleship: d(-1, -1),
			unitnames.BlackBoat:  d(-1, -1),
			unitnames.BlackBomb:  d(-1, -1),
			unitnames.Bomber:     d(-1, -1),
			unitnames.Carrier:    d(-1, -1),
			unitnames.Cruiser:    d(-1, -1),
			unitnames.Fighter:    d(-1, -1),
			unitnames.Infantry:   d(-1, 65),
			unitnames.Lander:     d(-1, -1),
			unitnames.MDTank:     d(15, 1),
			unitnames.Mech:       d(-1, 55),
			unitnames.MegaTank:   d(5, 1),
			unitnames.Missile:    d(85, 35),
			unitnames.Neotank:    d(15, 1),
			unitnames.Piperunner: d(55, 6),
			unitnames.Recon:      d(85, 18),
			unitnames.Stealth:    d(-1, -1),
			unitnames.Sub:        d(-1, -1),
			unitnames.TCopter:    d(-1, 35),
			unitnames.Tank:       d(55, 6),
		},
		unitnames.MegaTank: unitDamage{
			unitnames.Infantry: d(135, 0),
		},
		unitnames.Missile: unitDamage{
			// Fill
		},
		unitnames.Neotank: unitDamage{
			unitnames.AntiAir:    d(115, 17),
			unitnames.APC:        d(125, 65),
			unitnames.Artillery:  d(115, 65),
			unitnames.BCopter:    d(0, 22),
			unitnames.Battleship: d(15, 0),
			unitnames.BlackBoat:  d(40, 0),
			unitnames.Carrier:    d(15, 0),
			unitnames.Cruiser:    d(30, 0),
			unitnames.Infantry:   d(0, 125),
			unitnames.Lander:     d(40, 0),
			unitnames.MDTank:     d(75, 1),
			unitnames.Mech:       d(0, 115),
			unitnames.MegaTank:   d(35, 1),
			unitnames.Missile:    d(125, 55),
			unitnames.Neotank:    d(55, 1),
			unitnames.Piperunner: d(105, 10),
			unitnames.Recon:      d(125, 65),
			unitnames.Sub:        d(15, 0),
			unitnames.TCopter:    d(55, 0),
			unitnames.Tank:       d(105, 10),
		},
		unitnames.Piperunner: unitDamage{
			// Fill
		},
		unitnames.Recon: unitDamage{
			// Fill
		},
		unitnames.Rocket: unitDamage{
			// Fill
		},
		unitnames.Stealth: unitDamage{
			// Fill
		},
		unitnames.Sub: unitDamage{
			// Fill
		},
		unitnames.TCopter: unitDamage{},
		unitnames.Tank: unitDamage{
			unitnames.AntiAir:    d(65, 6),
			unitnames.APC:        d(75, 45),
			unitnames.Artillery:  d(70, 45),
			unitnames.BCopter:    d(0, 10),
			unitnames.Battleship: d(1, 0),
			unitnames.BlackBoat:  d(10, 0),
			unitnames.Carrier:    d(1, 0),
			unitnames.Cruiser:    d(5, 0),
			unitnames.Infantry:   d(0, 75),
			unitnames.Lander:     d(10, 0),
			unitnames.MDTank:     d(15, 1),
			unitnames.Mech:       d(0, 70),
			unitnames.MegaTank:   d(15, 1),
			unitnames.Missile:    d(85, 30),
			unitnames.Neotank:    d(15, 1),
			unitnames.Piperunner: d(55, 6),
			unitnames.Recon:      d(55, 6),
			unitnames.Sub:        d(1, 0),
			unitnames.TCopter:    d(0, 40),
			unitnames.Tank:       d(55, 6),
		},
	}
}

// Array for Primary & Secondary fire damage
func d(primaryFire int, secondaryFire int) [2]int {
	return [2]int{primaryFire, secondaryFire}
}
