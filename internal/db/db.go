package db

import (
	"errors"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Driver string `json:"driver" yaml:"driver" env:"DB_DRIVER" default:"sqlite"`
	DSN    string `json:"dsn" yaml:"dsn" env:"DB_DSN" default:"./db.sqlite"`
}

func Init(config Config) (db *gorm.DB, err error) {

	switch config.Driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(config.DSN))
	case "mysql":
		db, err = gorm.Open(mysql.Open(config.DSN))
	case "postgres":
		db, err = gorm.Open(postgres.Open(config.DSN))
	default:
		err = errors.New("unsupported driver: " + config.Driver)
	}

	if err != nil {
		return nil, err
	}
	return db, nil
}
