package repository

import (
	"context"
	"knives/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type KnifePhotoRepositoryInterface interface {
	Create(ctx context.Context, photo *models.KnifePhoto) error
	GetByID(ctx context.Context, id string) (*models.KnifePhoto, error)
	GetByKnifeID(ctx context.Context, knifeID string) ([]models.KnifePhoto, error)
	Delete(ctx context.Context, id string) error
}

type KnifePhotoRepository struct {
	db *pgxpool.Pool
}

func NewKnifePhotoRepository(db *pgxpool.Pool) *KnifePhotoRepository {
	return &KnifePhotoRepository{db: db}
}
