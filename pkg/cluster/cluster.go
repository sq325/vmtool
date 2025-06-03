package cluster

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sq325/cmdDaemon/daemon"
	"github.com/sq325/vmtool/pkg/cluster/instance"
)

type operator struct {
}

// hook
func (operator) PreStart(ins instance.Instancer) error {
	cmd := ins.Cmd()
	logdir := ins.LogDir()
	if logdir != "" {
		f, err := os.Open(filepath.Join(logdir, ins.Name()+".log"))
		if err != nil {
			cmd.Stdout = f
			cmd.Stderr = f
		}
	}

	return nil
}

func (operator) Start(ins instance.Instancer) error {

	return nil
}

// hook
func (operator) PostStart(ins instance.Instancer) error {
	return nil
}

func (operator) Stop(ins instance.Instancer) error {
	return nil
}

func (operator) Status(ins instance.Instancer) error {
	return nil
}

func (operator) Restart(ins instance.Instancer) error {
	return nil
}

type Instrumenter interface {
	WithLogger(logger *slog.Logger)
	WithMetricsRegistry(registry prometheus.Registerer)
}

type Cluster interface {
	Start() error
}

type vmcluster struct {
	d *daemon.Daemon
}
