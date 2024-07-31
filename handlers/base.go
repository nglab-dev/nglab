package handlers

import (
	"github.com/gofiber/fiber/v3"
)

type BaseHandler struct {
}

func (h *BaseHandler) Ok(c fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"code": 0,
		"data": data,
	})
}

func (h *BaseHandler) Error(c fiber.Ctx, message string) error {
	return c.JSON(fiber.Map{
		"code":    1,
		"message": message,
	})
}

func (h *BaseHandler) HTML(c fiber.Ctx, name string, data fiber.Map) error {
	return c.Render(name, data, "layouts/default")
}
