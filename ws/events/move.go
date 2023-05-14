package events

import "github.com/awbw/2040/models"

type MoveRequest struct {
	PlayerID int   `json:"players_id"`
	UnitID   int   `json:"units_id"`
	Path     []int `json:"path"`
}

type MoveResponse struct {
	Unit            models.Unit   `json:"unit"`
	Path            []int         `json:"path"`
	DiscoveredUnits []models.Unit `json:"discovered_units,omitempty"`
	Trapped         bool          `json:"trapped"`
}
