package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nglab-dev/nglab/auth"
)

type indexHandler struct {
	BaseHandler
}

func NewHomeHandler() indexHandler {
	return indexHandler{}
}

func (h indexHandler) Register(router fiber.Router) {
	router.Get("/", h.index)
	router.Get("/home", h.home)
}

func (h indexHandler) index(c fiber.Ctx) error {
	user := auth.GetUser(c)
	return c.Render("index", fiber.Map{
		"user": user,
	})
}

func (h indexHandler) home(c fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"title": "Home",
	})
}
