package terraintypes

import "strings"

type TerrainType string

const (
	Plain    TerrainType = "Plain"
	Mountain TerrainType = "Mountain"
	Wood     TerrainType = "Wood"
	River    TerrainType = "River"
	Road     TerrainType = "Road"
	Bridge   TerrainType = "Bridge"
	Sea      TerrainType = "Sea"
	Shoal    TerrainType = "Shoal"
	Reef     TerrainType = "Reef"
	City     TerrainType = "City"
	Base     TerrainType = "Base"
	Airport  TerrainType = "Airport"
	Port     TerrainType = "Port"
	HQ       TerrainType = "HQ"
	Pipe     TerrainType = "Pipe"
	Silo     TerrainType = "Silo"
	ComTower TerrainType = "ComTower"
	Lab      TerrainType = "Lab"
)

func (t TerrainType) Match(terrainName string) bool {
	return strings.Contains(terrainName, string(t))
}
