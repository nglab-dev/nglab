package app

import (
	"github.com/nglab-dev/nglab/internal/conf"
	"github.com/nglab-dev/nglab/internal/db"
	"github.com/nglab-dev/nglab/internal/log"
	"github.com/nglab-dev/nglab/internal/server"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppConfig struct {
	Log    log.Config    `json:"log" yaml:"log"`
	DB     db.Config     `json:"db" yaml:"db"`
	Server server.Config `json:"server" yaml:"server"`
}

var (
	Config *AppConfig
	DB     *gorm.DB
	Log    *zap.Logger
)

func Bootstrap(cfgFile string) error {

	// load and parse configuration
	c, err := conf.Load(cfgFile)
	if err != nil {
		return err
	}
	err = c.Unmarshal(&Config)
	if err != nil {
		return err
	}

	// initialize logger
	Log, err = log.Init(Config.Log)
	if err != nil {
		return err
	}

	// initialize database
	DB, err = db.Init(Config.DB)
	if err != nil {
		return err
	}

	// initialize server

	return nil
}
