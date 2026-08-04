package service

import (
	"context"
	"io"
	apperrors "knives/internal/errors"
	"knives/internal/models"
)

func (s *KnifeService) GetPhoto(ctx context.Context, id string) (io.ReadCloser, *models.KnifePhoto, error) {
	photo, err := s.photoRepo.GetByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	if photo == nil {
		return nil, nil, apperrors.ErrNotFound
	}

	obj, err := s.s3Client.Get(ctx, photo.S3Key)
	if err != nil {
		return nil, nil, err
	}

	return obj, photo, nil
}