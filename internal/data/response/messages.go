package response

import (
	"messages_api/internal/messages/model"
)

// Messages response object.
// The NextPage contains the full URL to access the next batch of messages, or it is empty if there is no other page.
// TotalCount is the amount of messages
type Messages struct {
	Message    []Message `json:"message"`
	NextPage   string    `json:"next_page,omitempty"`
	TotalCount uint64    `json:"total_count"`
}

// NewMessagesResponseFromModel factory for creating a Messages response object.
func NewMessagesResponseFromModel(messages []model.Message, nextPage string, totalCount uint64) *Messages {
	messageSlice := make([]Message, len(messages))
	for i := 0; i < len(messages); i++ {
		message := NewMessageResponseFromModel(&messages[i])
		messageSlice[i] = *message
	}

	return &Messages{
		Message:    messageSlice,
		NextPage:   nextPage,
		TotalCount: totalCount,
	}
}
