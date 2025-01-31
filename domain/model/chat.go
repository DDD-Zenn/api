package model

import (
	"context"
	"time"
)

type Chat struct {
	ID        string
	Name      string
	Response  string
	CreatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewChat(ctx context.Context, id, userId, response string) (chat Chat, err error) {

	chat.ID = id
	chat.Response = response
	chat.CreatedBy = userId
	chat.CreatedAt = time.Now()
	chat.UpdatedAt = time.Now()

	return
}
