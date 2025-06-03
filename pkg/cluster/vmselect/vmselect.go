package vmselect

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
)

type Opt struct {
	Path string // ./vmselect

	Addr              string   // :8481
	StorageNodes      []string // [ip:port ip:port ip:port]
	CacheDir          string   // ./tmp/vmselect
	MinScrapeInterval string   // 15s
	LatencyOffset     string   // 15s
	ReplicationFactor int      // 2
	Args              []string // 其他参数

	LogDir string
}

func (o Opt) Validate() error {

	return nil
}

// vmselect implements the Instancer interface
type vmselect struct {
	cmd *exec.Cmd
	opt Opt
}

// New creates a new vmselect instance with the given options
func New(o Opt) *vmselect {
	cmdPath := o.Path
	cmdArgs := []string{
		"-httpListenAddr=" + o.Addr,
		"-replicationFactor=" + strconv.Itoa(o.ReplicationFactor),
		"-dedup.minScrapeInterval=" + o.MinScrapeInterval,
		"-search.latencyOffset=" + o.LatencyOffset,
		"-cacheDataPath=" + o.CacheDir,
	}
	for _, node := range o.StorageNodes {
		cmdArgs = append(cmdArgs, "-storageNode="+node)
	}
	cmdArgs = append(cmdArgs, o.Args...)

	ctx := context.Background()
	return &vmselect{
		cmd: exec.CommandContext(ctx, cmdPath, cmdArgs...),
		opt: o,
	}
}

func (s *vmselect) Health() error {

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

func (s *vmselect) Port() int {
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

func (s *vmselect) MetricsPath() string {
	return "/metrics"
}

func (s *vmselect) HealthPath() string {
	return "/health"
}

func (v *vmselect) Cmd() *exec.Cmd {
	return v.cmd
}

func (v *vmselect) Name() string {
	return "vmselect"
}
