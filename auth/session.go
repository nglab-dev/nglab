package auth

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/nglab-dev/nglab/models"
)

var store *session.Store

func Setup() {
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		// CookieSecure: true, for https
		Expiration: time.Hour * 1,
	})
}

func SetSession(c fiber.Ctx, user *models.User) {
	store.Config.CookieHTTPOnly = true
	session := GetSession(c)

	userJsonStr, _ := json.Marshal(user)

	session.Set("user", string(userJsonStr))
	if err := session.Save(); err != nil {
		log.Errorf("Error saving session: %v", err)
	}
}

func GetSession(c fiber.Ctx) *session.Session {
	session, _ := store.Get(c)
	return session
}

func ClearSession(c fiber.Ctx) {
	session := GetSession(c)
	session.Delete("user")
}

func GetUser(c fiber.Ctx) (user *models.User) {
	session := GetSession(c)
	userJsonStr := session.Get("user")
	if userJsonStr == nil {
		return nil
	}

	json.Unmarshal([]byte(userJsonStr.(string)), &user)

	return
}
