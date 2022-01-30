package response

import (
	"messages_api/internal/messages/model"
	"time"
)

// Message object used for responses.
type Message struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Text         string    `json:"text"`
	CreationDate time.Time `json:"creation_date"`
}

// NewMessageResponseFromModel factory for creating a Message response object from a model.
func NewMessageResponseFromModel(message *model.Message) *Message {
	return &Message{
		ID:           message.ID,
		Name:         message.Name,
		Email:        message.Email,
		Text:         message.Text,
		CreationDate: message.CreationDate,
	}
}
