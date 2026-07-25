package handler

import (
	"knifes/internal/models"
	"knifes/internal/service"

	"github.com/gofiber/fiber/v2"
)

type KnifeHandler struct {
	service *service.KnifeService
}

func NewKnifeHandler(service *service.KnifeService) *KnifeHandler {
	return &KnifeHandler{service: service}
}

func (h *KnifeHandler) GetAll(c *fiber.Ctx) error {
	knives, err := h.service.GetAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(knives)
}

func (h *KnifeHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	knife, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if knife == nil {
		return c.Status(404).JSON(fiber.Map{"error": "knife not found"})
	}
	return c.JSON(knife)
}

func (h *KnifeHandler) Create(c *fiber.Ctx) error {
	knife := new(models.Knife)
	if err := c.BodyParser(knife); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	if err := h.service.Create(c.Context(), knife); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(knife)
}

func (h *KnifeHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	knife := new(models.Knife)
	if err := c.BodyParser(knife); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}
	knife.ID = id

	if err := h.service.Update(c.Context(), knife); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(knife)
}

func (h *KnifeHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.service.Delete(c.Context(), id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(204).Send(nil)
}
