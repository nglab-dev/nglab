package database

import (
	"log/slog"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"github.com/nglab-dev/nglab/api/model"
	"github.com/nglab-dev/nglab/internal/config"
	"github.com/nglab-dev/nglab/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	ORM *gorm.DB
}

func New(cfg config.Config) Database {
	dsn := cfg.Database.DSN()
	driver := cfg.Database.Dialect

	if driver == "sqlite" {
		dbFileDir := filepath.Dir(cfg.Database.Name)
		if err := utils.MkdirIfNotExist(dbFileDir); err != nil {
			slog.Error("failed to create database directory", "err", err)
		}
	}

	var dialector gorm.Dialector

	switch driver {
	case "mysql":
		dialector = mysql.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(cfg.Database.Name)
	default:
		slog.Error("unknown database driver", "driver", driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect to database", "err", err)
		return Database{}
	}

	return Database{ORM: db}
}

func (db *Database) AutoMigrate() error {
	return db.ORM.AutoMigrate(model.User{})
}
