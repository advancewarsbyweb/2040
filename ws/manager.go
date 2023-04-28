package ws

import (
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ErrEventNotSupported = errors.New("This event type is not supported")
	ClientManager        *Manager
)

// Keeps track of every connected client
type Manager struct {
	Clients ClientList
	sync.RWMutex
	Handlers map[string]EventHandler
}

func init() {
	ClientManager = NewManager()
}

func NewManager() *Manager {
	m := &Manager{
		Clients:  make(ClientList),
		Handlers: make(map[string]EventHandler),
	}
	m.SetupEventHandlers()
	return m
}

// Add every event Handler to the Manager
func (m *Manager) SetupEventHandlers() {
	m.Handlers[SendNotification] = SendNotificationHandler
}

// Make sure the event sent by the client is routed to the proper event handler
func (m *Manager) RouteEvent(e Event, c *Client) error {
	if eventHandler, ok := m.Handlers[e.Type]; ok {
		if err := eventHandler(e, c); err != nil {
			return nil
		}

		return nil
	} else {
		return ErrEventNotSupported
	}
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, m)
	m.AddClient(client)

}

func (m *Manager) AddClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	// Get the UserID

	m.Clients[0] = client
}

func (m *Manager) RemoveClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.Clients[0]; ok {
		client.Connection.Close()
		delete(m.Clients, 0)
	}
}
