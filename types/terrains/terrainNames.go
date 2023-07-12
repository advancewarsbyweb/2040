package terraintypes

import (
	"fmt"

	countrynames "github.com/awbw/2040/types/countries/names"
)

type TerrainNames map[int]string

type TerrainName TerrainType

var Names TerrainNames

func init() {
	Names = NewTerrainNames()
}

func NewTerrainNames() TerrainNames {

	roaddirs := Roaddirs()
	withdirs := Withdirs()

	Names = TerrainNames{
		1:   string(Plain),
		2:   string(Mountain),
		3:   string(Wood),
		26:  fmt.Sprintf("H%s", Bridge),
		27:  fmt.Sprintf("V%s", Bridge),
		28:  string(Sea),
		29:  fmt.Sprintf("H%s", Shoal),
		30:  fmt.Sprintf("H%sN", Shoal),
		31:  fmt.Sprintf("V%s", Shoal),
		32:  fmt.Sprintf("V%sE", Shoal),
		33:  string(Reef),
		34:  fmt.Sprintf("Neutral %s", City),
		35:  fmt.Sprintf("Neutral %s", Base),
		36:  fmt.Sprintf("Neutral %s", Airport),
		37:  fmt.Sprintf("Neutral %s", Port),
		127: FormatBuilding(countrynames.AmberBlaze, ComTower),
		128: FormatBuilding(countrynames.BlackHole, ComTower),
		129: FormatBuilding(countrynames.BlueMoon, ComTower),
		130: FormatBuilding(countrynames.BrownDesert, ComTower),
		131: FormatBuilding(countrynames.GreenEarth, ComTower),
		132: FormatBuilding(countrynames.JadeSun, ComTower),
		133: FormatBuilding(countrynames.Neutral, ComTower),
		134: FormatBuilding(countrynames.OrangeStar, ComTower),
		135: FormatBuilding(countrynames.RedFire, ComTower),
		136: FormatBuilding(countrynames.YellowComet, ComTower),
		137: FormatBuilding(countrynames.GreySky, ComTower),
		138: FormatBuilding(countrynames.AmberBlaze, Lab),
		139: FormatBuilding(countrynames.BlackHole, Lab),
		140: FormatBuilding(countrynames.BlueMoon, Lab),
		141: FormatBuilding(countrynames.BrownDesert, Lab),
		142: FormatBuilding(countrynames.GreenEarth, Lab),
		143: FormatBuilding(countrynames.GreySky, Lab),
		144: FormatBuilding(countrynames.JadeSun, Lab),
		145: FormatBuilding(countrynames.Neutral, Lab),
		146: FormatBuilding(countrynames.OrangeStar, Lab),
		147: FormatBuilding(countrynames.RedFire, Lab),
		148: FormatBuilding(countrynames.YellowComet, Lab),
	}

	// Not every country has the same terrain order
	terrainOrder := []TerrainType{City, Base, Airport, Port, HQ}
	FormatCountries(Names, 38, terrainOrder, countrynames.OrangeStar)
	FormatCountries(Names, 43, terrainOrder, countrynames.BlueMoon)
	FormatCountries(Names, 48, terrainOrder, countrynames.GreenEarth)
	FormatCountries(Names, 53, terrainOrder, countrynames.YellowComet)
	FormatCountries(Names, 81, terrainOrder, countrynames.RedFire)
	FormatCountries(Names, 86, terrainOrder, countrynames.GreySky)
	FormatCountries(Names, 91, terrainOrder, countrynames.BlackHole)
	FormatCountries(Names, 96, terrainOrder, countrynames.BrownDesert)

	terrainOrder = []TerrainType{Airport, Base, City, HQ, Port}
	FormatCountries(Names, 117, terrainOrder, countrynames.AmberBlaze)
	FormatCountries(Names, 122, terrainOrder, countrynames.JadeSun)

	terrainOrder = []TerrainType{Airport, Base, City, ComTower, HQ, Lab, Port}
	FormatCountries(Names, 149, terrainOrder, countrynames.CobaltIce)
	FormatCountries(Names, 156, terrainOrder, countrynames.CobaltIce)
	FormatCountries(Names, 163, terrainOrder, countrynames.CobaltIce)
	FormatCountries(Names, 170, terrainOrder, countrynames.CobaltIce)
	FormatCountries(Names, 181, terrainOrder, countrynames.CobaltIce)
	FormatCountries(Names, 188, terrainOrder, countrynames.CobaltIce)

	for terrainId, terrainName := range withdirs {
		for i, dir := range roaddirs {
			Names[terrainId+i] = fmt.Sprintf("%s%s", dir, terrainName)
		}
	}

	return Names
}

func FormatCountries(tNames TerrainNames, startIdx int, terrainOrder []TerrainType, c countrynames.CountryName) {
	for i, terrainName := range terrainOrder {
		tNames[startIdx+i] = FormatBuilding(c, terrainName)
	}
}

func FormatBuilding(c countrynames.CountryName, t TerrainType) string {
	return fmt.Sprintf("%s %s", c, t)
}

func Roaddirs() []string {
	return []string{
		"H",
		"V",
		"C",
		"ES",
		"SW",
		"WN",
		"NE",
		"ESW",
		"SWN",
		"WNE",
		"NES",
	}
}

func Withdirs() TerrainNames {
	return TerrainNames{
		4:  string(River),
		15: string(Road),
	}
}
