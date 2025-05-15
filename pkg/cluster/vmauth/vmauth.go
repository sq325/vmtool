package vmauth

import (
	"log/slog"
	"os"
)

type Opt struct {
	Path      string
	Addr      string
	ConfigDir string
	Args      []string // 其他参数

	LogDir string
}

func (o Opt) Validate() error {
	return nil
}

// vmauth implements the Instancer interface
type vmauth struct {
	logger *slog.Logger
}

func New(o Opt) *vmauth {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	if o.LogDir != "" {
		// 实际项目中应该设置logger输出到文件
		// 这里仅作为示例
	}

	return &vmauth{
		logger: logger,
	}
}
