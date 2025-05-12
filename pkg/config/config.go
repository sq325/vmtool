package config

type ConfigOpt func(*Config) error
type Visitor interface {
	Visit(ConfigOpt)
}

// defaultOpt sets default values for the Config if not explicitly specified.
// It populates the configuration with the following defaults:
//
// Cluster defaults:
//   - LogDir: "./logs"
//   - PidFile: "./victoriametrics.pid"
//   - ReplicationFactor: 1
//
// Storage defaults:
//   - Vmstorage.Addr: ":8482"
//   - Vmstorage.Retention: "180d"
//   - Vmstorage.DataPath: "./vmstorage-data"
//   - Vmstorage.InsertAddr: ":8400"
//   - Vmstorage.SelectAddr: ":8401"
//   - Vmstorage.MinScrapeInterval: "15s"
//
// Select defaults:
//   - Vmselect.CachePath: "./tmp/vmselect"
//   - Vmselect.MinScrapeInterval: "15s"
//   - Vmselect.LatencyOffset: same as MinScrapeInterval
//   - Vmselect.ReplicationFactor: same as global ReplicationFactor
//
// Insert defaults:
//   - Vminsert.Addr: ":8480"
//   - Vminsert.ReplicationFactor: same as global ReplicationFactor
//
// Auth defaults:
//   - Vmauth.Addr: ":8427"
func defaultOpt(c *Config) error {
	// cluster
	c.LogDir = "./logs"
	c.PidFile = "./victoriametrics.pid"
	c.ReplicationFactor = 1 // 默认福本数是1

	// storage
	if c.Vmstorage.Addr == "" {
		c.Vmstorage.Addr = ":8482"
	}
	if c.Vmstorage.Retention == "" {
		c.Vmstorage.Retention = "180d"
	}
	if c.Vmstorage.DataPath == "" {
		c.Vmstorage.DataPath = "./vmstorage-data"
	}
	if c.Vmstorage.InsertAddr == "" {
		c.Vmstorage.InsertAddr = ":8400"
		c.Vminsert.StorageAddr = c.Vmstorage.InsertAddr
	}
	if c.Vmstorage.SelectAddr == "" {
		c.Vmstorage.SelectAddr = ":8401"
		c.Vmselect.StorageAddr = c.Vmstorage.SelectAddr
	}
	if c.Vmstorage.MinScrapeInterval == "" {
		c.Vmstorage.MinScrapeInterval = "15s"
	}

	// select
	if c.Vmselect.CachePath == "" {
		c.Vmselect.CachePath = "./tmp/vmselect"
	}
	if c.Vmselect.MinScrapeInterval == "" {
		c.Vmselect.MinScrapeInterval = "15s"
	}
	if c.Vmselect.LatencyOffset == "" {
		c.Vmselect.LatencyOffset = c.Vmselect.MinScrapeInterval
	}
	if c.ReplicationFactor != 0 {
		c.Vmselect.ReplicationFactor = c.ReplicationFactor
	}

	// insert
	if c.Vminsert.Addr == "" {
		c.Vminsert.Addr = ":8480"
	}
	if c.Vminsert.ReplicationFactor != 0 {
		c.Vminsert.ReplicationFactor = c.ReplicationFactor
	}

	// vmauth
	if c.Vmauth.Addr == "" {
		c.Vmauth.Addr = ":8427"
	}
	return nil

}

var _ Visitor = (*Config)(nil)

type Config struct {
	Nodes             []string    `json:"nodes" yaml:"nodes"` // list of all nodes
	Vmstorage         StorageConf `json:"vmstorage" yaml:"vmstorage"`
	Vminsert          InsertConf  `json:"vminsert" yaml:"vminsert"`
	Vmselect          SelectConf  `json:"vmselect" yaml:"vmselect"`
	Vmauth            AuthConf    `json:"vmauth" yaml:"vmauth"`
	ReplicationFactor int         `json:"replication_factor" yaml:"replication_factor"`
	LogDir            string      `json:"log_dir" yaml:"log_dir"`
	PidFile           string      `json:"pid_file" yaml:"pid_file"`
}

type StorageConf struct {
	Retention         string `json:"retention" yaml:"retention"`
	Addr              string `json:"addr" yaml:"addr"`
	DataPath          string `json:"data_path" yaml:"data_path"`
	InsertAddr        string `json:"insert_addr" yaml:"insert_addr"`
	SelectAddr        string `json:"select_addr" yaml:"select_addr"`
	MinScrapeInterval string `json:"min_scrape_interval" yaml:"min_scrape_interval"`
}

type SelectConf struct {
	Addr              string `json:"addr" yaml:"addr"`
	CachePath         string `json:"cache_path" yaml:"cache_path"`
	MinScrapeInterval string `json:"min_scrape_interval" yaml:"min_scrape_interval"`
	LatencyOffset     string `json:"latency_offset" yaml:"latency_offset"`
	ReplicationFactor int    `json:"-" yaml:"-"`
	StorageAddr       string `json:"-" yaml:"-"`
}

type InsertConf struct {
	Addr              string `json:"addr" yaml:"addr"`
	ReplicationFactor int    `json:"replication_factor" yaml:"replication_factor"`
	StorageAddr       string `json:"-" yaml:"-"`
}

type AuthConf struct {
	Addr       string `json:"addr" yaml:"addr"`
	ConfigPath string `json:"config_path" yaml:"config_path"`
}

type BackupConf struct {
	BackupPath string `json:"backup_path" yaml:"backup_path"`
}

func (c *Config) Visit(opt ConfigOpt) {
	if err := opt(c); err != nil {
		panic(err)
	}
}
