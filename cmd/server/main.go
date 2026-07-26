package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"knives/internal/config"
	"knives/internal/handler"
	"knives/internal/repository"
	"knives/internal/service"
)

func main() {
	cfg := config.Load()

	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	defer pool.Close()

	repo := repository.NewKnifeRepository(pool)
	svc := service.NewKnifeService(repo)
	h := handler.NewKnifeHandler(svc)

	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api/v1")
	knives := api.Group("/knives")
	knives.Get("/", h.GetAll)
	knives.Get("/:id", h.GetByID)
	knives.Post("/", h.Create)
	knives.Patch("/:id", h.Update)
	knives.Delete("/:id", h.Delete)

	log.Fatal(app.Listen(":" + cfg.Port))
}
