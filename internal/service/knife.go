package service

import (
	"context"
	"io"
	"knives/internal/domain"
	"knives/internal/repository"
	"knives/internal/s3"
)

type KnifeServiceInterface interface {
	GetAll(ctx context.Context) ([]domain.Knife, error)
	GetByID(ctx context.Context, id string) (*domain.Knife, error)
	Create(ctx context.Context, knife *domain.Knife) error
	Update(ctx context.Context, knife *domain.Knife) (*domain.Knife, error)
	Delete(ctx context.Context, id string) error
	UploadPhoto(ctx context.Context, knifeID, filename string, file io.Reader, size int64) (*domain.KnifePhoto, error)
	GetPhotos(ctx context.Context, knifeID string) ([]domain.KnifePhoto, error)
	GetPhoto(ctx context.Context, id string) (io.ReadCloser, *domain.KnifePhoto, error)
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
