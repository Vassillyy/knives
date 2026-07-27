package handler

import (
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

	if err := h.service.Update(c.Context(), knife); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "id": id, "data": knife})
}
