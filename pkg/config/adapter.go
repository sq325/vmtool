// Package adapter provides adapters for configuration conversion
package config

import (
	"github.com/sq325/vmtool/pkg/cluster/vmauth"
	"github.com/sq325/vmtool/pkg/cluster/vminsert"
	"github.com/sq325/vmtool/pkg/cluster/vmselect"
	"github.com/sq325/vmtool/pkg/cluster/vmstorage"
)

// ConfigAdapter 定义了配置适配器接口，用于转换配置对象
type ConfigAdapter interface {
	AdaptToOpt() any
}

// StorageAdapter 是适配 storage 配置的适配器
type StorageAdapter struct {
	config *Config
}

// NewStorageAdapter 创建一个新的存储配置适配器
func NewStorageAdapter(cfg *Config) *StorageAdapter {
	return &StorageAdapter{
		config: cfg,
	}
}

// AdaptToOpt 将 Config 转换为 vmstorage.Opt
func (a *StorageAdapter) AdaptToOpt() any {
	return vmstorage.Opt{
		Addr:              a.config.Vmstorage.Addr,
		DataDir:           a.config.Vmstorage.DataDir,
		InsertAddr:        a.config.Vmstorage.InsertAddr,
		SelectAddr:        a.config.Vmstorage.SelectAddr,
		MinScrapeInterval: a.config.Vmstorage.MinScrapeInterval,
		RetentionPeriod:   a.config.Vmstorage.Retention,
	}
}

// SelectAdapter 是适配 select 配置的适配器
type SelectAdapter struct {
	config *Config
}

// NewSelectAdapter 创建一个新的查询配置适配器
func NewSelectAdapter(cfg *Config) *SelectAdapter {
	return &SelectAdapter{
		config: cfg,
	}
}

// AdaptToOpt 将 Config 转换为 vmselect.Opt
func (a *SelectAdapter) AdaptToOpt() any {
	return vmselect.Opt{
		Addr:              a.config.Vmselect.Addr,
		StorageNodes:      a.config.Nodes,
		CacheDir:          a.config.Vmselect.CacheDir,
		MinScrapeInterval: a.config.Vmselect.MinScrapeInterval,
		LatencyOffset:     a.config.Vmselect.LatencyOffset,
		ReplicationFactor: a.config.Vmselect.ReplicationFactor,
	}
}

// InsertAdapter 是适配 insert 配置的适配器
type InsertAdapter struct {
	config *Config
}

// NewInsertAdapter 创建一个新的插入配置适配器
func NewInsertAdapter(cfg *Config) *InsertAdapter {
	return &InsertAdapter{
		config: cfg,
	}
}

// AdaptToOpt 将 Config 转换为 vminsert.Opt
func (a *InsertAdapter) AdaptToOpt() any {
	return vminsert.Opt{
		Addr:              a.config.Vminsert.Addr,
		ReplicationFactor: a.config.Vminsert.ReplicationFactor,
		StorageNodes:      a.config.Nodes,
	}
}

// AuthAdapter 是适配 auth 配置的适配器
type AuthAdapter struct {
	config *Config
}

// NewAuthAdapter 创建一个新的认证配置适配器
func NewAuthAdapter(cfg *Config) *AuthAdapter {
	return &AuthAdapter{
		config: cfg,
	}
}

// AdaptToOpt 将 Config 转换为 vmauth.Opt
func (a *AuthAdapter) AdaptToOpt() any {
	return vmauth.Opt{
		Addr:      a.config.Vmauth.Addr,
		ConfigDir: a.config.Vmauth.ConfigDir,
	}
}

// AdapterFactory 用于创建配置适配器
type AdapterFactory struct {
	config *Config
}

// NewAdapterFactory 创建一个新的适配器工厂
func NewAdapterFactory(cfg *Config) *AdapterFactory {
	return &AdapterFactory{
		config: cfg,
	}
}

// CreateAdapter 根据组件类型创建相应的适配器
func (f *AdapterFactory) CreateAdapter(componentType string) ConfigAdapter {
	switch componentType {
	case "storage":
		return NewStorageAdapter(f.config)
	case "select":
		return NewSelectAdapter(f.config)
	case "insert":
		return NewInsertAdapter(f.config)
	case "auth":
		return NewAuthAdapter(f.config)
	default:
		return nil
	}
}
