package service

import (
	"context"
	"io"
	"knives/internal/models"
	"knives/internal/repository"
	"knives/internal/s3"
)

type KnifeServiceInterface interface {
	GetAll(ctx context.Context) ([]models.Knife, error)
	GetByID(ctx context.Context, id string) (*models.Knife, error)
	Create(ctx context.Context, knife *models.Knife) error
	Update(ctx context.Context, knife *models.Knife) (*models.Knife, error)
	Delete(ctx context.Context, id string) error
	UploadPhoto(ctx context.Context, knifeID, filename string, file io.Reader, size int64) (*models.KnifePhoto, error)
	GetPhotos(ctx context.Context, knifeID string) ([]models.KnifePhoto, error)
	GetPhoto(ctx context.Context, id string) (io.ReadCloser, *models.KnifePhoto, error)
	DeletePhoto(ctx context.Context, id string) error
}

type KnifeService struct {
	repo      repository.KnifeRepositoryInterface
	photoRepo repository.KnifePhotoRepositoryInterface
	s3Client  s3.StorageInterface
}

func NewKnifeService(repo repository.KnifeRepositoryInterface, photoRepo repository.KnifePhotoRepositoryInterface, s3Client s3.StorageInterface) *KnifeService {
	return &KnifeService{repo: repo, photoRepo: photoRepo, s3Client: s3Client}
}
