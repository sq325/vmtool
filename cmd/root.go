package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sq325/vmtool/cmd/operation"
	"github.com/sq325/vmtool/pkg/config"
)

var (
	_versionInfo   string
	buildTime      string
	buildGoVersion string
	_version       string
	author         string
	projectName    string
)

func init() {
	RootCmd.Flags().BoolP("version", "v", false, "show version info")
	RootCmd.PersistentFlags().String("config", "./config/vmtool.yml", "config file (default is vmtool.yml)")
	RootCmd.PersistentFlags().String("config.format", "yaml", "config file format (default is yaml), one of [yaml, json]")

	RootCmd.AddCommand(operation.DeployCmd)
	RootCmd.AddCommand(ConfigCmd)
}

var RootCmd = &cobra.Command{
	Use:   "vmtool",
	Short: "A CLI tool to manage vm-clsuter",
	Long:  `A CLI tool to manage vm-clsuter`,
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion, err := cmd.Flags().GetBool("version"); err != nil {
			fmt.Println("failed to get version flag:", err)
			return
		} else if showVersion {
			fmt.Println(projectName, _version)
			fmt.Println("build time:", buildTime)
			fmt.Println("go version:", buildGoVersion)
			fmt.Println("author:", author)
			fmt.Println("version info:", _versionInfo)
			return
		}

		// ctx 传递给接下来的子命令

		cmd.Help()
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() error {

	// config 初始化

	configFilePath, err := RootCmd.Flags().GetString("config")
	if err != nil {
		return fmt.Errorf("failed to get config file path: %w", err)
	}
	c := getConfig(configFilePath)
	if c == nil {
		return fmt.Errorf("failed to get config")
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, ConfigKey, c)
	_, err = RootCmd.ExecuteContextC(ctx)
	return err
}

func getConfig(filePath string) *config.Config {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("failed to open config file:", err)
		return nil
	}
	defer f.Close()

	c := &config.Config{}
	if err := c.YamlDecoder(f); err != nil {
		fmt.Println("failed to unmarshal config file:", err)
		return nil
	}
	return c
}

type Contextkey string

const (
	ConfigKey Contextkey = "config"
)
