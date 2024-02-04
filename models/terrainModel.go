package models

import (
	countrycodes "github.com/awbw/2040/types/countries/codes"
	terraintypes "github.com/awbw/2040/types/terrains"
)

type Terrain struct {
	ID          int
	Name        terraintypes.TerrainName
	Defense     int
	Symbol      string
	CountryCode countrycodes.CountryCode
	BuildType   string
	Offset      string
}
