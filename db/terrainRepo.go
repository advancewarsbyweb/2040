package db

import (
	"errors"
	"fmt"

	models "github.com/awbw/2040/models"
)

type MapRepository struct{}

var MapRepo MapRepository

func NewMapRepo() MapRepository {
	return MapRepository{}
}

func init() {
	MapRepo = NewMapRepo()
}

func (r *MapRepository) FindMapTiles(mapId int) ([]models.Tile, error) {

	var mapTiles []models.Tile
	findQuery := `SELECT awbw_tiles.tiles_x, awbw_tiles.tiles_y, awbw_terrain.* FROM awbw_tiles, awbw_terrain
		WHERE tiles_maps_id = ?
		AND tiles_terrain_id = terrain_id`
	err := DB.Get(mapTiles, findQuery, mapId)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to find map tiles with given id: %d", mapId))
	}
	return mapTiles, nil
}
