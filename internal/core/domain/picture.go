package domain

import "github.com/google/uuid"

type Picture struct {
	ID        uuid.UUID `json:"id"`
	Content   []byte    `json:"content"`
	Extension string    `json:"extension"`
}

func NewPicture(id uuid.UUID, content []byte, extension string) *Picture {
	return &Picture{
		ID:        id,
		Content:   content,
		Extension: extension,
	}
}
