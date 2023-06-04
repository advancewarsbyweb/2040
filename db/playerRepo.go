package db

import (
	"errors"

	"github.com/awbw/2040/models"
)

type PlayerRepository struct{}

var PlayerRepo PlayerRepository

func NewPlayerRepo() PlayerRepository {
	return PlayerRepository{}
}

func init() {
	PlayerRepo = NewPlayerRepo()
}

func (r PlayerRepository) FindPlayer(id int) (models.Player, error) {
	var playerModel models.Player
	playerQuery := "SELECT * FROM awbw_players WHERE players_id = ?"

	err := DB.Get(playerModel, playerQuery, id)
	if err != nil {
		return models.Player{}, errors.New("Failed to find player with given ID")
	}

	return playerModel, nil
}

func (r PlayerRepository) FindPlayersByGame(gameId int) ([]models.Player, error) {

	var playerModels []models.Player
	playersQuery := "SELECT * FROM awbw_players WHERE players_games_id = ?"

	err := DB.Select(&playerModels, playersQuery, gameId)
	if err != nil {
		return []models.Player{}, errors.New("Failed to find players with given game ID")
	}

	return playerModels, nil
}

func (r PlayerRepository) CreatePlayer(body models.Player) (*models.Player, error) {
	return nil, nil
}
