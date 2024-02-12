package ws

import (
	"time"

	"github.com/awbw/2040/db"
	eventdata "github.com/awbw/2040/ws/events/data"
	eventtypes "github.com/awbw/2040/ws/events/types"
)

func MoveHandler(event Event) error {
	// Validate action

	// Format response
	// MoveResponse would be the response regardless of vision for individual players
	// This would be saved in the database
	_ = eventdata.MoveUnformated{}

	// Here we format a new MoveResponse for individual users if they are connected to the game
	playerModels, err := db.PlayerRepo.FindPlayersByGame(1)
	if err != nil {
		return err
	}

	for _, playerModel := range playerModels {
		connectedUser := ClientManager.Clients[playerModel.UserID]
		if connectedUser == nil {
			continue
		}

		// For each user we format the response. Utils function to format to each specific user

		_ = Event{
			Type:      eventtypes.MoveResponse,
			Timestamp: time.Now(),
			Data:      eventdata.MoveResponse{},
		}

	}

	return nil
}
