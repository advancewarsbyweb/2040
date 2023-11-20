package unittypes

type tank struct {
	directUnit
}

func NewTank() BaseUnit {
	return &tank{
		directUnit{},
	}
}
