package utils

import (
	"github.com/awbw/2040/conf"
	"github.com/awbw/2040/models"
)

func GetGame(gameId string) models.Game {
	var game models.Game
	conf.DB.First(&game, gameId)
	return game
}
