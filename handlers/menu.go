package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nglab-dev/nglab/models"
)

type menuHandler struct {
	BaseHandler
}

func NewMenuHandler() menuHandler {
	return menuHandler{}
}

func (h menuHandler) Register(router fiber.Router) {
	router.Get("/menu.html", h.html)
	router.Get("/menu/tree", h.tree)
	router.Post("/menu", h.create)
}

func (h menuHandler) html(c fiber.Ctx) error {
	return c.Render("system/menu", fiber.Map{})
}

func (h menuHandler) tree(c fiber.Ctx) error {
	menus, err := models.GetMenuTree()
	if err != nil {
		h.Error(c, err.Error())
	}
	return h.Ok(c, menus)
}

func (h menuHandler) create(c fiber.Ctx) error {
	var menu models.Menu
	if err := c.Bind().JSON(&menu); err != nil {
		h.Error(c, err.Error())
	}
	if err := models.CreateMenu(&menu); err != nil {
		h.Error(c, err.Error())
	}
	return h.Ok(c, nil)
}
