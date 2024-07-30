package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/nglab-dev/nglab/db"
	"github.com/nglab-dev/nglab/models"
	"github.com/nglab-dev/nglab/router"
	"github.com/nglab-dev/nglab/utils/env"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := db.Init(); err != nil {
		panic(err)
	}

	db.Get().AutoMigrate(&models.User{})

	addr := env.GetString("SERVER_ADDR", ":3000")

	engine := html.NewFileSystem(http.Dir("./views"), ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	router.Setup(app)

	log.Fatal(app.Listen(addr))
}
