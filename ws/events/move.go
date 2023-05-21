package events

import "github.com/awbw/2040/models"

type Team string

type PathTile struct {
	UnitVisible bool `json:"unit_visible"`
	X           int  `json:"x"`
	Y           int  `json:"y"`
}

type Discovereds struct {
	Units     []models.Unit
	Buildings []models.Building
}

// Information for all players. This needs to be formatted for individual players
type MoveUnformated struct {
	Unit       models.Unit
	Paths      map[Team][]PathTile    `json:"paths"`
	Dist       int                    `json:"dist"`
	Trapped    bool                   `json:"trapped"`
	Discovered map[Team][]Discovereds `json:"discovered"`
}

type MoveRequest struct {
	PlayerID int   `json:"players_id"`
	UnitID   int   `json:"units_id"`
	Path     []int `json:"path"`
}

type MoveResponse struct {
	Unit        models.Unit   `json:"unit"`
	Path        []PathTile    `json:"path"`
	Discovereds []Discovereds `json:"discovered_units,omitempty"`
	Trapped     bool          `json:"trapped"`
}
