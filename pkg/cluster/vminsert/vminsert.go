package vminsert

import (
	"log/slog"
	"os"
)

type Opt struct {
	Path              string
	Addr              string   // :8480
	ReplicationFactor int      // 2
	StorageNodes      []string // [ip:port ip:port ip:port]
	Args              []string // 其他参数

	LogDir string
}

func (o Opt) Validate() error {
	return nil
}

type vminsert struct {
	logger *slog.Logger
}

func New(o Opt) *vminsert {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	if o.LogDir != "" {
		// 实际项目中应该设置logger输出到文件
		// 这里仅作为示例
	}

	return &vminsert{
		logger: logger,
	}
}
