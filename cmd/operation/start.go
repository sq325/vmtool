package operation

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use: "start",
	// 根据配置文件以 daemon 模式运行整个 cluster
	Short: "Start vm-cluster in daemon mode with config file",
	Long:  `Start vm-cluster in daemon mode with config file`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	startCmd.Flags().String("log.daemon", "./log/daemon.log", "log file")
	startCmd.Flags().String("log.level", "info", "log leve, one of [debug, info, warn, error]")
	startCmd.Flags().String("log.format", "text", "log format, one of [text, json]")

	// logger := logger()

}

func logger() *slog.Logger {
	// 解析命令行参数
	logLevel, _ := startCmd.Flags().GetString("log.level")
	logFormat, _ := startCmd.Flags().GetString("log.format")
	logFilePath, _ := startCmd.Flags().GetString("log.daemon")

	f, err := os.Open(logFilePath)
	if err != nil {
		// 如果文件不存在，则创建文件
		f, err = os.Create(logFilePath)
		if err != nil {
			fmt.Println("无法创建日志文件:", err)
		}
	}
	defer f.Close()

	// 创建日志记录器
	var logger *slog.Logger
	var opt = &slog.HandlerOptions{}
	switch logLevel {
	case "debug":
		opt.Level = slog.LevelDebug
	case "info":
		opt.Level = slog.LevelInfo
	case "warn":
		opt.Level = slog.LevelWarn
	case "error":
		opt.Level = slog.LevelError
	default:
		opt.Level = slog.LevelInfo
	}
	opt.AddSource = true
	switch logFormat {
	case "json":
		logger = slog.New(slog.NewTextHandler(f, opt))
	case "text":
		logger = slog.New(slog.NewTextHandler(f, opt))
	default:
		logger = slog.New(slog.NewTextHandler(f, opt))
	}

	return logger

}
