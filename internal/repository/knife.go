package repository

import (
	"context"
	"knives/internal/domain"
)

type KnifeRepositoryInterface interface {
	GetAll(ctx context.Context) ([]domain.Knife, error)
	GetByID(ctx context.Context, id string) (*domain.Knife, error)
	Create(ctx context.Context, knife *domain.Knife) error
	Update(ctx context.Context, knife *domain.Knife) (*domain.Knife, error)
	Delete(ctx context.Context, id string) error
}
