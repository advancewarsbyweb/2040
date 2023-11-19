package unittypes

type infantry struct {
	directUnit
}

func (u *infantry) Load() {
}

func NewInfantry() BaseUnit {
	return &infantry{
		directUnit{},
	}
}
