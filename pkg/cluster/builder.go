package cluster

import (
	"github.com/sq325/vmtool/pkg/cluster/instance"
	"github.com/sq325/vmtool/pkg/cluster/vminsert"
	"github.com/sq325/vmtool/pkg/cluster/vmselect"
	"github.com/sq325/vmtool/pkg/cluster/vmstorage"
)

// type Option interface {
// 	Validate() error
// }

type Option interface {
	vmstorage.Opt | vmselect.Opt | vminsert.Opt
}

type Builder[T Option] interface {
	Build(T) (instance.Instancer, error)
}

type StorageBuilder struct{}
type SelectBuilder struct{}
type InsertBuilder struct{}
type AuthBuilder struct{}

var _ Builder[vmstorage.Opt] = &StorageBuilder{}

func (s *StorageBuilder) Build(opt vmstorage.Opt) (instance.Instancer, error) {
	if err := opt.Validate(); err != nil {
		return nil, err
	}

	return vmstorage.New(opt), nil
}
