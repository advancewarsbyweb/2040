package ws

import (
	"time"

	eventtypes "github.com/awbw/2040/ws/events/types"
)

// This handler is typically fired by the backend only
func NotificationHandler(event Event) error {

	event = Event{
		Type:      eventtypes.NotificationResponse,
		Timestamp: time.Now(),
		Data:      event.Data,
	}

	return nil
}
