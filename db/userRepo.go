package db

import (
	"errors"

	"github.com/awbw/2040/models"
)

type UserRepository struct{}

var UserRepo UserRepository

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func init() {
	UserRepo = NewUserRepository()
}

func (r UserRepository) FindUser(id int) (models.User, error) {
	return models.User{}, nil
}

func (r UserRepository) FindUsersByGame(gameId int) ([]models.User, error) {

	var userModels []models.User
	usersQuery := "SELECT u.* FROM awbw_players AS p, ofua_users AS u WHERE players_games_id = ? AND players_users_id = users_id"

	err := DB.Select(&userModels, usersQuery, gameId)
	if err != nil {
		return []models.User{}, errors.New("Failed to find players with given game ID")
	}

	return userModels, nil
}
