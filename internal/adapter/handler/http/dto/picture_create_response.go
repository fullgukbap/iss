package dto

import (
	"letsgo-mini-is/internal/core/domain"

	"github.com/google/uuid"
)

type PictureCreateResponse struct {
	ID uuid.UUID `json:"id"`
}

func PictureCreateResponseOf(p *domain.Picture) PictureCreateResponse {
	return PictureCreateResponse{
		ID: p.ID,
	}
}
