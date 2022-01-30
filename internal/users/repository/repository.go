package repository

import (
	"context"
	"messages_api/internal/users/model"
)

// Repository defines the available function for the user repository.
//go:generate mockery --name Repository
type Repository interface {
	GetAdmin(ctx context.Context, username string) (*model.User, error)
}

type inMemoryRepository struct {
	admins map[string]model.User
}

// NewInMemoryUserRepository initializes the inMemoryRepository with an example admin user.
func NewInMemoryUserRepository() *inMemoryRepository {
	return &inMemoryRepository{
		admins: map[string]model.User{
			"admin": {
				Username: "admin",
				Password: "back-challenge",
				Type:     model.AdminType,
			},
		},
	}
}

// GetAdmin returns the Admin model.User by username. If the user is not found or the user is not model.AdminType
// it returns nil.
func (r *inMemoryRepository) GetAdmin(_ context.Context, username string) (*model.User, error) {
	user, found := r.admins[username]
	if !found {
		return nil, nil
	}

	if user.Type != model.AdminType {
		return nil, nil
	}

	return &user, nil
}
