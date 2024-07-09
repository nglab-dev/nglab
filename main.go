package main

import (
	"flag"

	"github.com/nglab-dev/nglab/internal/conf"
	"github.com/nglab-dev/nglab/internal/db"
	"github.com/nglab-dev/nglab/internal/log"
	"github.com/nglab-dev/nglab/internal/server"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configFile := flag.String("c", "config.yaml", "config file path")
	flag.Parse()

	// load config
	if err := conf.Load(*configFile); err != nil {
		panic(err)
	}

	// init logger
	if err := log.Init(); err != nil {
		panic(err)
	}

	// init database
	if err := db.Init(); err != nil {
		panic(err)
	}

	// run server
	if err := server.Run(); err != nil {
		panic(err)
	}
}
