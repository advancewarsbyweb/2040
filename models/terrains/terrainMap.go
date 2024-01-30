package models

type TerrainMap map[int]map[int]Tile

func CreateTerrainMap(mapTiles []Tile) TerrainMap {
	terrainMap := make(TerrainMap)

	for _, t := range mapTiles {
		if terrainMap[t.X] == nil {
			terrainMap[t.X] = make(map[int]Tile)
		}
		terrainMap[t.X][t.Y] = t
	}

	return terrainMap
}
