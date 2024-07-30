package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nglab-dev/nglab/handlers"
)

func Setup(app *fiber.App) {
	handlers.NewAuthHandler().Register(app)
}
