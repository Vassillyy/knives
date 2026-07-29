package service

import (
	"context"
	"fmt"
	apperrors "knives/internal/errors"
	"knives/internal/models"
)

func (s *KnifeService) Create(ctx context.Context, knife *models.Knife) error {
	if knife.Name == nil || *knife.Name == "" {
		return fmt.Errorf("%w: name is required", apperrors.ErrValidation)
	}
	if knife.Price == nil || *knife.Price <= 0 {
		return fmt.Errorf("%w: price must be greater than 0", apperrors.ErrValidation)
	}
	return s.repo.Create(ctx, knife)
}
