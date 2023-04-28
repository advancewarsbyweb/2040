package ws

import (
	"encoding/json"
	"fmt"

	"github.com/awbw/2040/models"
	"github.com/awbw/2040/ws/events"
)

func SendNotificationHandler(notification events.NewNotification, sendTo []models.User) error {

	data, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("Failed to marshal notification")
	}

	var outgoingEvent Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = SendNotification

	c.Egress <- outgoingEvent

	return nil
}
