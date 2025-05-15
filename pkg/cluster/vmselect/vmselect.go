package vmselect

import (
	"context"
	"log/slog"
	"os/exec"
	"strconv"
	"strings"

	"github.com/sq325/cmdDaemon/daemon"
)

type Opt struct {
	Path              string   // ./vmselect
	Addr              string   // :8481
	StorageNodes      []string // [ip:port ip:port ip:port]
	CacheDir          string   // ./tmp/vmselect
	MinScrapeInterval string   // 15s
	LatencyOffset     string   // 15s
	ReplicationFactor int      // 2
	Args              []string // 其他参数
	LogDir            string
}

func (o Opt) Validate() error {

	return nil
}

// vmselect implements the Instancer interface
type vmselect struct {
	logger *slog.Logger
	dcmd   *daemon.DaemonCmd
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
	cmd := exec.Command(cmdPath, strings.Join(cmdArgs, " "))
	ctx := context.Background()
	return &vmselect{
		dcmd: daemon.NewDaemonCmd(ctx, cmd),
	}
}

// Start starts the vmselect service
func (s *vmselect) Start() error {

	// 实现启动逻辑
	return nil
}

// Stop stops the vmselect service
func (s *vmselect) Stop() error {
	s.logger.Info("Stopping vmselect service")
	// 实现停止逻辑
	return nil
}
