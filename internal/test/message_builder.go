package test

import (
	"messages_api/internal/messages/model"
	"time"
)

type messageBuilder struct {
	m *model.Message
}

func NewMessageBuilder(id string) *messageBuilder {
	return &messageBuilder{
		m: &model.Message{
			ID:           id,
			Name:         "some-name",
			Email:        "email@gmail.com",
			Text:         "some-text",
			CreationDate: MustParseTimeFromFormat(time.RFC3339, "2022-01-01T00:00:00Z"),
		},
	}
}

func (b *messageBuilder) WithText(s string) *messageBuilder {
	b.m.Text = s
	return b
}

func (b *messageBuilder) WithCreationDate(t time.Time) *messageBuilder {
	b.m.CreationDate = t
	return b
}

func (b *messageBuilder) Build() *model.Message {
	return b.m
}
