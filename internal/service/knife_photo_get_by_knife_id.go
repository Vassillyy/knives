package service

import (
	"context"
	"knives/internal/domain"
)

func (s *KnifeService) GetPhotos(ctx context.Context, knifeID string) ([]domain.KnifePhoto, error) {
	return s.photoRepo.GetByKnifeID(ctx, knifeID)
}