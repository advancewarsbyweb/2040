package eventdata

import "github.com/awbw/2040/models"

type Team string

type TileType string

type PathTile struct {
	UnitVisible bool `json:"unitVisible"`
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
	Paths      map[Team][]PathTile
	Dist       int
	Trapped    bool
	Discovered map[Team][]Discovereds
}

type MoveRequest struct {
	PlayerID int   `json:"playerId"`
	UnitID   int   `json:"unitId"`
	Path     []int `json:"path"`
}

type MoveResponse struct {
	Unit        models.Unit   `json:"unit"`
	Path        []PathTile    `json:"path,omitempty"`
	Discovereds []Discovereds `json:"discoveredUnits,omitempty"`
	Trapped     bool          `json:"trapped"`
}
