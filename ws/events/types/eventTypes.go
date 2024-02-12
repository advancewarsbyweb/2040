package eventtypes

type EventType string

// Event names received from the client
const (
	MoveRequest         EventType = "moveRequest"
	NotificationRequest EventType = "notificationRequest"
)

// Event names sent back to the client
const (
	MoveResponse         EventType = "moveResponse"
	NotificationResponse EventType = "notificationResponse"
)
