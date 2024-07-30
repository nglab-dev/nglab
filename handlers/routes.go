package handlers

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

var store *session.Store

const (
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

func Setup(app *fiber.App) {
	/* Sessions Config */
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		// CookieSecure: true, for https
		Expiration: time.Hour * 1,
	})

	/* Views */
	app.Get("/register", HandleRegisterView)
	app.Get("/login", HandleLoginView)

	app.Post("/register", HandleRegister)
	app.Post("/login", HandleLogin)

	auth := app.Group("", withAuth)
	auth.Get("/", HandleIndexView)

}

func withAuth(c fiber.Ctx) error {
	session, err := store.Get(c)

	if err != nil {
		return c.Redirect().Status(fiber.StatusMovedPermanently).To("/login")
	}

	if session.Get(AUTH_KEY) == nil {
		return c.Redirect().Status(fiber.StatusMovedPermanently).To("/login")
	}

	userId := session.Get(USER_ID)
	if userId == nil {
		return c.Redirect().Status(fiber.StatusMovedPermanently).To("/login")
	}

	c.Locals("userId", userId)

	return c.Next()
}
