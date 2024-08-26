package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/nglab-dev/nglab/internal/config"
	"github.com/nglab-dev/nglab/internal/router"
	"github.com/nglab-dev/nglab/internal/server"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {

	c := config.Load()

	r := router.InitRouter(c)

	srv := server.NewHTTPServer(c, r)

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
