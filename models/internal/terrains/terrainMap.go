package models

import terraintypes "github.com/awbw/2040/types/terrains"

type TerrainMap map[int]map[int]Tile

type testMaps struct{}

var TestMaps testMaps

// Create from Tiles received from the database
func CreateTerrainMap(mapTiles []Tile) TerrainMap {
	terrainMap := make(TerrainMap)

	for _, t := range mapTiles {
		if terrainMap[t.Y] == nil {
			terrainMap[t.Y] = make(map[int]Tile)
		}
		terrainMap[t.Y][t.X] = t
	}

	return terrainMap
}

func (m *testMaps) Tile(id int) Tile {
	return Tile{
		Terrain: Terrain{
			ID:   id,
			Name: terraintypes.Names[id],
		},
	}
}

func (m *testMaps) FromTiles(tileIDs [][]int) []Tile {
	var mapTiles []Tile
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
	return mapTiles
}

func (m *testMaps) ShangriLa() []Tile {
	tileIDs := [][]int{
		{3, 112, 3, 5, 1, 1, 1, 1, 34, 16, 2, 2, 1, 10, 8, 1, 1, 34, 10, 8, 2, 34, 1, 32, 28},
		{112, 47, 112, 26, 3, 36, 1, 1, 3, 21, 19, 1, 1, 3, 133, 1, 18, 20, 1, 10, 8, 1, 1, 1, 30},
		{3, 112, 3, 5, 1, 1, 2, 1, 1, 1, 21, 15, 15, 15, 26, 15, 20, 1, 3, 18, 26, 19, 3, 1, 34},
		{4, 27, 4, 9, 3, 1, 1, 34, 1, 1, 3, 34, 1, 3, 10, 8, 2, 1, 18, 20, 34, 21, 19, 1, 1},
		{2, 3, 1, 1, 1, 3, 1, 1, 1, 1, 18, 15, 15, 19, 34, 5, 2, 1, 44, 1, 10, 4, 27, 27, 4},
		{1, 44, 19, 34, 1, 29, 29, 2, 34, 1, 16, 1, 1, 21, 15, 5, 15, 19, 34, 1, 1, 3, 1, 1, 1},
		{1, 1, 16, 1, 1, 28, 28, 28, 3, 3, 27, 8, 1, 34, 1, 5, 2, 21, 44, 1, 1, 1, 3, 1, 1},
		{34, 3, 16, 1, 34, 2, 30, 30, 1, 18, 20, 10, 8, 1, 32, 31, 1, 2, 1, 3, 1, 34, 2, 1, 34},
		{1, 18, 20, 3, 2, 1, 3, 1, 18, 20, 1, 7, 13, 8, 32, 31, 34, 3, 1, 1, 34, 2, 2, 18, 15},
		{1, 16, 1, 2, 3, 1, 1, 34, 20, 1, 7, 12, 145, 14, 9, 1, 18, 34, 1, 1, 3, 2, 1, 16, 1},
		{15, 20, 2, 2, 34, 1, 1, 3, 34, 32, 31, 10, 11, 9, 1, 18, 20, 1, 3, 1, 2, 3, 18, 20, 1},
		{34, 1, 2, 34, 1, 3, 1, 2, 1, 32, 31, 1, 10, 8, 18, 20, 1, 29, 29, 2, 34, 1, 16, 3, 34},
		{1, 1, 3, 1, 1, 1, 39, 19, 2, 5, 1, 34, 1, 10, 27, 3, 3, 28, 28, 28, 1, 1, 16, 1, 1},
		{1, 1, 1, 3, 1, 1, 34, 21, 15, 5, 15, 19, 1, 1, 16, 1, 34, 2, 30, 30, 1, 34, 21, 39, 1},
		{4, 27, 27, 4, 8, 1, 39, 1, 2, 5, 34, 21, 15, 15, 20, 1, 1, 1, 1, 3, 1, 1, 1, 3, 2},
		{1, 1, 21, 19, 34, 18, 20, 1, 2, 10, 8, 3, 1, 34, 3, 1, 1, 34, 1, 1, 3, 7, 4, 27, 4},
		{34, 1, 3, 21, 26, 20, 3, 1, 18, 15, 26, 15, 15, 15, 19, 1, 1, 1, 2, 1, 1, 5, 3, 112, 3},
		{29, 1, 1, 1, 10, 8, 1, 18, 20, 1, 133, 3, 1, 1, 21, 19, 3, 1, 1, 36, 3, 26, 112, 42, 112},
		{28, 31, 1, 34, 2, 10, 8, 34, 1, 1, 10, 8, 1, 2, 2, 16, 34, 1, 1, 1, 1, 5, 3, 112, 3},
	}
	return TestMaps.FromTiles(tileIDs)
}
