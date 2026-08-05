package handler

import (
	"context"
	"knives/internal/handler/dto"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *KnifeHandler) GetAll(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	knives, err := h.service.GetAll(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "data": dto.KnivesFromDomain(knives)})
}
