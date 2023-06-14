package terraintypes

import "strings"

type TerrainType string

type TerrainStars map[TerrainType]int

var Stars TerrainStars

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
	Seam     TerrainType = "Seam"
	Silo     TerrainType = "Silo"
	ComTower TerrainType = "ComTower"
	Lab      TerrainType = "Lab"
)

func (t TerrainType) Match(terrainName string) bool {
	return strings.Contains(terrainName, string(t))
}

func (t TerrainType) Stars() int {
	return Stars[t]
}

func init() {
	Stars = TerrainStars{
		Plain:    1,
		Mountain: 4,
		Wood:     2,
		River:    0,
		Road:     0,
		Bridge:   0,
		Sea:      0,
		Shoal:    0,
		Reef:     0,
		City:     3,
		Base:     3,
		Airport:  3,
		Port:     3,
		HQ:       4,
		Pipe:     0,
		Seam:     0,
		Silo:     3,
		ComTower: 3,
		Lab:      3,
	}
}
