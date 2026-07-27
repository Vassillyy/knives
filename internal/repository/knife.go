package repository

import (
	"context"
	"knives/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type KnifeRepositoryInterface interface {
	GetAll(ctx context.Context) ([]models.Knife, error)
	GetByID(ctx context.Context, id string) (*models.Knife, error)
	Create(ctx context.Context, knife *models.Knife) error
	Update(ctx context.Context, knife *models.Knife) error
	Delete(ctx context.Context, id string) error
}

type KnifeRepository struct {
	db *pgxpool.Pool
}

func NewKnifeRepository(db *pgxpool.Pool) *KnifeRepository {
	return &KnifeRepository{db: db}
}
