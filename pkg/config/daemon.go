package config

type DaemonConf struct {
	PidFile string `json:"pid_file" yaml:"pid_file"`
	LogDir  string `json:"log_dir" yaml:"log_dir"`
}
