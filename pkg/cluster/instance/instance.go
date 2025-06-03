package instance

import "os/exec"

type Instancer interface {
	Cmd() *exec.Cmd
	Name() string // vmstorage, vmselect
	LogDir() string
	// common admin methods
	Manager
}

type Manager interface {
	Health() error
	Port() int
	MetricsPath() string // if "" means no metrics path
	HealthPath() string  // if "" means no health path
}
