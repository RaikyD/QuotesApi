package entity

import "github.com/google/uuid"

type Quote struct {
	ID     uuid.UUID `json:"id"`
	Author string    `json:"author"`
	Text   string    `json:"quote"`
}
