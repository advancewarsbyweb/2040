package ws

import (
	"encoding/json"

	"github.com/awbw/2040/db"
	"github.com/awbw/2040/ws/events"
)

func MoveHandler(moveRequest events.MoveRequest, c *Client) error {
	// Validate action

	// Format response
	// MoveResponse would be the response regardless of vision for individual players
	// This would be saved in the database
	_ = events.MoveUnformated{}

	// Here we format a new MoveResponse for individual users if they are connected to the game
	playerModels, err := db.PlayerRepo.FindPlayersByGame(1)
	if err != nil {
		return err
	}

	for _, playerModel := range playerModels {
		connectedUser := c.Manager.Clients[playerModel.UserID]
		if connectedUser == nil {
			continue
		}

		// For each user we format the response. Utils function to format to each specific user

		jsonResponse, err := json.Marshal(events.MoveResponse{})
		if err != nil {
			return err
		}

		_ = Event{
			Payload: jsonResponse,
			Type:    MoveResponse,
		}

	}

	return nil
}
