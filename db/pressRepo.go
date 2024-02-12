package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/awbw/2040/models"
	presscolumns "github.com/awbw/2040/models/columnNames/press"
	presstocolumns "github.com/awbw/2040/models/columnNames/pressTo"
)

type PressRepository struct{}

var PressRepo PressRepository

func NewPressRepo() PressRepository {
	return PressRepository{}
}

func init() {
	PressRepo = NewPressRepo()
}

func (r *PressRepository) FindPress(id int) (*models.Press, error) {
	var pressModel models.Press
	query := `SELECT p.*, GROUP_CONCAT(r.users_username SEPARATOR ',') AS recipients FROM awbw_press AS p
        INNER JOIN (
            SELECT pt.press_to_press_id, u.users_username FROM awbw_press_to AS pt
            INNER JOIN awbw_players AS pl
                ON pl.players_id = pt.press_to_players_id
            INNER JOIN ofua_users AS u
                ON u.users_id = pl.players_users_id
        ) AS r
            ON r.press_to_press_id = p.press_id
        WHERE p.press_id = ?
        GROUP BY p.press_id`
	err := DB.Get(&pressModel, query, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find press with given ID: %s", err.Error()))
	}
	return &pressModel, nil
}

func (r *PressRepository) FindRecipients(pressId int) ([]models.PressTo, error) {
	var pressToModels []models.PressTo
	findQuery := `SELECT pt.* FROM awbw_press AS p, awbw_press_to AS pt, awbw_players AS pl, ofua_users as u
		WHERE p.press_id = ?
		AND p.press_id = pt.press_to_press_id
		AND pt.press_to_players_id = pl.players_id
		AND pl.players_users_id = u.users_id`
	err := DB.Select(&pressToModels, findQuery, pressId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find press recipients with given Press ID (%d): %s", pressId, err.Error()))
	}
	return pressToModels, nil
}

func (r *PressRepository) FindAllPress(playerId int) ([]models.Press, error) {
	var pressModels []models.Press
	query := `SELECT p.*, GROUP_CONCAT(r.users_username SEPARATOR ',') AS recipients FROM awbw_press AS p
        INNER JOIN (
            SELECT pt.press_to_press_id, u.users_username FROM awbw_press_to AS pt
            INNER JOIN awbw_players AS pl
                ON pl.players_id = pt.press_to_players_id
            INNER JOIN ofua_users AS u
                ON u.users_id = pl.players_users_id
        ) AS r
            ON r.press_to_press_id = p.press_id
        WHERE p.press_players_id = ?
        GROUP BY p.press_id`

	err := DB.Select(&pressModels, query, playerId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find player's press with given Player ID (%d): %s", playerId, err.Error()))
	}
	return pressModels, nil
}

func (r *PressRepository) CreatePress(body models.Press, recipients []int) (int, error) {
	pressColumns := []string{
		presscolumns.Text,
		presscolumns.PlayerID,
		presscolumns.Subject,
		presscolumns.Date,
		presscolumns.Type,
	}
	pressToColumns := []string{
		presstocolumns.PressID,
		presstocolumns.PlayerID,
	}
	pressQuery := FormatCreateQuery("awbw_press", pressColumns)
	pressToQuery := FormatCreateQuery("awbw_press_to", pressToColumns)

	body.Date = time.Now()
	tx, err := DB.Beginx()
	pressRes, err := tx.NamedExec(pressQuery, body)
	if err != nil {
		fmt.Printf("%v", err)
		return -1, errors.New(fmt.Sprintf("Could not create Press: %s", err.Error()))
	}
	pressId, err := pressRes.LastInsertId()

	if err != nil {
		fmt.Printf("%v", err)
		return -1, errors.New(fmt.Sprintf("Could not get last inserted ID: %s", err.Error()))
	}

	for _, recipientId := range recipients {
		_, err := tx.NamedExec(pressToQuery, map[string]interface{}{
			presstocolumns.PressID:  pressId,
			presstocolumns.PlayerID: recipientId,
		})
		if err != nil {
			fmt.Printf("%v", err)
			return -1, errors.New(fmt.Sprintf("Could not create Press To: %s", err.Error()))
		}
	}

	err = tx.Commit()

	if err != nil {
		return -1, errors.New(fmt.Sprintf("Could not commit transaction to create new Press: %s", err.Error()))
	}

	return int(pressId), nil
}
