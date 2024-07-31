package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/nglab-dev/nglab/auth"
	"github.com/nglab-dev/nglab/models"
	"github.com/nglab-dev/nglab/utils/env"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Username        string `json:"username" form:"username"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword"`
}

type LoginRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Remember bool   `json:"remember" form:"remember"`
}

type authHandler struct {
	BaseHandler
}

func NewAuthHandler() authHandler {
	return authHandler{}
}

func (h authHandler) Register(router fiber.Router) {
	router.Get("/login", h.login)
	router.Post("/login", h.login)
	router.Get("/register", h.register)
	router.Post("/register", h.register)
	router.Post("/logout", h.logout)
}

func (h authHandler) login(c fiber.Ctx) error {
	appName := env.GetString("APP_NAME", "nglab")
	if c.Method() == "GET" {
		user := auth.GetUser(c)
		if user != nil {
			return c.Redirect().To("/")
		}

		return c.Render("login", fiber.Map{
			"title": appName,
		})
	}
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := models.GetUserByUsername(username)
	if err != nil {
		return h.Error(c, "User not found")
	}

	if err := user.ComparePassword(password); err != nil {
		return h.Error(c, "Incorrect password")
	}

	// save session
	auth.SetSession(c, &user)

	return h.Ok(c, fiber.Map{
		"location": "/",
	})
}

func (h authHandler) register(c fiber.Ctx) error {
	appName := env.GetString("APP_NAME", "nglab")
	if c.Method() == "GET" {
		return c.Render("register", fiber.Map{
			"title": appName,
		})
	}

	req := new(RegisterRequest)
	if err := c.Bind().Form(req); err != nil {
		return h.Error(c, "Invalid input")
	}

	if req.Password != req.ConfirmPassword {
		return h.Error(c, "Passwords do not match")
	}

	user, err := models.GetUserByUsername(req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return h.Error(c, err.Error())
	}

	if user.ID != 0 {
		return h.Error(c, "Username already exists")
	}

	newUser := new(models.User)
	newUser.Username = req.Username
	newUser.Password = req.Password

	err = models.CreateUser(newUser)
	if err != nil {
		return h.Error(c, err.Error())
	}

	return h.Ok(c, fiber.Map{
		"location": "/login",
	})
}

func (h authHandler) logout(c fiber.Ctx) error {
	auth.ClearSession(c)
	return h.Ok(c, nil)
}
