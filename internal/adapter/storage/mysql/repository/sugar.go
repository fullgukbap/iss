package repository

import (
	"letsgo-mini-is/internal/adapter/storage/mysql/ent"
	"letsgo-mini-is/internal/core/domain"
)

func domainOf(entity *ent.Picture) *domain.Picture {
	if entity == nil {
		return nil
	}
	return &domain.Picture{
		ID:        entity.ID,
		Content:   entity.Content,
		Extension: entity.Extension,
	}
}
