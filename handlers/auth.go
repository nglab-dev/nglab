package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/nglab-dev/nglab/models"
	"gorm.io/gorm"
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

func HandleLoginView(c fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}

func HandleRegisterView(c fiber.Ctx) error {
	return c.Render("register", fiber.Map{})
}

func HandleRegister(c fiber.Ctx) error {
	req := new(RegisterRequest)

	if err := c.Bind().Form(req); err != nil {
		return c.Render("signup", fiber.Map{
			"error": "Invalid input",
		})
	}

	if req.Password != req.ConfirmPassword {
		return c.Render("signup", fiber.Map{
			"error": "Passwords do not match",
		})
	}

	user, err := models.GetUserByUsername(req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		return c.Render("signup", fiber.Map{
			"error": err.Error(),
		})
	}

	if user.ID != 0 {
		log.Infof("user id is %s", user.ID)
		return c.Render("signup", fiber.Map{
			"error": "Username already exists",
		})
	}

	newUser := new(models.User)
	newUser.Username = req.Username
	newUser.Password = req.Password

	err = models.CreateUser(newUser)
	if err != nil {
		return c.Render("signup", fiber.Map{
			"error": err.Error(),
		})
	}
	// to home
	return c.Redirect().Status(fiber.StatusMovedPermanently).To("/login")
}

func HandleLogin(c fiber.Ctx) error {
	req := new(LoginRequest)
	if err := c.Bind().Form(req); err != nil {
		return c.Render("login", fiber.Map{
			"error": "",
		})
	}

	user, err := models.GetUserByUsername(req.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Render("login", fiber.Map{
				"error": "User not found",
			})
		}
	}

	if err := user.ComparePassword(req.Password); err != nil {
		return c.Render("login", fiber.Map{
			"error": "Incorrect password",
		})
	}

	session, err := store.Get(c)
	if err != nil {
		return c.Render("login", fiber.Map{
			"error": "Error getting session",
		})
	}

	session.Set(AUTH_KEY, true)
	session.Set(USER_ID, user.ID)

	err = session.Save()
	if err != nil {
		return c.Render("login", fiber.Map{
			"error": "Error saving session",
		})
	}

	return c.Redirect().Status(fiber.StatusMovedPermanently).To("/")
}
