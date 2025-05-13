package cluster

type Instance interface {
	Start() error
	Stop() error
	Restart() error
	Reload() error
	Cmd() (string, error)
}

type Cluster interface {
	Backup() error
	Restore() error
	List() error
	Cmds() error
}

type cluster struct{}
