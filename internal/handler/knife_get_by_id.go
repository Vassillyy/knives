package handler

import (
	"context"
	apperrors "knives/internal/errors"
	"knives/internal/handler/dto"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *KnifeHandler) GetByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	id := c.Params("id")

	knife, err := h.service.GetByID(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": apperrors.MsgInternalServerError})
	}
	if knife == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "error": apperrors.MsgKnifeNotFound})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "data": dto.KnifeFromDomain(knife)})
}
