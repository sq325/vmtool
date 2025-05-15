package config

import (
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestConfig_YamlMarshal(t *testing.T) {
	// Create a test config
	cfg := &Config{
		Nodes:             []string{"node1", "node2"},
		ReplicationFactor: 2,
		Vmstorage: StorageConf{
			Retention:         "30d",
			Addr:              ":9090",
			DataDir:          "/data",
			InsertAddr:        ":8400",
			SelectAddr:        ":8401",
			MinScrapeInterval: "10s",
		},
		Vminsert: InsertConf{
			Addr:              ":8480",
			ReplicationFactor: 2,
			StorageAddr:       ":8400", // This will be excluded in YAML
		},
		Vmselect: SelectConf{
			Addr:              ":8481",
			CacheDir:         "/cache",
			MinScrapeInterval: "15s",
			LatencyOffset:     "5s",
			ReplicationFactor: 2, // This will be excluded in YAML
			StorageAddr:       ":8401", // This will be excluded in YAML
		},
		Vmauth: AuthConf{
			Addr:       ":8427",
			ConfigDir: "/config",
		},
	}

	// Marshal to YAML
	yamlData, err := cfg.YamlMarshal()
	if err != nil {
		t.Fatalf("YamlMarshal failed: %v", err)
	}

	// Check that the YAML output is not empty
	if len(yamlData) == 0 {
		t.Error("YamlMarshal produced empty output")
	}

	// Unmarshal back to verify
	var newCfg Config
	err = yaml.Unmarshal(yamlData, &newCfg)
	if err != nil {
		t.Fatalf("Failed to unmarshal YAML data: %v", err)
	}

	// Compare key fields
	if !reflect.DeepEqual(newCfg.Nodes, cfg.Nodes) {
		t.Errorf("Nodes mismatch: got %v, want %v", newCfg.Nodes, cfg.Nodes)
	}
	if newCfg.ReplicationFactor != cfg.ReplicationFactor {
		t.Errorf("ReplicationFactor mismatch: got %v, want %v", newCfg.ReplicationFactor, cfg.ReplicationFactor)
	}
	if newCfg.Vmstorage.Retention != cfg.Vmstorage.Retention {
		t.Errorf("Vmstorage.Retention mismatch: got %v, want %v", newCfg.Vmstorage.Retention, cfg.Vmstorage.Retention)
	}
	if newCfg.Vmstorage.Addr != cfg.Vmstorage.Addr {
		t.Errorf("Vmstorage.Addr mismatch: got %v, want %v", newCfg.Vmstorage.Addr, cfg.Vmstorage.Addr)
	}
	if newCfg.Vminsert.Addr != cfg.Vminsert.Addr {
		t.Errorf("Vminsert.Addr mismatch: got %v, want %v", newCfg.Vminsert.Addr, cfg.Vminsert.Addr)
	}
	if newCfg.Vminsert.ReplicationFactor != cfg.Vminsert.ReplicationFactor {
		t.Errorf("Vminsert.ReplicationFactor mismatch: got %v, want %v", newCfg.Vminsert.ReplicationFactor, cfg.Vminsert.ReplicationFactor)
	}
	if newCfg.Vmauth.Addr != cfg.Vmauth.Addr {
		t.Errorf("Vmauth.Addr mismatch: got %v, want %v", newCfg.Vmauth.Addr, cfg.Vmauth.Addr)
	}
	if newCfg.Vmauth.ConfigDir != cfg.Vmauth.ConfigDir {
		t.Errorf("Vmauth.ConfigDir mismatch: got %v, want %v", newCfg.Vmauth.ConfigDir, cfg.Vmauth.ConfigDir)
	}

	// Check that fields with json:"-" yaml:"-" were not marshaled
	if newCfg.Vmselect.ReplicationFactor != 0 {
		t.Errorf("Vmselect.ReplicationFactor should not be marshaled: got %v", newCfg.Vmselect.ReplicationFactor)
	}
	if newCfg.Vmselect.StorageAddr != "" {
		t.Errorf("Vmselect.StorageAddr should not be marshaled: got %v", newCfg.Vmselect.StorageAddr)
	}
	if newCfg.Vminsert.StorageAddr != "" {
		t.Errorf("Vminsert.StorageAddr should not be marshaled: got %v", newCfg.Vminsert.StorageAddr)
	}
}
