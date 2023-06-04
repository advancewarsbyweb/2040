package ws

import (
	"encoding/json"
)

type Event struct {
	Type    EventType       `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

type EventType string

// Event names received from the client
const (
	MoveRequest         EventType = "move_request"
	NotificationRequest EventType = "notification_request"
)

// Event names sent back to the client
const (
	MoveResponse         EventType = "move_response"
	NotificationResponse EventType = "notification_response"
)
