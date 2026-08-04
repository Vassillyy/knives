package handler

import (
	"context"
	"errors"
	apperrors "knives/internal/errors"
	"knives/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *KnifeHandler) Update(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	id := c.Params("id")

	knife := new(models.Knife)
	if err := c.BodyParser(knife); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": apperrors.MsgInvalidRequestBody})
	}
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "data": updatedKnife})
}
