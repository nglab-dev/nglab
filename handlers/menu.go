package handlers

import "github.com/gofiber/fiber/v3"

type menuHandler struct {
	BaseHandler
}

func NewMenuHandler() menuHandler {
	return menuHandler{}
}

func (h menuHandler) Register(router fiber.Router) {
	router.Get("/menu.html", h.html)
	router.Post("/menu", h.create)
}

func (h menuHandler) html(c fiber.Ctx) error {
	return c.Render("system/menu", fiber.Map{})
}

func (h menuHandler) create(c fiber.Ctx) error {
	return c.SendString("create menu")
}
