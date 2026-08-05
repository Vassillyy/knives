package repository

import (
	"context"
	"knives/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type KnifePhotoRepositoryInterface interface {
	Create(ctx context.Context, photo *domain.KnifePhoto) error
	GetByID(ctx context.Context, id string) (*domain.KnifePhoto, error)
	GetByKnifeID(ctx context.Context, knifeID string) ([]domain.KnifePhoto, error)
	Delete(ctx context.Context, id string) error
}

type KnifePhotoRepository struct {
	db *pgxpool.Pool
}

func NewKnifePhotoRepository(db *pgxpool.Pool) *KnifePhotoRepository {
	return &KnifePhotoRepository{db: db}
}
