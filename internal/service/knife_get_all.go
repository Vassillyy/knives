package service

import (
	"context"
	"knives/internal/models"
)

func (s *KnifeService) GetAll(ctx context.Context) ([]models.Knife, error) {
	return s.repo.GetAll(ctx)
}
