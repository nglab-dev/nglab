package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/nglab-dev/nglab/auth"
	"github.com/nglab-dev/nglab/handlers"
	"github.com/nglab-dev/nglab/middleware"
)

func Setup(app *fiber.App) {

	auth.Setup()

	app.Use("/static", static.New("./web/static"))

	handlers.NewAuthHandler().Register(app)

	rootRouter := app.Group("", middleware.AuthRequired)
	handlers.NewHomeHandler().Register(rootRouter)

	systemRouter := app.Group("system", middleware.AuthRequired)
	handlers.NewMenuHandler().Register(systemRouter)
}
