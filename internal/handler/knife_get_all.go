package handler

import "github.com/gofiber/fiber/v2"

func (h *KnifeHandler) GetAll(c *fiber.Ctx) error {
	knives, err := h.service.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "data": knives})
}
