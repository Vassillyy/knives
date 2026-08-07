package service

import (
	"context"
	"knives/internal/domain"
)

func (s *KnifeService) DeletePhoto(ctx context.Context, id string) error {
	photo, err := s.photoRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	return s.deletePhoto(ctx, photo)
}

func (s *KnifeService) deletePhoto(ctx context.Context, photo *domain.KnifePhoto) error {
	if err := s.storage.Remove(ctx, photo.S3Key); err != nil {
		return err
	}
	return s.photoRepo.Delete(ctx, photo.ID)
}