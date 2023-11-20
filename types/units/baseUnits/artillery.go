package unittypes

type artillery struct {
	indirectUnit
}

func NewArtillery() BaseUnit {
	return &artillery{
		indirectUnit{},
	}
}
