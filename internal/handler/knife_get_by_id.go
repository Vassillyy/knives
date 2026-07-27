package handler

import "github.com/gofiber/fiber/v2"

func (h *KnifeHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	knife, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	if knife == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "error": "knife not found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "id": id, "data": knife})
}
