package db

import (
	"errors"
	"fmt"

	models "github.com/awbw/2040/models"
)

type UnitRepository struct {
	Columns []string
}

var (
	UnitRepo UnitRepository

	unitRelationsQuery = `SELECT u.* FROM awbw_units AS u
        INNER JOIN awbw_players AS p
            ON p.players_id = u.units_players_id
        INNER JOIN awbw_games AS g
            ON g.games_id = u.units_games_id
        INNER JOIN awbw_tiles AS ti
            ON ti.tiles_maps_id = g.games_maps_id
            AND ti.tiles_x = u.units_x
            AND ti.tiles_y = u.units_y
        INNER JOIN awbw_terrain AS te
            ON ti.tiles_terrain_id = te.terrain_id
        %s`
)

func NewUnitRepo() UnitRepository {
	return UnitRepository{
		Columns: []string{
			"units_id",
			"units_players_id",
			"units_games_id",
			"units_x",
			"units_y",
			"units_sub_dive",
			"units_moved",
			"units_fired",
			"units_hp",
			"units_cargo1_units_id",
			"units_cargo2_units_id",
			"units_carried",
			"units_name",
			"units_movement_points",
			"units_movement_type",
			"units_vision",
			"units_fuel",
			"units_fuel_per_turn",
			"units_short_range",
			"units_long_range",
			"units_cost",
			"units_ammo",
		},
	}
}

func init() {
	UnitRepo = NewUnitRepo()
}

func (r *UnitRepository) FindUnit(id int) (*models.Unit, error) {
	var u models.Unit
	query := "SELECT * FROM awbw_units WHERE players_id = ?"

	err := DB.Get(&u, query, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find unit with given ID: %s", err.Error()))
	}

	return &u, nil
}

// Relations are the Player, Game and Tile
func (r *UnitRepository) FindUnitRelations(id int) (*models.Unit, error) {
	var u models.Unit
	query := fmt.Sprintf(unitRelationsQuery, "WHERE units_id = ?")

	err := DB.Get(&u, query, id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find unit relations with given ID: %s", err.Error()))
	}

	return &u, nil
}

// Relations are the Player, Game and Tile
func (r *UnitRepository) FindUnitsRelationsByGame(gameId int) ([]models.Unit, error) {
	var u []models.Unit
	query := fmt.Sprintf(unitRelationsQuery, "WHERE units_games_id = ?")

	err := DB.Select(&u, query, gameId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find units relations with given game ID: %s", err.Error()))
	}

	return u, nil
}

func (r *UnitRepository) CreateUnit(body models.Unit) (int, error) {
	query := FormatCreateQuery("awbw_units", UserRepo.Columns)
	res, err := DB.NamedExec(query, body)
	if err != nil {
		return -1, errors.New(fmt.Sprintf("Could not create new Unit: %s", err.Error()))
	}

	userId, err := res.LastInsertId()

	return int(userId), nil
}
