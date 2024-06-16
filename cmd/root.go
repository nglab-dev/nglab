package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
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
