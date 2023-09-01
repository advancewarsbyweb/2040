package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/awbw/2040/models"
	gamecolumns "github.com/awbw/2040/models/columnNames/game"
	"github.com/awbw/2040/types"
)

type GameRepository struct{}

var GameRepo GameRepository

func NewGameRepo() GameRepository {
	return GameRepository{}
}

func init() {
	GameRepo = NewGameRepo()
}

func (r *GameRepository) FindGame(id int) (*types.Game, error) {
	gameModel := models.Game{}
	findQuery := "SELECT * FROM awbw_games WHERE games_id = ?"
	err := DB.Get(&gameModel, findQuery, id)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find game with given ID: %s", err.Error()))
	}
	game := types.NewGame(gameModel)
	return &game, nil
}

func (r GameRepository) CreateGame(body types.Game) (int, error) {
	columns := []string{
		"games_name",
		"games_start_date",
		"games_end_date",
		"games_creator",
		"games_activity_date",
		"games_maps_id",
		"games_weather_type",
		"games_weather_start",
		"games_weather_code",
		"games_password",
		"games_turn",
		"games_day",
		"games_active",
		"games_funds",
		"games_capture_win",
		"games_days",
		"games_fog",
		"games_comment",
		"games_type",
		"games_official",
		"games_min_rating",
		"games_unit_limit",
		"games_league",
		"games_team",
	}
	createQuery := FormatCreateQuery("awbw_games", columns)
	res, err := DB.NamedExec(createQuery, body)

	if err != nil {
		return -1, errors.New(fmt.Sprintf("Could not create new Game: %s", err.Error()))
	}

	gameId, err := res.LastInsertId()

	return int(gameId), nil
}

// This function assumes that the given fields can be updated
func (r GameRepository) UpdateGame(id int, updatedFields map[string]interface{}) (map[string]interface{}, error) {
	var updateStatements []string

	for column, _ := range updatedFields {
		updateStatements = append(updateStatements, fmt.Sprintf("%s = :%s", column, column))
	}
	updateQuery := fmt.Sprintf("UPDATE awbw_games SET %s WHERE %s = :%s", strings.Join(updateStatements, ","), gamecolumns.ID, gamecolumns.ID)
	updatedFields[gamecolumns.ID] = id
	_, err := DB.NamedExec(updateQuery, updatedFields)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not update game: %s", err.Error()))
	}

	return updatedFields, nil
}

func (r GameRepository) DeleteGame(id int) error {
	deleteQuery := fmt.Sprintf("DELETE FROM awbw_games WHERE %s = :%s", gamecolumns.ID, gamecolumns.ID)

	params := map[string]interface{}{
		gamecolumns.ID: id,
	}
	_, err := DB.NamedExec(deleteQuery, params)

	if err != nil {
		return errors.New(fmt.Sprintf("Could not delete game with ID %d: %s", id, err.Error()))
	}

	return nil
}
