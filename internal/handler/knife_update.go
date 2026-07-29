package handler

import (
	"errors"
	apperrors "knives/internal/errors"
	"knives/internal/models"

	"github.com/gofiber/fiber/v2"
)

func (h *KnifeHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	knife := new(models.Knife)
	if err := c.BodyParser(knife); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid request body"})
	}
	knife.ID = id

	updatedKnife, err := h.service.Update(c.Context(), knife)

	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "error": "knife not found"})
		}
		if errors.Is(err, apperrors.ErrValidation) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": "internal server error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "data": updatedKnife})
}
