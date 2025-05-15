package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/sq325/vmtool/pkg/config"
)

// ConfigCmd represents the command to generate a default configuration file
// in either JSON or YAML format.
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "生成默认配置文件",
	Long: `生成默认配置文件
该命令会在指定路径生成默认的配置文件，支持JSON和YAML格式。`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取命令行参数
		format, err := cmd.Flags().GetString("config.format")
		if err != nil {
			fmt.Fprintf(os.Stderr, "获取配置格式失败: %v\n", err)
			return
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Fprintf(os.Stderr, "获取输出路径失败: %v\n", err)
			return
		}

		// 创建并初始化配置对象
		c := &config.Config{}
		if err := config.DefaultVisitor(c); err != nil {
			fmt.Fprintf(os.Stderr, "初始化默认配置失败: %v\n", err)
			return
		}

		// 根据格式生成配置文件
		var configData []byte
		switch format {
		case "json":
			configData, err = c.JsonMarshal()
		case "yaml":
			configData, err = c.YamlMarshal()
		default:
			fmt.Fprintf(os.Stderr, "不支持的配置格式: %s，请使用 json 或 yaml\n", format)
			return
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "序列化配置失败: %v\n", err)
			return
		}

		// 写入配置文件
		if err := writeToFile(output, string(configData)); err != nil {
			fmt.Fprintf(os.Stderr, "写入配置文件失败: %v\n", err)
			return
		}

		fmt.Printf("配置文件已生成: %s\n", output)
	},
}

func init() {
	// 注册命令行参数
	ConfigCmd.Flags().String("output", "./config/vmtool.yml", "配置文件路径")
}

// writeToFile writes content to a file at the specified path.
// If the file doesn't exist, it creates the file (and its directory if needed).
// If the file exists, it returns nil without modifying the file.
func writeToFile(path string, content string) error {
	// 确保目录存在
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	// 检查文件是否存在
	_, err := os.Stat(path)
	if err == nil {
		// 文件已存在
		return fmt.Errorf("文件已存在: %s", path)
	}

	// 创建并写入文件
	if os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			return fmt.Errorf("创建文件失败: %w", err)
		}
		defer f.Close()

		if _, err := io.WriteString(f, content); err != nil {
			return fmt.Errorf("写入内容失败: %w", err)
		}
		return nil
	}

	return err
}
