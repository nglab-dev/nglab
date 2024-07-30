package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nglab-dev/nglab/models"
	"github.com/nglab-dev/nglab/utils/env"
)

type RegisterRequest struct {
	Username        string `json:"username" form:"username"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type authHandler struct {
	BaseHandler
}

func NewAuthHandler() authHandler {
	return authHandler{}
}

func (h authHandler) Register(app *fiber.App) {
	app.Get("/login", h.login)
	app.Post("/login", h.login)
}

func (h authHandler) login(c fiber.Ctx) error {
	appName := env.GetString("APP_NAME", "nglab")
	// get request method
	if c.Method() == "GET" {
		return c.Render("login", fiber.Map{
			"title": appName,
		})
	} else {
		username := c.FormValue("username")
		password := c.FormValue("password")

		user, err := models.GetUserByUsername(username)
		if err != nil {

		}
	}

}

// func HandleLoginView(c fiber.Ctx) error {
// 	return c.Render("login", fiber.Map{})
// }

// func HandleRegisterView(c fiber.Ctx) error {
// 	return c.Render("register", fiber.Map{})
// }

// func HandleRegister(c fiber.Ctx) error {
// 	req := new(RegisterRequest)

// 	if err := c.Bind().Form(req); err != nil {
// 		return c.Render("signup", fiber.Map{
// 			"error": "Invalid input",
// 		})
// 	}

// 	if req.Password != req.ConfirmPassword {
// 		return c.Render("signup", fiber.Map{
// 			"error": "Passwords do not match",
// 		})
// 	}

// 	user, err := models.GetUserByUsername(req.Username)
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return c.Render("signup", fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	if user.ID != 0 {
// 		log.Infof("user id is %s", user.ID)
// 		return c.Render("signup", fiber.Map{
// 			"error": "Username already exists",
// 		})
// 	}

// 	newUser := new(models.User)
// 	newUser.Username = req.Username
// 	newUser.Password = req.Password

// 	err = models.CreateUser(newUser)
// 	if err != nil {
// 		return c.Render("signup", fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	return c.Redirect().Status(fiber.StatusMovedPermanently).To("/login")
// }

// func HandleLogin(c fiber.Ctx) error {
// 	req := new(LoginRequest)
// 	if err := c.Bind().Form(req); err != nil {
// 		return c.Render("login", fiber.Map{
// 			"error": "Invalid input",
// 		})
// 	}

// 	user, err := models.GetUserByUsername(req.Username)
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return c.Render("login", fiber.Map{
// 				"error": "User not found",
// 			})
// 		}
// 	}

// 	if err := user.ComparePassword(req.Password); err != nil {
// 		return c.Render("login", fiber.Map{
// 			"error": "Incorrect password",
// 		})
// 	}

// 	SetSession(c, user.Username)

// 	return c.Redirect().Status(fiber.StatusMovedPermanently).To("/")
// }
