package handlers

import "github.com/gofiber/fiber/v3"

func HandleIndexView(c fiber.Ctx) error {
	userId := c.Locals("userId")
	return c.Render("index", fiber.Map{
		"Title": userId,
	})
}
