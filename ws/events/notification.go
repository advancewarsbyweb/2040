package events

import "time"

// A User joined a waiting-to-start game
type GameJoinNotification struct {
	GameID  int       `json:"gameId"`
	Message string    `json:"message"`
	Sent    time.Time `json:"sent"`
}

type NewNotification struct {
}
