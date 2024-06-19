package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

func init() {
	runCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "configs/config.yaml", "config file (default is configs/config.yaml)")
	rootCmd.AddCommand(runCmd, migrateCmd)
}

var rootCmd = &cobra.Command{
	Use:   "nglab",
	Short: "nglab is a Next Generation Lab Information System",
	Long:  `nglab is a Next Generation Lab Information System.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(70)
	}
}
