package repository

import (
	"context"
	"knives/internal/domain"
)

type KnifePhotoRepositoryInterface interface {
	Create(ctx context.Context, photo *domain.KnifePhoto) error
	GetByID(ctx context.Context, id string) (*domain.KnifePhoto, error)
	GetByKnifeID(ctx context.Context, knifeID string) ([]domain.KnifePhoto, error)
	Delete(ctx context.Context, id string) error
}
