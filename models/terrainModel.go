package models

import (
	countrycodes "github.com/awbw/2040/types/countries/codes"
	terrainnames "github.com/awbw/2040/types/terrains/names"
)

type Terrain struct {
	ID          int
	Name        terrainnames.TerrainName
	Defense     int
	Symbol      string
	CountryCode countrycodes.CountryCode
	BuildType   string
	Offset      string
}
