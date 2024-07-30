package handlers

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

type BaseHandler struct {
}

var store = session.New(session.Config{
	CookieHTTPOnly: true,
	// CookieSecure: true, for https
	Expiration: time.Hour * 1,
})

func SetSession(c fiber.Ctx, username string) {
	store.Config.CookieHTTPOnly = true
	session := GetSession(c)
	session.Set("username", username)
	session.Save()

	c.Locals("session", session)
}

func GetSession(c fiber.Ctx) *session.Session {
	session, _ := store.Get(c)
	return session
}

func (b *BaseHandler) Ok(c fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"code": 0,
		"data": data,
	})
}

func (b *BaseHandler) Error(c fiber.Ctx, code int, message string) error {
	return c.JSON(fiber.Map{
		"code":    code,
		"message": message,
	})
}

func (b *BaseHandler) HTML(c fiber.Ctx, name string, data fiber.Map) error {
	return c.Render(name, data, "layouts/default")
}
