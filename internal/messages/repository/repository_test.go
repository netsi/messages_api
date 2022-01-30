package repository_test

import (
	"context"
	"messages_api/internal/messages/model"
	"messages_api/internal/messages/repository"
	"messages_api/internal/test"
	"reflect"
	"testing"
	"time"
)

func Test_inMemoryRepository_Count(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name       string
		messageIDs []string
		want       uint64
		wantErr    bool
	}{
		{name: "empty", want: 0},
		{name: "1 message", messageIDs: []string{"id-1"}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := repository.NewInMemoryMessagesRepository()

			for _, id := range tt.messageIDs {
				_ = i.Store(ctx, test.NewMessageBuilder(id).Build())
			}

			got, err := i.Count(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Count() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemoryRepository_FetchByID(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name       string
		messageIDs []string
		fetchId    string
		want       *model.Message
		wantErr    bool
	}{
		{name: "not found", fetchId: "id-1"},
		{name: "found", fetchId: "id-1", messageIDs: []string{"id-2", "id-1"}, want: test.NewMessageBuilder("id-1").Build()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := repository.NewInMemoryMessagesRepository()

			for _, id := range tt.messageIDs {
				_ = i.Store(ctx, test.NewMessageBuilder(id).Build())
			}

			got, err := i.FetchByID(ctx, tt.fetchId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemoryRepository_FetchMessages(t *testing.T) {
	ctx := context.Background()

	type args struct {
		offset uint64
		limit  uint64
	}
	tests := []struct {
		name     string
		messages []*model.Message
		args     args
		want     []model.Message
		wantErr  bool
	}{
		{
			name: "empty",
			args: args{
				offset: 0,
				limit:  2,
			},
			want: nil,
		},
		{
			name:     "one message",
			messages: []*model.Message{test.NewMessageBuilder("id-1").Build()},
			args: args{
				offset: 0,
				limit:  2,
			},
			want: []model.Message{*(test.NewMessageBuilder("id-1").Build())},
		},
		{
			name: "two messages sorted",
			messages: []*model.Message{
				test.NewMessageBuilder("id-1").
					WithCreationDate(test.MustParseTimeFromFormat(time.RFC3339, "2022-01-01T00:00:00Z")).
					Build(),
				test.NewMessageBuilder("id-2").
					WithCreationDate(test.MustParseTimeFromFormat(time.RFC3339, "2022-01-02T00:00:00Z")).
					Build(),
			},
			args: args{
				offset: 0,
				limit:  2,
			},
			want: []model.Message{
				*(test.NewMessageBuilder("id-2").
					WithCreationDate(test.MustParseTimeFromFormat(time.RFC3339, "2022-01-02T00:00:00Z")).
					Build()),
				*(test.NewMessageBuilder("id-1").
					WithCreationDate(test.MustParseTimeFromFormat(time.RFC3339, "2022-01-01T00:00:00Z")).
					Build()),
			},
		},
		{
			name: "two messages limit 1",
			messages: []*model.Message{
				test.NewMessageBuilder("id-1").
					WithCreationDate(test.MustParseTimeFromFormat(time.RFC3339, "2022-01-01T00:00:00Z")).
					Build(),
				test.NewMessageBuilder("id-2").
					WithCreationDate(test.MustParseTimeFromFormat(time.RFC3339, "2022-01-02T00:00:00Z")).
					Build(),
			},
			args: args{
				offset: 0,
				limit:  1,
			},
			want: []model.Message{
				*(test.NewMessageBuilder("id-2").
					WithCreationDate(test.MustParseTimeFromFormat(time.RFC3339, "2022-01-02T00:00:00Z")).
					Build()),
			},
		},
		{
			name: "two messages limit 1 offset 1",
			messages: []*model.Message{
				test.NewMessageBuilder("id-1").
					WithCreationDate(test.MustParseTimeFromFormat(time.RFC3339, "2022-01-01T00:00:00Z")).
					Build(),
				test.NewMessageBuilder("id-2").
					WithCreationDate(test.MustParseTimeFromFormat(time.RFC3339, "2022-01-02T00:00:00Z")).
					Build(),
			},
			args: args{
				offset: 1,
				limit:  1,
			},
			want: []model.Message{
				*(test.NewMessageBuilder("id-1").
					WithCreationDate(test.MustParseTimeFromFormat(time.RFC3339, "2022-01-01T00:00:00Z")).
					Build()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := repository.NewInMemoryMessagesRepository()
			for _, msg := range tt.messages {
				_ = i.Store(ctx, msg)
			}

			got, err := i.FetchMessages(ctx, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchMessages() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemoryRepository_UpdateText(t *testing.T) {
	ctx := context.Background()

	type args struct {
		id   string
		text string
	}
	tests := []struct {
		name       string
		messageIDs []string
		args       args
		want       *model.Message
		wantErr    bool
	}{
		{
			name:       "update",
			messageIDs: []string{"id-1"},
			args: args{
				id:   "id-1",
				text: "new text",
			},
			want: test.NewMessageBuilder("id-1").WithText("new text").Build(),
		},
		{
			name:       "update non existent item",
			messageIDs: []string{"id-1"},
			args: args{
				id:   "id-2",
				text: "new text",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := repository.NewInMemoryMessagesRepository()
			for _, id := range tt.messageIDs {
				_ = i.Store(ctx, test.NewMessageBuilder(id).Build())
			}

			got, err := i.UpdateText(ctx, tt.args.id, tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateText() got = %v, want %v", got, tt.want)
			}
		})
	}
}
