package db

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/nglab-dev/nglab/pkg/env"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	var dialector gorm.Dialector

	switch driver {
	case "sqlite":
		dialector = sqlite.Open(dsn)
	case "mysql":
		dialector = mysql.Open(dsn)
	case "postgres":
		dialector = postgres.Open(dsn)
	default:
		return nil, errors.New("unsupported driver")
	}

	db, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				ParameterizedQueries:      false,       // Don't include params in the SQL log
				Colorful:                  false,       // Disable color
			},
		),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
