package countrycodes

import countrynames "github.com/awbw/2040/types/countries/names"

type CountryCode string

var CountryCodes map[countrynames.CountryName]CountryCode

const (
	OrangeStar     CountryCode = "OS"
	BlueMoon       CountryCode = "BM"
	GreenEarth     CountryCode = "GE"
	YellowComet    CountryCode = "YC"
	BlackHole      CountryCode = "BH"
	RedFire        CountryCode = "RF"
	GreySky        CountryCode = "GS"
	BrownDesert    CountryCode = "BD"
	AmberBlaze     CountryCode = "AB"
	JadeSun        CountryCode = "JS"
	CobaltIce      CountryCode = "CI"
	PinkCosmos     CountryCode = "PC"
	TealGalaxy     CountryCode = "TG"
	PurpleLighting CountryCode = "PL"
	AcidRain       CountryCode = "AR"
	WhiteNova      CountryCode = "WN"
)

func init() {
	CountryCodes = map[countrynames.CountryName]CountryCode{
		countrynames.OrangeStar:  OrangeStar,
		countrynames.BlueMoon:    BlueMoon,
		countrynames.GreenEarth:  GreenEarth,
		countrynames.YellowComet: YellowComet,
		countrynames.BlackHole:   BlackHole,
	}
}
