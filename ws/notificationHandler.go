package ws

import (
	"encoding/json"
	"fmt"

	"github.com/awbw/2040/ws/events"
)

func SendNotificationHandler(e Event, c *Client) error {
	var notifEvent events.GameJoinNotification
	if err := json.Unmarshal(e.Payload, &notifEvent); err != nil {
		return fmt.Errorf("Bad payload in request: %v", err)
	}

	data, err := json.Marshal(notifEvent)
	if err != nil {
		return fmt.Errorf("Failed to marshal notification")
	}

	var outgoingEvent Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = GameJoinNotification

	c.Egress <- outgoingEvent

	return nil
}
