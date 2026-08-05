package handler

import (
	"context"
	apperrors "knives/internal/errors"
	"knives/internal/handler/dto"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *KnifeHandler) GetPhotos(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	knifeID := c.Params("id")

	photos, err := h.service.GetPhotos(ctx, knifeID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": apperrors.MsgInternalServerError})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "data": dto.KnifePhotosFromDomain(photos)})
}