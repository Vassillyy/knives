package service

import (
	"context"
	"knives/internal/models"
	"knives/internal/repository"
)

type KnifeServiceInterface interface {
	GetAll(ctx context.Context) ([]models.Knife, error)
	GetByID(ctx context.Context, id string) (*models.Knife, error)
	Create(ctx context.Context, knife *models.Knife) error
	Update(ctx context.Context, knife *models.Knife) (*models.Knife, error)
	Delete(ctx context.Context, id string) error
}

type KnifeService struct {
	repo repository.KnifeRepositoryInterface
}

func NewKnifeService(repo repository.KnifeRepositoryInterface) *KnifeService {
	return &KnifeService{repo: repo}
}
