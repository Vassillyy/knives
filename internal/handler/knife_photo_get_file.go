package handler

import (
	"errors"
	apperrors "knives/internal/errors"
	"mime"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func (h *KnifeHandler) GetPhoto(c *fiber.Ctx) error {
	photoID := c.Params("photoId")

	file, photo, err := h.service.GetPhoto(c.Context(), photoID)
	if err != nil {
		if errors.Is(err, apperrors.ErrNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "error": apperrors.MsgPhotoNotFound})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": apperrors.MsgInternalServerError})
	}
	contentType := mime.TypeByExtension(filepath.Ext(photo.Filename))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	c.Set("Content-Type", contentType)

	return c.SendStream(file)
}