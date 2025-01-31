package chat

import "time"

type ChatDTO struct {
	ID        string    `json:"id"`
	Response  string    `json:"response"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
