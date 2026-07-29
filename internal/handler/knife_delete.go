package handler

import (
	"errors"
	apperrors "knives/internal/errors"

	"github.com/gofiber/fiber/v2"
)

func (h *KnifeHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.service.Delete(c.Context(), id); err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "error": "knife not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": "internal server error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "id": id})
}
