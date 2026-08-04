package service

import (
	"context"
	"fmt"
	"io"
	"knives/internal/models"

	"github.com/google/uuid"
)

func (s *KnifeService) UploadPhoto(ctx context.Context, knifeID, filename string, file io.Reader, size int64) (*models.KnifePhoto, error) {
	key := fmt.Sprintf("%s/%s/%s", knifeID, uuid.NewString(), filename)

	if err := s.s3Client.Put(ctx, key, file, size); err != nil {
		return nil, err
	}

	photo := &models.KnifePhoto{
		KnifeID:  knifeID,
		S3Key:    key,
		Filename: filename,
	}

	if err := s.photoRepo.Create(ctx, photo); err != nil {
		return nil, err
	}

	return photo, nil
}
