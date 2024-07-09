package db

import (
	"errors"

	"github.com/glebarez/sqlite"
	"github.com/nglab-dev/nglab/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func Get() *gorm.DB {
	return dbInstance
}

func Init() (err error) {

	config := conf.Get().DB

	switch config.Driver {
	case "sqlite":
		dbInstance, err = gorm.Open(sqlite.Open(config.DSN))
	case "mysql":
		dbInstance, err = gorm.Open(mysql.Open(config.DSN))
	case "postgres":
		dbInstance, err = gorm.Open(postgres.Open(config.DSN))
	default:
		err = errors.New("unsupported driver: " + config.Driver)
	}

	if err != nil {
		return err
	}
	return nil
}
