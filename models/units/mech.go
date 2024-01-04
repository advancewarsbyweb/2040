package models

type mech struct {
	directUnit
}

func (u *mech) Load() {
}

func NewMech() BaseUnit {
	return &mech{
		directUnit{},
	}
}
