package cluster

import (
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
)

type Instancer interface {
	WithLogger(logger *slog.Logger) Instancer
	WithMetricsRegistry(registry prometheus.Registerer) Instancer

	Start() error
	Stop() error
	Restart() error

	Reload() error
	Health() error
	Port() int
	MetricsPath() string
	HealthPath() string

	Path() string // /app/bin/vmstorage
	Cmd() (string, error)
}

type Cluster interface {
	Backup() error
	Restore() error
	List() error

	Reload() error
	Health() error

	Cmds() error
}

type cluster struct{}
