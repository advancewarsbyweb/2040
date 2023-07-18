package models

type Unit struct {
	ID       int
	PlayerID int
	GameID   int
	X        int
	Y        int
	SubDive  string
	Moved    int
	Fired    int
	HP       float64
	Cargo1ID int
	Cargo2ID int
	Carried  string
	BaseUnit
}
