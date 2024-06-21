package storage

import (
	"log/slog"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"github.com/nglab-dev/nglab/api/model"
	"github.com/nglab-dev/nglab/internal/config"
	"github.com/nglab-dev/nglab/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func New(cfg config.Config) Storage {
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
		return Storage{}
	}

	return Storage{db}
}

func (db *Storage) AutoMigrate() error {
	return db.DB.AutoMigrate(model.User{})
}
