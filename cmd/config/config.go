package config

import (
	"github.com/spf13/cobra"
	"github.com/sq325/vmtool/pkg/config"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "生成默认配置文件",
	Long:  `生成默认配置文件`,
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("config.format")
		c := &config.Config{}
		if err := config.DefaultOpt(c); err != nil {
			panic(err)
		}
		switch format {
		case "json":
			if bys, err := config.JsonMarshal(c); err != nil {
				panic(err)
			} else {
				println(string(bys))
			}
		case "yaml":
			if bys, err := config.YamlMarshal(c); err != nil {
				panic(err)
			} else {
				println(string(bys))
			}
		}
	},
}

func init() {
	ConfigCmd.Flags().String("config.format", "yaml", "config file format (default is yaml), one of [yaml, json]")

	ConfigCmd.AddCommand(vmConfCmd)
	ConfigCmd.AddCommand(daemonConfCmd)
}
