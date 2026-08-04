package service

import (
	"context"
	"knives/internal/models"
)

func (s *KnifeService) GetByID(ctx context.Context, id string) (*models.Knife, error) {
	return s.repo.GetByID(ctx, id)
}