package main

import (
	"context"
	"flag"

	"github.com/nglab-dev/nglab/internal/boot"
	"github.com/nglab-dev/nglab/internal/conf"
	"github.com/xbmlz/ungo"
	"github.com/xbmlz/ungo/cfg"
	"github.com/xbmlz/ungo/log"
	"github.com/xbmlz/ungo/server"
)

var configFile = flag.String("c", "config.yaml", "config file path")

func main() {
	flag.Parse()

	var config conf.Config
	cfg.MustLoad(*configFile, &config)

	app := ungo.NewApp(
		ungo.WithServer(
			server.NewHTTPServer(boot.NewGinRouter(), &config.Server),
		),
	)

	log.Infof("Starting server on %s", config.Server.Addr())

	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}
