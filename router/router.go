package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nglab-dev/nglab/auth"
	"github.com/nglab-dev/nglab/handlers"
	"github.com/nglab-dev/nglab/middleware"
)

func Setup(app *fiber.App) {

	auth.Setup()

	handlers.NewAuthHandler().Register(app)

	authRouter := app.Group("", middleware.AuthRequired)

	handlers.NewHomeHandler().Register(authRouter)
}
