package db

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"github.com/nglab-dev/nglab/utils/env"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Get() *gorm.DB {
	return db
}

func Init() (err error) {
	driver := env.GetString("DB_DRIVER", "sqlite")
	dsn := env.GetString("DB_URL", "./db.sqlite")

	if driver == "sqlite" {
		if _, err := os.Stat(dsn); os.IsNotExist(err) {
			os.MkdirAll(filepath.Dir(dsn), os.ModePerm)
		}
	}

	switch driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn))
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn))
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn))
	default:
		err = errors.New("unsupported driver")
	}

	if err != nil {
		return err
	}

	return nil
}
