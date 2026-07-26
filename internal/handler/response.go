package handler

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func successResponse(c *fiber.Ctx, status int, data any) error {
	return c.Status(status).JSON(Response{Success: true, Data: data})
}

func errorResponse(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(Response{Success: false, Error: msg})
}
