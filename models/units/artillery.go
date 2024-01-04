package models

type artillery struct {
	indirectUnit
}

func NewArtillery() Unit {
	return &artillery{
		indirectUnit{},
	}
}
