package main

import (
	"context"
	"log"

	"knives/internal/config"
	"knives/internal/handler"
	"knives/internal/repository/postgres"
	"knives/internal/service"
	"knives/internal/storage/minio"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	miniogo "github.com/minio/minio-go/v7"
)

func main() {
	cfg := config.Load()

	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	defer pool.Close()

	storageClient, err := minio.New(cfg.MinIOEndpoint, cfg.MinIOAccessKey, cfg.MinIOSecretKey, cfg.MinIOBucket)
	if err != nil {
		log.Fatal("failed to connect to minio: ", err)
	}

	bucketExists, err := storageClient.BucketExists(context.Background(), cfg.MinIOBucket)
	if err != nil {
		log.Fatal("failed to check minio bucket: ", err)
	}
	if !bucketExists {
		if err := storageClient.MakeBucket(context.Background(), cfg.MinIOBucket, miniogo.MakeBucketOptions{}); err != nil {
			log.Fatal("failed to create minio bucket: ", err)
		}
	}

	repo := postgres.NewKnifeRepository(pool)
	photoRepo := postgres.NewKnifePhotoRepository(pool)
	svc := service.NewKnifeService(repo, photoRepo, storageClient)
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
