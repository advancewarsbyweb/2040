package models

import (
	"testing"

	terraintypes "github.com/awbw/2040/types/terrains"
)

func TestCreateTerrainMap(t *testing.T) {
	var (
		tileIDs  [][]int
		mapTiles []Tile
	)

	tileIDs = [][]int{
		{1, 2, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	for y, row := range tileIDs {
		for x, id := range row {
			mapTiles = append(mapTiles, Tile{
				X: x,
				Y: y,
				Terrain: Terrain{
					ID:   id,
					Name: terraintypes.Names[id],
				},
			})
		}
	}

	terrainMap := CreateTerrainMap(mapTiles)

	if terrainMap[1][1].Name != string(terraintypes.Plain) {

	}

}
