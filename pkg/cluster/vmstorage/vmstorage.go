package vmstorage

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/sq325/vmtool/pkg/cluster/instance"
)

type Opt struct {
	Path string // ./vmstorage

	Addr              string   // :8482
	InsertAddr        string   // :8400
	SelectAddr        string   // :8401
	MinScrapeInterval string   // 15s
	DataDir           string   // ./vmstorage-data
	RetentionPeriod   string   // 180d
	Args              []string // 其他参数 ["key1=value1", "key2=value2"]

	LogDir string // 日志目录
}

func (o Opt) Validate() error {
	if o.Path == "" {
		return errors.New("vmstorage Path is required")
	}
	if o.Addr == "" {
		return errors.New("vmstorage Addr is required")
	}
	if o.InsertAddr == "" {
		return errors.New("vmstorage InsertAddr is required")
	}
	if o.SelectAddr == "" {
		return errors.New("vmstorage SelectAddr is required")
	}
	if o.MinScrapeInterval == "" {
		return errors.New("vmstorage MinScrapeInterval is required")
	}
	if o.DataDir == "" {
		return errors.New("vmstorage DataDir is required")
	}
	if o.RetentionPeriod == "" {
		return errors.New("vmstorage RetentionPeriod is required")
	}

	return nil
}

// storage implements the Instancer interface
type storage struct {
	cmd *exec.Cmd
	opt Opt
}

var _ instance.Instancer = &storage{}

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

	cmdArgs = append(cmdArgs, o.Args...)
	ctx := context.Background()

	return &storage{
		cmd: exec.CommandContext(ctx, cmdPath, cmdArgs...),
		opt: o,
	}
}

func (s *storage) Health() error {

	// process exited
	if s.cmd != nil && s.cmd.ProcessState != nil && s.cmd.ProcessState.Exited() {
		return errors.New("vmstorage process has exited")
	}

	// /health
	u, _ := url.Parse("http://localhost" + s.opt.Addr)
	u = u.JoinPath(s.HealthPath())
	http.DefaultClient.Timeout = 5 * time.Second // 设置超时时间为5秒
	resp, err := http.Get(u.String())            // 发送健康检查请求
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if string(body) != "ok" {
		return errors.New("vmstorage health check failed")
	}

	return nil
}

func (s *storage) Port() int {
	if port, err := strconv.Atoi(s.opt.Addr); err == nil {
		return port // 如果 Addr 是端口号，直接返回
	}

	// 如果 Addr 是形如 ":8482" 的格式，提取端口号
	if strings.HasPrefix(s.opt.Addr, ":") {
		portStr := strings.TrimPrefix(s.opt.Addr, ":")
		if port, err := strconv.Atoi(portStr); err == nil {
			return port // 返回提取的端口号
		}
	}

	return 0 // 如果没有找到端口，返回 0 或者其他默认值
}

func (s *storage) MetricsPath() string {
	return "/metrics"
}

func (s *storage) HealthPath() string {
	return "/health"
}

func (s *storage) Cmd() *exec.Cmd {
	if s.cmd == nil {
		return nil // 如果 Cmd 为 nil，返回 nil
	}
	return s.cmd
}

func (s *storage) Name() string {
	return "vmstorage"
}

func (s *storage) LogDir() string {
	return s.opt.LogDir
}
