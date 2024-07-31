package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nglab-dev/nglab/auth"
)

func HandleIndexView(c fiber.Ctx) error {
	userId := c.Locals("userId")
	return c.Render("index", fiber.Map{
		"Title": userId,
	})
}

type homeHandler struct {
	BaseHandler
}

func NewHomeHandler() homeHandler {
	return homeHandler{}
}

func (h homeHandler) Register(router fiber.Router) {
	router.Get("/", h.home)
}

func (h homeHandler) home(c fiber.Ctx) error {
	user := auth.GetUser(c)
	return h.HTML(c, "home", fiber.Map{
		"user": user,
	})
}
