package handler

import (
	"context"
	apperrors "knives/internal/errors"
	"knives/internal/handler/dto"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *KnifeHandler) UploadPhoto(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	knifeID := c.Params("id")

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": apperrors.MsgFileRequired})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": apperrors.MsgInternalServerError})
	}
	defer file.Close()

	photo, err := h.service.UploadPhoto(ctx, knifeID, fileHeader.Filename, file, fileHeader.Size)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": apperrors.MsgInternalServerError})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "data": dto.KnifePhotoFromDomain(photo)})
}