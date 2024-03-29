package ws

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/awbw/2040/utils/auth"
	eventtypes "github.com/awbw/2040/ws/events/types"
	"github.com/gin-gonic/gin"
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
	Handlers map[eventtypes.EventType]EventHandler
}

func init() {
	ClientManager = NewManager()
}

func NewManager() *Manager {
	m := &Manager{
		Clients:  make(ClientList),
		Handlers: make(map[eventtypes.EventType]EventHandler),
	}
	m.SetupEventHandlers()
	return m
}

// Add every event Handler to the Manager
func (m *Manager) SetupEventHandlers() {
	m.Handlers[eventtypes.NotificationResponse] = NotificationHandler
}

// Make sure the event sent by the client is routed to the proper event handler
func (m *Manager) RouteEvent(e Event, c *Client) error {
	if eventHandler, ok := m.Handlers[e.Type]; ok {
		if err := eventHandler(e); err != nil {
			return nil
		}

		return nil
	} else {
		return ErrEventNotSupported
	}
}

func (m *Manager) ServeWS(c *gin.Context) {
	loggedUser, err := auth.RequireAuth(c)
	if err != nil {
		log.Println(fmt.Sprintf("Could not authorize user: %s", err.Error()))
		return
	}

	conn, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, m, loggedUser.ID)
	m.Subscribe(client)
}

func (m *Manager) Subscribe(client *Client) {
	m.Lock()
	defer m.Unlock()

	// Get the UserID

	m.Clients[client.UserID] = client
}

func (m *Manager) Delete(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.Clients[client.UserID]; ok {
		client.Connection.Close()
		delete(m.Clients, client.UserID)
	}
}

func (m *Manager) Publish(event Event) {
	for _, u := range event.Users {
		if connectedClient := ClientManager.Clients[u.ID]; connectedClient != nil {
			connectedClient.Egress <- event
		}
	}
}
