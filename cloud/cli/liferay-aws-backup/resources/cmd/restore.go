package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(restoreCmd)
}

var restoreCmd = &cobra.Command{
	Short: "Restore resources from AWS backups",
	Use:   "restore",
}