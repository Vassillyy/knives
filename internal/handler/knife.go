package handler

import (
	"knives/internal/models"
	"knives/internal/service"

	"github.com/gofiber/fiber/v2"
)

type KnifeHandler struct {
	service service.KnifeServiceInterface
}

func NewKnifeHandler(service service.KnifeServiceInterface) *KnifeHandler {
	return &KnifeHandler{service: service}
}

func (h *KnifeHandler) GetAll(c *fiber.Ctx) error {
	knives, err := h.service.GetAll(c.Context())
	if err != nil {
		return errorResponse(c, 500, err.Error())
	}
	return successResponse(c, 200, knives)
}

func (h *KnifeHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	knife, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return errorResponse(c, 500, err.Error())
	}
	if knife == nil {
		return errorResponse(c, 404, "knife not found")
	}
	return c.Status(200).JSON(fiber.Map{"success": true, "id": id, "data": knife})
}

func (h *KnifeHandler) Create(c *fiber.Ctx) error {
	knife := new(models.Knife)
	if err := c.BodyParser(knife); err != nil {
		return errorResponse(c, 400, "invalid request body")
	}

	if err := h.service.Create(c.Context(), knife); err != nil {
		return errorResponse(c, 500, err.Error())
	}
	return c.Status(201).JSON(fiber.Map{"success": true, "id": knife.ID, "data": knife})
}

func (h *KnifeHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	knife := new(models.Knife)
	if err := c.BodyParser(knife); err != nil {
		return errorResponse(c, 400, "invalid request body")
	}
	knife.ID = id

	if err := h.service.Update(c.Context(), knife); err != nil {
		return errorResponse(c, 500, err.Error())
	}
	return c.Status(200).JSON(fiber.Map{"success": true, "id": id, "data": knife})
}

func (h *KnifeHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.service.Delete(c.Context(), id); err != nil {
		return errorResponse(c, 500, err.Error())
	}
	return c.Status(200).JSON(fiber.Map{"success": true})
}
