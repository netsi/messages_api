package repository_test

import (
	"context"
	"messages_api/internal/users/model"
	"messages_api/internal/users/repository"
	"reflect"
	"testing"
)

func Test_inMemoryRepository_GetAdmin(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name     string
		username string
		want     *model.User
		wantErr  bool
	}{
		{name: "not found", username: "unknown", want: nil, wantErr: false},
		{
			name:     "found",
			username: "admin",
			want: &model.User{
				Username: "admin",
				Password: "back-challenge",
				Type:     model.AdminType,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository.NewInMemoryUserRepository()
			got, err := r.GetAdmin(ctx, tt.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}
