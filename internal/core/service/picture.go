package service

import (
	"context"
	"letsgo-mini-is/internal/core/domain"
	"letsgo-mini-is/internal/core/port"

	"github.com/google/uuid"
)

type PictureService struct {
	pictureRepository port.PictureRepository
}

func NewPictureService(pictureRepository port.PictureRepository) *PictureService {
	return &PictureService{
		pictureRepository: pictureRepository,
	}
}

func (s *PictureService) Create(c context.Context, p *domain.Picture) (*domain.Picture, error) {
	return s.Create(c, p)
}

func (s *PictureService) Find(c context.Context, id uuid.UUID) (*domain.Picture, error) {
	return s.Find(c, id)
}

func (s *PictureService) Update(c context.Context, p *domain.Picture) (*domain.Picture, error) {
	return s.Update(c, p)
}

func (s *PictureService) Delete(c context.Context, id uuid.UUID) error {
	return s.Delete(c, id)
}
