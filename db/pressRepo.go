package db

import (
	"errors"
	"fmt"

	"github.com/awbw/2040/models"
	presscolumns "github.com/awbw/2040/models/columnNames/press"
	presstocolumns "github.com/awbw/2040/models/columnNames/pressTo"
	"github.com/awbw/2040/types"
)

type PressRepository struct{}

var PressRepo PressRepository

func NewPressRepo() PressRepository {
	return PressRepository{}
}

func init() {
	PressRepo = NewPressRepo()
}

func (r *PressRepository) FindPress(id int) (*types.Press, error) {
	var pressModel models.Press
	findQuery := `SELECT awbw_press.*, GROUP_CONCAT(users_username) AS recipients FROM awbw_press, awbw_press_to, awbw_players, ofua_users 
		WHERE press_id = ?
		AND press_to_press_id = press_id
		AND press_to_players_id = players_id
		AND users_id = players_users_id;`
	err := DB.Get(&pressModel, findQuery, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find press with given ID: %s", err.Error()))
	}
	pressType := types.NewPress(pressModel)
	return &pressType, nil
}

func (r *PressRepository) FindRecipients(pressId int) ([]types.PressTo, error) {
	var pressToModels []models.PressTo
	findQuery := `SELECT awbw_press_to.* FROM awbw_press, awbw_press_to, awbw_players, ofua_users
		WHERE press_id = ?
		AND press_id = press_to_press_id
		AND press_to_players_id = players_id
		AND players_users_id = users_id`
	err := DB.Select(&pressToModels, findQuery, pressId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find press recipients with given Press ID (%d): %s", pressId, err.Error()))
	}
	var pressToTypes []types.PressTo
	for _, p := range pressToModels {
		pressToTypes = append(pressToTypes, types.NewPressTo(p))
	}

	return pressToTypes, nil
}

func (r *PressRepository) CreatePress(body types.Press, recipients []int) (int, error) {
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

	tx, err := DB.Beginx()
	pressRes, err := tx.NamedExec(pressQuery, body)
	pressId, err := pressRes.LastInsertId()

	if err != nil {
		return -1, errors.New(fmt.Sprintf("Could not get last inserted ID: %s", err.Error()))
	}

	for _, recipientId := range recipients {
		_, err := tx.NamedExec(pressToQuery, map[string]interface{}{
			presstocolumns.PressID:  pressId,
			presstocolumns.PlayerID: recipientId,
		})
		if err != nil {
			return -1, errors.New(fmt.Sprintf("Could not create Press To: %s", err.Error()))
		}
	}

	err = tx.Commit()

	if err != nil {
		return -1, errors.New(fmt.Sprintf("Could not create new Press: %s", err.Error()))
	}

	return int(pressId), nil
}
