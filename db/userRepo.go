package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/awbw/2040/models"
	"github.com/awbw/2040/types"
)

type UserRepository struct {
	Columns []string
}

var UserRepo UserRepository

func NewUserRepository() UserRepository {
	return UserRepository{
		Columns: []string{
			"users_id",
			"users_username",
			"users_firstname",
			"users_lastname",
			"users_password",
			"users_email",
			"users_email_display",
			"users_inactive",
			"users_games",
			"users_uniq_id",
			"users_discord_id",
			"users_session_id",
			"users_activity_date",
			"users_boot_record",
			"users_shoals_display",
			"users_countries_id",
			"users_awbw_image_dir",
			"users_friend_date",
			"users_auto_scroll",
			"users_changelog_date",
			"ip",
			"users_awbw_last_page",
			"users_awbw_moderator",
			"users_awbw_days",
			"users_yourgames_date",
			"users_awbw_turn_email",
			"users_timezone",
			"users_email_messages",
			"users_show_warnings",
			"users_co_portraits",
			"users_tournament_date",
			"users_gridlines",
			"users_vacation",
			"users_last_vacation",
			"users_map_committee",
			"users_dor_fc",
			"users_game_animations",
		},
	}
}

func init() {
	UserRepo = NewUserRepository()
}

func (r *UserRepository) FindUser(id int) (*types.User, error) {
	var userModel models.User
	userQuery := fmt.Sprintf("SELECT %s FROM ofua_users WHERE users_id = ?", strings.Join(UserRepo.Columns, ","))
	err := DB.Get(&userModel, userQuery, id)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find player with given ID (%d): %s", id, err.Error()))
	}

	u := types.NewUser(userModel)
	return &u, nil
}

func (r *UserRepository) FindUserByPlayerId(id int) (*types.User, error) {
	var m models.User
	query := `SELECT ofua_users.* from ofua_users, awbw_players
		WHERE players_id = ?
		AND players_users_id = users_id`
	err := DB.Get(&m, query, id)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find player with given ID (%d): %s", id, err.Error()))
	}

	t := types.NewUser(m)
	return &t, nil
}

func (r *UserRepository) CreateUser(body types.User) (int, error) {
	createQuery := FormatCreateQuery("ofua_users", UserRepo.Columns)
	res, err := DB.NamedExec(createQuery, body)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("Could not create new User: %s", err.Error()))
	}

	userId, err := res.LastInsertId()

	return int(userId), nil
}

func (r *UserRepository) FindUsersByGame(gameId int) ([]models.User, error) {

	var userModels []models.User
	usersQuery := "SELECT ofua_users.* FROM awbw_players, ofua_users WHERE players_games_id = ? AND players_users_id = users_id"

	err := DB.Select(&userModels, usersQuery, gameId)
	if err != nil {
		return []models.User{}, errors.New("Failed to find players with given game ID")
	}

	return userModels, nil
}
