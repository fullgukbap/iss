package repository

import (
	"context"
	"letsgo-mini-is/internal/adapter/storage/mysql"
	"letsgo-mini-is/internal/core/domain"

	"github.com/google/uuid"
)

type PictureRepository struct {
	client *mysql.DB
}

func NewPictureRepository(client *mysql.DB) *PictureRepository {
	return &PictureRepository{
		client: client,
	}
}

func (r *PictureRepository) Create(c context.Context, p *domain.Picture) (*domain.Picture, error) {
	res, err := r.client.Picture.Create().
		SetID(p.ID).
		SetContent(p.Content).
		SetExtension(p.Extension).
		Save(c)
	return domainOf(res), err
}

func (r *PictureRepository) Find(c context.Context, id uuid.UUID) (*domain.Picture, error) {
	res, err := r.client.Picture.Get(c, id)
	return domainOf(res), err
}

func (r *PictureRepository) Update(c context.Context, p *domain.Picture) error {
	return r.client.Picture.Update().
		SetContent(p.Content).
		Exec(c)

}

func (r *PictureRepository) Delete(c context.Context, id uuid.UUID) error {
	return r.client.Picture.DeleteOneID(id).Exec(c)
}
