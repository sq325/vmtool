package cluster

import (
	"errors"

	"github.com/sq325/vmtool/pkg/cluster/vmstorage"
)

type Option interface {
	Validate() error
}

type Builder interface {
	Build(Option) (any, error)
}

type StorageBuilder struct{}
type SelectBuilder struct{}
type InsertBuilder struct{}
type AuthBuilder struct{}

var _ Builder = &StorageBuilder{}

func (s *StorageBuilder) Build(opt Option) (any, error) {
	if err := opt.Validate(); err != nil {
		return nil, err
	}

	// 假设 opt 是 vmstorage.Opt 类型
	vmStorageOpt, ok := opt.(vmstorage.Opt)
	if !ok {
		return nil, errors.New("无效的存储配置选项")
	}

	return vmstorage.New(vmStorageOpt), nil
}
