package countrycodes

import countrynames "github.com/awbw/2040/types/countries/names"

type CountryCode string

type CountryCodes map[countrynames.CountryName]CountryCode

var Codes CountryCodes

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
	Codes = CountryCodes{
		countrynames.OrangeStar:     OrangeStar,
		countrynames.BlueMoon:       BlueMoon,
		countrynames.GreenEarth:     GreenEarth,
		countrynames.YellowComet:    YellowComet,
		countrynames.BlackHole:      BlackHole,
		countrynames.RedFire:        RedFire,
		countrynames.GreySky:        GreySky,
		countrynames.BrownDesert:    BrownDesert,
		countrynames.AmberBlaze:     AmberBlaze,
		countrynames.JadeSun:        JadeSun,
		countrynames.CobaltIce:      CobaltIce,
		countrynames.PinkCosmos:     PinkCosmos,
		countrynames.TealGalaxy:     TealGalaxy,
		countrynames.PurpleLighting: PurpleLighting,
		countrynames.AcidRain:       AcidRain,
		countrynames.WhiteNova:      WhiteNova,
	}
}
