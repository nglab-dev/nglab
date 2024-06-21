package cmd

import (
	"github.com/nglab-dev/nglab/internal/config"
	"github.com/nglab-dev/nglab/internal/storage"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Long:  `Run database migrations`,
	Run: func(_ *cobra.Command, _ []string) {
		config.SetConfigFilePath(cfgFile)
		runMigrateCmd()
	},
}

func runMigrateCmd() {
	cfg := config.New()

	storage := storage.New(cfg)
	sqlDB, err := storage.DB.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	if err = storage.AutoMigrate(); err != nil {
		panic(err)
	}
}
