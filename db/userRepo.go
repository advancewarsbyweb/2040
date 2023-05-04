package db

import "github.com/awbw/2040/models"

type UserRepository struct{}

var UserRepo UserRepository

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func init() {

}

func (r UserRepository) FindUser(id int) models.User {
	return models.User{}
}
