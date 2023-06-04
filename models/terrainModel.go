package models

import (
	countrycodes "github.com/awbw/2040/types/countries/codes"
)

type Terrain struct {
	ID          int
	Name        string
	Defense     int
	Symbol      string
	CountryCode countrycodes.CountryCode
	BuildType   string
	Offset      string
}
