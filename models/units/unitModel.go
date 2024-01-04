package models

import (
	"github.com/awbw/2040/models"
)

type unit struct {
	id       int
	playerID int
	gameID   int
	x        int
	y        int
	subDive  string
	moved    int
	fired    int
	hp       float64
	cargo1ID int
	cargo2ID int
	carried  string
	game     *models.Game   `db:""`
	player   *models.Player `db:""`
	baseUnit
}

// Create a Unit from a model retrieved from the database
func CreateUnit(m *unit) Unit {
	return MakeUnit[m.name](m)
}

// Set the Unit's properties with the model received from the database
// This is used after creating a Unit type struct to have the appropriate Unit methods
// For example in NewInfantry
func (u *unit) SetUnitProperties(m *unit) {
	u.id = m.id
	u.playerID = m.playerID
	u.gameID = m.gameID
	u.x = m.x
	u.y = m.y
	u.subDive = m.subDive
	u.moved = m.moved
	u.fired = m.fired
	u.hp = m.hp
	u.cargo1ID = m.cargo1ID
	u.cargo2ID = m.cargo2ID
	u.carried = m.carried
	u.game = m.game
	u.player = m.player
	u.fuel = m.fuel
}
