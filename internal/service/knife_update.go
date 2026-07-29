package service

import (
	"context"
	"fmt"
	apperrors "knives/internal/errors"
	"knives/internal/models"
)

func (s *KnifeService) Update(ctx context.Context, knife *models.Knife) (*models.Knife, error) {
	if knife.Name != nil && *knife.Name == "" {
		return nil, fmt.Errorf("%w: name cannot be empty", apperrors.ErrValidation)
	}
	if knife.Price != nil && *knife.Price <= 0 {
		return nil, fmt.Errorf("%w: price must be greater than 0", apperrors.ErrValidation)
	}

	return s.repo.Update(ctx, knife)
}
