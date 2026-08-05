package service

import (
	"context"
	"knives/internal/domain"
	apperrors "knives/internal/errors"
)

func (s *KnifeService) DeletePhoto(ctx context.Context, id string) error {
	photo, err := s.photoRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if photo == nil {
		return apperrors.ErrNotFound
	}
	return s.deletePhoto(ctx, photo)
}

func (s *KnifeService) deletePhoto(ctx context.Context, photo *domain.KnifePhoto) error {
	if err := s.s3Client.Remove(ctx, photo.S3Key); err != nil {
		return err
	}
	return s.photoRepo.Delete(ctx, photo.ID)
}