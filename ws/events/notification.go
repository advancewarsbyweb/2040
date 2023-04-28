package events

import "time"

// A User joined a waiting-to-start game
type GameJoinNotification struct {
	GameID int       `json:"gameId"`
	Sent   time.Time `json:"sent"`
}

// Notification sent back to the client
type NewNotification struct {
	Message string `json:"message"`
	Url     string `json:"url,omitempty"` // The url when clicking on the notification. Can be present or absent
}
