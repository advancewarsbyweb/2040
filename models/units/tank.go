package models

type tank struct {
	directUnit
}

func NewTank() Unit {
	return &tank{
		directUnit{},
	}
}
