package terraintypes

import (
	"github.com/awbw/2040/db"
	"github.com/awbw/2040/models"
)

type TerrainMap map[int]map[int]models.Tile

func CreateTerrainMap(mapId int) (*TerrainMap, error) {

	mapTiles, err := db.MapRepo.FindMapTiles(mapId)
	if err != nil {
		return nil, err
	}

	terrainMap := make(TerrainMap)

	for _, t := range *mapTiles {
		if terrainMap[t.X] == nil {
			terrainMap[t.X] = make(map[int]models.Tile)
		}
		terrainMap[t.X][t.Y] = t
	}

	return &terrainMap, nil
}
