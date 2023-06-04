package terraintypes

import (
	"github.com/awbw/2040/db"
	"github.com/awbw/2040/models"
)

type TerrainMap map[int]map[int]models.Tile

func NewTerrainMap(mapId int) *TerrainMap {
	mapTiles, _ := MapTiles(mapId)
	return CreateTerrainMap(mapTiles)
}

func MapTiles(mapId int) (*[]models.Tile, error) {
	mapTiles, err := db.MapRepo.FindMapTiles(mapId)
	if err != nil {
		return nil, err
	}
	return mapTiles, nil
}

func CreateTerrainMap(mapTiles *[]models.Tile) *TerrainMap {

	terrainMap := make(TerrainMap)

	for _, t := range *mapTiles {
		if terrainMap[t.X] == nil {
			terrainMap[t.X] = make(map[int]models.Tile)
		}
		terrainMap[t.X][t.Y] = t
	}

	return &terrainMap
}
