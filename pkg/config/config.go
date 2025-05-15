package config

import (
	"encoding/json"
	"io"

	"gopkg.in/yaml.v3"
)

type ConfigVisitor func(*Config) error
type Configer interface {
	Visit(ConfigVisitor)
}

// defaultVisitor
func DefaultVisitor(c *Config) error {
	// cluster
	c.ReplicationFactor = 1 // 默认福本数是1

	defaultStorage(c)
	defaultSelect(c)
	defaultInsert(c)
	defaultAuth(c)
	defaultDaemon(c)

	return nil

}

func defaultSelect(c *Config) error {
	// select
	if c.Vmselect.Addr == "" {
		c.Vmselect.Addr = ":8481"
	}
	if c.Vmselect.CacheDir == "" {
		c.Vmselect.CacheDir = "./tmp/vmselect"
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
	return nil
}

func defaultStorage(c *Config) error {

	// storage
	if c.Vmstorage.Addr == "" {
		c.Vmstorage.Addr = ":8482"
	}
	if c.Vmstorage.Retention == "" {
		c.Vmstorage.Retention = "180d"
	}
	if c.Vmstorage.DataDir == "" {
		c.Vmstorage.DataDir = "./vmstorage-data"
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
	return nil
}

func defaultInsert(c *Config) error {
	// insert
	if c.Vminsert.Addr == "" {
		c.Vminsert.Addr = ":8480"
	}
	if c.Vminsert.ReplicationFactor != 0 {
		c.Vminsert.ReplicationFactor = c.ReplicationFactor
	}

	return nil
}

func defaultAuth(c *Config) error {
	// vmauth
	if c.Vmauth.Addr == "" {
		c.Vmauth.Addr = ":8427"
	}
	return nil
}

func defaultDaemon(c *Config) error {
	// log
	if c.LogDir == "" {
		c.LogDir = "./logs"
	}
	if c.LogLevel == "" {
		c.LogLevel = "info"
	}

	// tmp
	if c.TmpDir == "" {
		c.TmpDir = "./tmp"
	}

	// config dir
	if c.ConfigDir == "" {
		c.ConfigDir = "./config"
	}

	// bin dir
	if c.BinDir == "" {
		c.BinDir = "./bin"
	}

	return nil
}

var _ Configer = (*Config)(nil)

type Config struct {
	Nodes             []string    `json:"nodes" yaml:"nodes"` // list of all nodes
	ReplicationFactor int         `json:"replication_factor" yaml:"replication_factor"`
	Vmstorage         StorageConf `json:"vmstorage" yaml:"vmstorage"`
	Vminsert          InsertConf  `json:"vminsert" yaml:"vminsert"`
	Vmselect          SelectConf  `json:"vmselect" yaml:"vmselect"`
	Vmauth            AuthConf    `json:"vmauth" yaml:"vmauth"`

	LogDir   string `json:"log_dir" yaml:"log_dir"`     // directory where logs are stored
	LogLevel string `json:"log_level" yaml:"log_level"` // log level, e.g., "info", "debug", "error"

	BackupDir string `json:"backup_dir" yaml:"backup_dir"`
	TmpDir    string `json:"tmp_dir" yaml:"tmp_dir"`       // temporary directory for operations
	ConfigDir string `json:"config_dir" yaml:"config_dir"` // directory where configuration files are stored
	BinDir    string `json:"bin_dir" yaml:"bin_dir"`       // directory where binaries are located
}

func (c *Config) Visit(opt ConfigVisitor) {
	if err := opt(c); err != nil {
		panic(err)
	}
}

// visitor

// yaml
func (c *Config) YamlMarshal() ([]byte, error) {
	return yaml.Marshal(c)
}

func (c *Config) YamlUnmarshal(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func (c *Config) YamlDecoder(r io.Reader) error {
	decoder := yaml.NewDecoder(r)
	return decoder.Decode(c)
}

// json
func (c *Config) JsonMarshal() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}

func (c *Config) JsonUnmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}

type StorageConf struct {
	Retention         string   `json:"retention" yaml:"retention"`
	Addr              string   `json:"addr" yaml:"addr"`
	DataDir           string   `json:"data_dir" yaml:"data_dir"`
	InsertAddr        string   `json:"insert_addr" yaml:"insert_addr"`
	SelectAddr        string   `json:"select_addr" yaml:"select_addr"`
	MinScrapeInterval string   `json:"min_scrape_interval" yaml:"min_scrape_interval"`
	Args              []string `json:"args" yaml:"args"` // 其他参数
}

type SelectConf struct {
	Addr              string   `json:"addr" yaml:"addr"`
	CacheDir          string   `json:"cache_dir" yaml:"cache_dir"`
	MinScrapeInterval string   `json:"min_scrape_interval" yaml:"min_scrape_interval"`
	LatencyOffset     string   `json:"latency_offset" yaml:"latency_offset"`
	ReplicationFactor int      `json:"-" yaml:"-"`
	StorageAddr       string   `json:"-" yaml:"-"`
	Args              []string `json:"args" yaml:"args"` // 其他参数
}

type InsertConf struct {
	Addr              string   `json:"addr" yaml:"addr"`
	ReplicationFactor int      `json:"replication_factor" yaml:"replication_factor"`
	StorageAddr       string   `json:"-" yaml:"-"`
	Args              []string `json:"args" yaml:"args"` // 其他参数
}

type AuthConf struct {
	Addr      string   `json:"addr" yaml:"addr"`
	ConfigDir string   `json:"config_dir" yaml:"config_dir"`
	Args      []string `json:"args" yaml:"args"` // 其他参数
}

type BackupConf struct {
	BackupDir string   `json:"backup_dir" yaml:"backup_dir"`
	Args      []string `json:"args" yaml:"args"` // 其他参数
}
