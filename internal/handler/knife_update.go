package handler

import (
	"context"
	"errors"
	apperrors "knives/internal/errors"
	"knives/internal/handler/dto"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *KnifeHandler) Update(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	id := c.Params("id")

	var req dto.KnifeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": apperrors.MsgInvalidRequestBody})
	}
	knife := req.ToDomain()
	knife.ID = id

	updatedKnife, err := h.service.Update(ctx, knife)

	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "error": apperrors.MsgKnifeNotFound})
		}
		if errors.Is(err, apperrors.ErrValidation) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": apperrors.MsgInternalServerError})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "data": dto.KnifeFromDomain(updatedKnife)})
}
