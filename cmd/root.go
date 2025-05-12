package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	_versionInfo   string
	buildTime      string
	buildGoVersion string
	_version       string
	author         string
	projectName    string
)

var RootCmd = &cobra.Command{
	Use:   "vmtool",
	Short: "A CLI tool to manage vm-clsuter",
	Long:  `A CLI tool to manage vm-clsuter`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println(projectName, _version)
			fmt.Println("build time:", buildTime)
			fmt.Println("go version:", buildGoVersion)
			fmt.Println("author:", author)
			fmt.Println("version info:", _versionInfo)
			return
		}

		cmd.Help()
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() error {
	return RootCmd.Execute()
}

var (
	version bool
)

func init() {
	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "show version info")
}
