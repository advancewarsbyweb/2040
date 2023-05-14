package events

type NotificationRequest struct {
}

// Notification sent back to the client
type NotificationResponse struct {
	Message string `json:"message"`
	Url     string `json:"url,omitempty"` // The url when clicking on the notification. Can be present or absent
}
