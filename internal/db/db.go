package db

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/nglab-dev/nglab/pkg/env"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (db *gorm.DB, err error) {
	driver := env.GetString("DB_DRIVER", "sqlite")
	dsn := env.GetString("DB_DSN", "db.sqlite")

	if driver == "sqlite" {
		// if path is not exist, create it
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
		return nil, err
	}

	return db, nil
}
