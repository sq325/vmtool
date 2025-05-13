package config

import "github.com/spf13/cobra"

var daemonConfCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Manage daemon configuration",
	Long:  `Manage daemon configuration`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
