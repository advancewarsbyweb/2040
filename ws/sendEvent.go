package ws

import (
	"github.com/awbw/2040/models"
)

// Send event to users
func SendEventResponse(outgoingEvent Event, users []models.User) {
	for _, user := range users {
		if connectedClient := ClientManager.Clients[user.ID]; connectedClient != nil {
			connectedClient.Egress <- outgoingEvent
		}
	}
}
