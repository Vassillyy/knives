package service

import (
	"context"
	"io"
	"knives/internal/domain"
)

func (s *KnifeService) GetPhoto(ctx context.Context, id string) (io.ReadCloser, *domain.KnifePhoto, error) {
	photo, err := s.photoRepo.GetByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	obj, err := s.storage.Get(ctx, photo.S3Key)
	if err != nil {
		return nil, nil, err
	}

	return obj, photo, nil
}
