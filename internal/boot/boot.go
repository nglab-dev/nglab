package boot

import (
	"github.com/nglab-dev/nglab/internal/app"
	"github.com/nglab-dev/nglab/internal/conf"
	"github.com/nglab-dev/nglab/internal/db"
	"github.com/nglab-dev/nglab/internal/log"
	"github.com/nglab-dev/nglab/internal/server"
)

func Bootstrap(cfgFile string) error {

	// load and parse configuration
	c, err := conf.Load(cfgFile)
	if err != nil {
		return err
	}
	err = c.Unmarshal(&app.Config)
	if err != nil {
		return err
	}

	// initialize logger
	app.Log, err = log.Init(app.Config.Log)
	if err != nil {
		return err
	}

	// initialize database
	app.DB, err = db.Init(app.Config.DB)
	if err != nil {
		return err
	}

	// initialize server
	if err := server.Run(app.Config.Server); err != nil {
		return err
	}

	return nil
}
