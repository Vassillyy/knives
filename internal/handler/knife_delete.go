package handler

import (
	"context"
	"errors"
	apperrors "knives/internal/errors"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *KnifeHandler) Delete(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	id := c.Params("id")

	if err := h.service.Delete(ctx, id); err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "error": apperrors.MsgKnifeNotFound})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": apperrors.MsgInternalServerError})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "id": id})
}
