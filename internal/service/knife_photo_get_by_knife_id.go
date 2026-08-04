package service

import (
	"context"
	"knives/internal/models"
)

func (s *KnifeService) GetPhotos(ctx context.Context, knifeID string) ([]models.KnifePhoto, error) {
	return s.photoRepo.GetByKnifeID(ctx, knifeID)
}