package ws

import (
	"time"

	"github.com/awbw/2040/models"
	eventtypes "github.com/awbw/2040/ws/events/types"
)

type Event struct {
	Type      eventtypes.EventType `json:"type"`
	Timestamp time.Time            `json:"timestamp"`
	Users     []models.User        `json:"-"`
	Data      interface{}          `json:"data"`
}

type EventHandler func(event Event) error
