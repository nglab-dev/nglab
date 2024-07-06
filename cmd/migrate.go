package cmd

import (
	"github.com/nglab-dev/nglab/api/model"
	"github.com/nglab-dev/nglab/internal/config"
	"github.com/nglab-dev/nglab/internal/database"
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

	db := database.New(cfg)

	db.DB.AutoMigrate(
		&model.User{},
		&model.Menu{},
	)
	// Add more models here as needed
}
