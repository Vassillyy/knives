package handler

import (
	"errors"
	apperrors "knives/internal/errors"
	"knives/internal/models"

	"github.com/gofiber/fiber/v2"
)

func (h *KnifeHandler) Create(c *fiber.Ctx) error {
	knife := new(models.Knife)
	if err := c.BodyParser(knife); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "invalid request body"})
	}

	if err := h.service.Create(c.Context(), knife); err != nil {
		if errors.Is(err, apperrors.ErrValidation) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": "internal server error"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "data": knife})
}
