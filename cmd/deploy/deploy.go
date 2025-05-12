package deploy

import "github.com/spf13/cobra"

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy vm-cluster",
	Long:  `Deploy vm-cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
