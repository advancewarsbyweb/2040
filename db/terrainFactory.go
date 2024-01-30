package db

import "github.com/awbw/2040/models/terrains"

type terrainFactory struct {
	Terrain models.Terrain
}

var TerrainFactory terrainFactory

func NewTerrainFactory() terrainFactory {
	return terrainFactory{}
}

func init() {
	TerrainFactory = NewTerrainFactory()
}
