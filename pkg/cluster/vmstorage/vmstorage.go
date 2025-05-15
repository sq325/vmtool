package vmstorage

import (
	"context"
	"log/slog"
	"os/exec"
	"strings"

	"github.com/sq325/cmdDaemon/daemon"
)

type Opt struct {
	Path string // ./vmstorage

	Addr              string   // :8482
	InsertAddr        string   // :8400
	SelectAddr        string   // :8401
	MinScrapeInterval string   // 15s
	DataDir           string   // ./vmstorage-data
	RetentionPeriod   string   // 180d
	Args              []string // 其他参数
}

func (Opt) Validate() error {
	// 这里可以添加必要的验证逻辑
	// 例如检查地址格式、目录是否存在等
	return nil
}

// storage implements the Instancer interface
type storage struct {
	logger *slog.Logger
	dcmd   *daemon.DaemonCmd
}

func New(o Opt) *storage {
	cmdPath := o.Path
	cmdArgs := []string{
		// port
		"-httpListenAddr=" + o.Addr,
		"-vminsertAddr=" + o.InsertAddr,
		"-vmselectAddr=" + o.SelectAddr,

		"-dedup.minScrapeInterval=" + o.MinScrapeInterval,
		"-storageDataPath=" + o.DataDir,
		"-retentionPeriod=" + o.RetentionPeriod,
	}
	cmd := exec.Command(cmdPath, strings.Join(cmdArgs, " "))

	ctx := context.Background()
	return &storage{
		dcmd: daemon.NewDaemonCmd(ctx, cmd),
	}
}

func (s *storage) Start() error {
	return nil
}

func (s *storage) Stop() error {
	return nil
}

func (s *storage) MetricsPath() string {
	return "/metrics"
}
