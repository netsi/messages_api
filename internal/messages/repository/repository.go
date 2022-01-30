package repository

import (
	"context"
	"errors"
	"messages_api/internal/messages/helper"
	"messages_api/internal/messages/model"
	"sort"
	"sync"
	"sync/atomic"
)

// Repository defines the functions provided by a message repository.
type Repository interface {
	Count(ctx context.Context) (uint64, error)
	Store(ctx context.Context, message *model.Message) error
	UpdateText(_ context.Context, id, text string) (*model.Message, error)
	FetchByID(ctx context.Context, id string) (*model.Message, error)
	FetchMessages(ctx context.Context, offset, limit uint64) ([]model.Message, error)
}

type inMemoryRepository struct {
	count    uint64
	messages sync.Map
}

// NewInMemoryRepository returns a new instance of inMemoryRepository.
func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{
		count:    0,
		messages: sync.Map{},
	}
}

// Count returns the number of messages stored.
func (i *inMemoryRepository) Count(_ context.Context) (uint64, error) {
	return i.count, nil
}

// Store a new message.
func (i *inMemoryRepository) Store(_ context.Context, message *model.Message) error {
	atomic.AddUint64(&i.count, 1)
	i.messages.Store(message.Identifier(), *message)

	return nil
}

// UpdateText of an existing message and returns the updated model.
func (i *inMemoryRepository) UpdateText(_ context.Context, id, text string) (*model.Message, error) {
	messageI, found := i.messages.Load(id)
	if !found {
		return nil, errors.New("attempt to update a non existing message")
	}

	message := messageI.(model.Message)
	message.Text = text
	i.messages.Store(id, message)

	return &message, nil
}

// FetchByID attempts to fetch a message by id, if it is not found returns nil.
func (i *inMemoryRepository) FetchByID(_ context.Context, id string) (*model.Message, error) {
	messageI, found := i.messages.Load(id)
	if !found {
		return nil, nil
	}

	message := messageI.(model.Message)

	return &message, nil
}

// FetchMessages returns limit amount of messages sorted by Descending CreationDate.
func (i *inMemoryRepository) FetchMessages(ctx context.Context, offset, limit uint64) ([]model.Message, error) {
	messages := []model.Message{}

	i.messages.Range(func(key, value interface{}) bool {
		messages = append(messages, value.(model.Message))

		return true
	})

	totalItems, err := i.Count(ctx)
	if err != nil {
		return nil, err
	}

	if totalItems == 0 {
		return nil, nil
	}

	sort.Sort(helper.ByCreationDateDesc(messages))

	// we have less items than the limit
	if totalItems <= limit {
		return messages, nil
	}

	// we requested messages with an offset larger than the amount of messages
	if offset >= totalItems {
		return nil, errors.New("invalid offset")
	}

	upperBound := offset + limit
	if upperBound >= totalItems {
		upperBound = totalItems
	}

	return messages[offset:upperBound], nil
}
