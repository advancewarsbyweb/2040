package ws

import (
	"encoding/json"
	"fmt"

	"github.com/awbw/2040/models"
	"github.com/awbw/2040/ws/events"
)

// This handler is typically fired by the backend only
func NotificationHandler(notification events.NotificationResponse, sendTo []models.User) error {

	data, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("Failed to marshal notification response")
	}

	var outgoingEvent Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = NotificationResponse

	SendEventResponse(outgoingEvent, sendTo)

	return nil
}
