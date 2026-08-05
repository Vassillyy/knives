package main

import (
	"context"
	"knives/internal/s3"
	"log"

	"knives/internal/config"
	"knives/internal/handler"
	"knives/internal/repository/postgres"
	"knives/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
)

func main() {
	cfg := config.Load()

	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	defer pool.Close()

	s3Client, err := s3.New(cfg.MinIOEndpoint, cfg.MinIOAccessKey, cfg.MinIOSecretKey, cfg.MinIOBucket)
	if err != nil {
		log.Fatal("failed to connect to minio: ", err)
	}

	bucketExists, err := s3Client.BucketExists(context.Background(), cfg.MinIOBucket)
	if err != nil {
		log.Fatal("failed to check minio bucket: ", err)
	}
	if !bucketExists {
		if err := s3Client.MakeBucket(context.Background(), cfg.MinIOBucket, minio.MakeBucketOptions{}); err != nil {
			log.Fatal("failed to create minio bucket: ", err)
		}
	}

	repo := postgres.NewKnifeRepository(pool)
	photoRepo := postgres.NewKnifePhotoRepository(pool)
	svc := service.NewKnifeService(repo, photoRepo, s3Client)
	h := handler.NewKnifeHandler(svc)

	app := fiber.New(fiber.Config{
		BodyLimit: 25 * 1024 * 1024,
	})
	app.Use(cors.New())

	api := app.Group("/api/v1")
	knives := api.Group("/knives")
	knives.Get("/", h.GetAll)
	knives.Get("/:id", h.GetByID)
	knives.Post("/", h.Create)
	knives.Patch("/:id", h.Update)
	knives.Delete("/:id", h.Delete)
	knives.Post("/:id/photos", h.UploadPhoto)
	knives.Get("/:id/photos", h.GetPhotos)
	knives.Get("/:id/photos/:photoId/file", h.GetPhoto)
	knives.Delete("/:id/photos/:photoId", h.DeletePhoto)

	log.Fatal(app.Listen(":" + cfg.Port))
}
