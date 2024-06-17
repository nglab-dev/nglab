package cmd

import (
	"github.com/nglab-dev/nglab/internal/bootstrap"
	"github.com/nglab-dev/nglab/internal/config"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var (
	cfgFile string
)

func init() {
	runCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "configs/config.yaml", "config file (default is configs/config.yaml)")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a application",
	Long:  `Run a application`,
	Run: func(_ *cobra.Command, _ []string) {
		config.SetConfigFilePath(cfgFile)

		runApp()
	},
}

func runApp() {
	fx.New(bootstrap.Module).Run()
}
