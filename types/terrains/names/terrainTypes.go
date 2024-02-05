package terrainnames

import "strings"

type TerrainName string

type TerrainStars map[TerrainName]int

var Stars TerrainStars

const (
	Plain    TerrainName = "Plain"
	Mountain TerrainName = "Mountain"
	Wood     TerrainName = "Wood"
	River    TerrainName = "River"
	Road     TerrainName = "Road"
	Bridge   TerrainName = "Bridge"
	Sea      TerrainName = "Sea"
	Shoal    TerrainName = "Shoal"
	Reef     TerrainName = "Reef"
	City     TerrainName = "City"
	Base     TerrainName = "Base"
	Airport  TerrainName = "Airport"
	Port     TerrainName = "Port"
	HQ       TerrainName = "HQ"
	Pipe     TerrainName = "Pipe"
	Seam     TerrainName = "Seam"
	Rubble   TerrainName = "Rubble"
	Silo     TerrainName = "Silo"
	ComTower TerrainName = "ComTower"
	Lab      TerrainName = "Lab"
)

func (t TerrainName) Match(haystack TerrainName) bool {
	return strings.Contains(string(haystack), string(t))
}

func (t TerrainName) Stars() int {
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
