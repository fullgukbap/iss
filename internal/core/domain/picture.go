package domain

import "github.com/google/uuid"

type Picture struct {
	ID      uuid.UUID `json:"id"`
	Content []byte    `json:"content"`
}

func NewPicture(id uuid.UUID, content []byte) *Picture {
	return &Picture{
		ID:      id,
		Content: content,
	}
}
