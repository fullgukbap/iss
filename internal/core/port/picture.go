package port

import (
	"context"
	"letsgo-mini-is/internal/core/domain"

	"github.com/google/uuid"
)

type PictureService interface {
	Create(context.Context, *domain.Picture) (*domain.Picture, error)
	Find(context.Context, uuid.UUID) (*domain.Picture, error)
	Update(context.Context, *domain.Picture) (*domain.Picture, error)
	Delete(context.Context, uuid.UUID) error
}

type PictureRepository interface {
	Create(context.Context, *domain.Picture) (*domain.Picture, error)
	Find(context.Context, uuid.UUID) (*domain.Picture, error)
	Update(context.Context, *domain.Picture) (*domain.Picture, error)
	Delete(context.Context, uuid.UUID) error
}
