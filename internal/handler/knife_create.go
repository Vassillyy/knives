package handler

import (
	"context"
	"errors"
	apperrors "knives/internal/errors"
	"knives/internal/handler/dto"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *KnifeHandler) Create(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	var req dto.KnifeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": apperrors.MsgInvalidRequestBody})
	}

	knife := req.ToDomain()
	if err := h.service.Create(ctx, knife); err != nil {
		if errors.Is(err, apperrors.ErrValidation) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": apperrors.MsgInternalServerError})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "data": dto.KnifeFromDomain(knife)})
}
