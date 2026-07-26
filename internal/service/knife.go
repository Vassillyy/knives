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
	Update(ctx context.Context, knife *models.Knife) error
	Delete(ctx context.Context, id string) error
}

type KnifeService struct {
	repo repository.KnifeRepositoryInterface
}

func NewKnifeService(repo repository.KnifeRepositoryInterface) *KnifeService {
	return &KnifeService{repo: repo}
}

func (s *KnifeService) GetAll(ctx context.Context) ([]models.Knife, error) {
	return s.repo.GetAll(ctx)
}

func (s *KnifeService) GetByID(ctx context.Context, id string) (*models.Knife, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *KnifeService) Create(ctx context.Context, knife *models.Knife) error {
	if err := validateKnifeCreate(knife); err != nil {
		return err
	}
	return s.repo.Create(ctx, knife)
}

func (s *KnifeService) Update(ctx context.Context, knife *models.Knife) error {
	if err := validateKnifeUpdate(knife); err != nil {
		return err
	}
	return s.repo.Update(ctx, knife)
}

func (s *KnifeService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
