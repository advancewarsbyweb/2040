package eventdata

// e.g: Type: "GameStart" ContextType: "Game", ContextID: gameID
type NotificationResponse struct {
	Type        string `json:"type"`
	ContextType string `json:"contextType"`
	ContextID   int    `json:"contextId"`
}
