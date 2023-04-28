package ws

import (
	"encoding/json"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

// Event names received from the client
const (
	EventSendMessage = "send_message"
	SendNotification = "send_notification"
)

// Event names sent back to the client
const (
	NewNotification = "new_notification"
)
