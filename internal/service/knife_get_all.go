package service

import (
	"context"
	"knives/internal/domain"
)

func (s *KnifeService) GetAll(ctx context.Context) ([]domain.Knife, error) {
	return s.repo.GetAll(ctx)
}
