package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/awbw/2040/models"
)

type PlayerRepository struct {
	Columns []string
}

var PlayerRepo PlayerRepository

func NewPlayerRepo() PlayerRepository {
	return PlayerRepository{
		Columns: []string{
			"players_games_id",
			"players_users_id",
			"players_countries_id",
			"players_co_id",
			"players_funds",
			"players_turn",
			"players_old_interface",
			"players_eliminated",
			"players_last_read",
			"players_turn_start",
			"players_turn_clock",
			"players_pause_timer",
			"players_co_power",
			"players_co_power_on",
			"players_order",
			"players_accept_draw",
			"players_co_max_power",
			"players_co_max_spower",
			"players_co_image",
			"players_team",
			"players_aet_count",
			"players_turn_flag",
		},
	}
}

func init() {
	PlayerRepo = NewPlayerRepo()
}

// PlayerModel with null UserModel in it
func (r *PlayerRepository) FindPlayer(id int) (*models.Player, error) {
	var playerModel models.Player
	playerQuery := "SELECT * FROM awbw_players WHERE players_id = ?"

	err := DB.Get(&playerModel, playerQuery, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find player with given ID: %s", err.Error()))
	}

	return &playerModel, nil
}

// PlayerModel with non null UserModel in it
func (r *PlayerRepository) FindPlayerUser(id int) (*models.Player, error) {
	var playerModel models.Player
	playerQuery := fmt.Sprintf(`SELECT awbw_players.*, %s FROM awbw_players, ofua_users
        WHERE players_id = ?
        AND users_id = players_users_id`,
		strings.Join(UserRepo.Columns, ","),
	)

	err := DB.Get(&playerModel, playerQuery, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find playerUser with given ID: %s", err.Error()))
	}

	return &playerModel, nil
}

func (r *PlayerRepository) FindPlayersByGame(gameId int) ([]models.Player, error) {

	var playerModels []models.Player
	query := "SELECT * FROM awbw_players WHERE players_games_id = ?"

	err := DB.Select(&playerModels, query, gameId)
	if err != nil {
		return nil, errors.New("Failed to find players with given game ID")
	}

	var pTypes []models.Player

	for _, m := range playerModels {
		pTypes = append(pTypes, m)
	}

	return pTypes, nil
}

func (r *PlayerRepository) CreatePlayer(body models.Player) (int, error) {
	createQuery := FormatCreateQuery("awbw_players", PlayerRepo.Columns)

	res, err := DB.NamedExec(createQuery, body)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("Could not create new Player: %s", err.Error()))
	}

	playerId, err := res.LastInsertId()

	return int(playerId), nil
}
