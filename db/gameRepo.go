package db

import (
	"errors"
	"fmt"
	"strings"

	"github.com/awbw/2040/models"
)

type GameRepository struct{}

var GameRepo GameRepository

func NewGameRepo() GameRepository {
	return GameRepository{}
}

func init() {
	GameRepo = NewGameRepo()
}

func (r GameRepository) FindGame(id int) (models.Game, error) {
	var gameModel models.Game
	findQuery := "SELECT * FROM awbw_games WHERE games_id = ?"
	err := DB.Get(&gameModel, findQuery, id)

	if err != nil {
		return models.Game{}, errors.New("Failed to find game with given ID")
	}
	return gameModel, nil
}

func (r GameRepository) CreateGame(body models.Game) (models.Game, error) {
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
	var (
		insertCols strings.Builder
		values     strings.Builder
	)

	// Format columns to insert with the values from the list above to make sure they always match
	for _, col := range columns {
		insertCols.WriteString(fmt.Sprintf("%s,", col))
		values.WriteString(fmt.Sprintf(":%s,", col))

	}
	createQuery := fmt.Sprintf("INSERT INTO awbw_games (%s) VALUES (%s)", insertCols, values)
	_, err := DB.NamedExec(createQuery, body)

	if err != nil {
		return body, errors.New("Could not create new game")
	}

	return body, nil
}

// This function assumes that the given fields can be updated
func (r GameRepository) UpdateGame(body models.Game, updatedFields []string) (models.Game, error) {
	var updateStatements []string

	for _, column := range updatedFields {
		updateStatements = append(updateStatements, fmt.Sprintf("%s = :%s", column, column))
	}

	updateQuery := fmt.Sprintf("UPDATE awbw_games SET %s WHERE games_id = ?", strings.Join(updateStatements, ","))
	_, err := DB.NamedExec(updateQuery, updatedFields)

	if err != nil {
		return body, errors.New("Could not update game")
	}

	return body, nil
}
