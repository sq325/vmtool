package cluster

type Instance interface {
	Start() error
	Stop() error
	Restart() error
}

type Cluster interface {
	Instance
	Backup() error
	Restore() error
	List() error
	Cmds() error
}

type cluster struct{}
