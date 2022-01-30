package model

import (
	"github.com/google/uuid"
	"messages_api/internal/data/request"
	"strings"
	"time"
)

// Message model
type Message struct {
	ID           string
	Name         string
	Email        string
	Text         string
	CreationDate time.Time
}

// NewMessageFromRequest returns a model.Message instance based on the request.CreateMessage request.
func NewMessageFromRequest(message *request.CreateMessage) *Message {
	return &Message{
		Name:  message.Name,
		Email: message.Email,
		Text:  message.Text,
	}
}

// Init the ID with a random UUID and the CreationDate with the current time.
func (m *Message) Init() {
	m.ID = strings.ToUpper(uuid.NewString())
	m.CreationDate = time.Now()
}

// Identifier returns the identifier of the message
func (m *Message) Identifier() string {
	return m.ID
}
