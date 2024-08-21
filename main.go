package main

import (
	"github.com/joho/godotenv"
	"github.com/nglab-dev/nglab/internal/router"
	"github.com/nglab-dev/nglab/internal/server"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	r := router.InitRouter()

	srv := server.NewHTTPServer(r)

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
