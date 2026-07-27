package service

import (
	"context"
	"errors"
	"knives/internal/models"
)

func (s *KnifeService) Update(ctx context.Context, knife *models.Knife) error {
	if knife.Name != nil && *knife.Name == "" {
		return errors.New("name cannot be empty")
	}
	if knife.Price != nil && *knife.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	return s.repo.Update(ctx, knife)
}
