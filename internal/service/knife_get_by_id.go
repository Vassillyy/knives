package service

import (
	"context"
	"knives/internal/domain"
)

func (s *KnifeService) GetByID(ctx context.Context, id string) (*domain.Knife, error) {
	return s.repo.GetByID(ctx, id)
}