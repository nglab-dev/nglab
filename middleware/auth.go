package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/nglab-dev/nglab/auth"
)

func AuthRequired(c fiber.Ctx) error {

	user := auth.GetUser(c)
	if user == nil {
		c.Set("Content-Type", "text/html; charset=utf-8")
		return c.SendString(`<script type="text/javascript">top.location.href="/login"</script>`)
	}

	c.Set("user_id", strconv.FormatInt(user.ID, 10))
	c.Set("user_name", user.Username)

	return c.Next()
}
