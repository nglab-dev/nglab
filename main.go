package main

import (
	"flag"

	"github.com/nglab-dev/nglab/internal/conf"
	"github.com/nglab-dev/nglab/internal/model"
	"github.com/nglab-dev/nglab/internal/query"
	"github.com/nglab-dev/nglab/internal/router"
	"github.com/xbmlz/ungo/unconf"
	"github.com/xbmlz/ungo/undb"
	"github.com/xbmlz/ungo/unhttp"
	"github.com/xbmlz/ungo/unlog"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configFile := flag.String("c", "config.yaml", "config file path")
	flag.Parse()

	cfg, err := unconf.New(*configFile)
	if err != nil {
		panic(err)
	}

	config := &conf.Config{}
	cfg.Parse(config)

	// initialize db
	db, err := undb.New(config.DB)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})

	query.SetDefault(db)

	srv := unhttp.NewServer(config.Server)

	router.RegisterRoutes(srv.Router)

	unlog.Infof("Starting server on %s:%d", config.Server.Host, config.Server.Port)

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
