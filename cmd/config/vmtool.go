package config

import "github.com/spf13/cobra"

var vmConfCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Manage cluster configuration",
	Long:  `Manage cluster configuration`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
