package asset

import (
	"context"

	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services"
)

type HostGetter interface {
	Host() IHost
}

type host struct {
	factory services.FactoryImp
}

func NewHost(f services.FactoryImp) IHost {
	return &host{
		factory: f,
	}
}

// IHost 主机操作接口
type IHost interface {
	Create(ctx context.Context, host *asset.Host)
	Delete(ctx context.Context, hostId uint64)
	Update(ctx context.Context, hostId uint64, host *asset.Host)
	List(ctx context.Context, page, limit int) *form.PageHost
	GetByHostId(ctx context.Context, hostId uint64) (host *asset.Host)
	UpdateFiledStatus(ctx context.Context, hostId uint64, updateFiled string, status int8)
	CheckHostExist(ctx context.Context, hostId uint64) bool
}

func (m *host) Create(ctx context.Context, host *asset.Host) {
	m.factory.Host().Create(ctx, host)
}

func (m *host) Delete(ctx context.Context, hostId uint64) {
	m.factory.Host().Delete(ctx, hostId)

}

func (m *host) Update(ctx context.Context, hostId uint64, host *asset.Host) {
	m.factory.Host().Update(ctx, hostId, host)

}

func (m *host) List(ctx context.Context, page, limit int) *form.PageHost {

	return m.factory.Host().List(page, limit)

}

func (m *host) GetByHostId(ctx context.Context, hostId uint64) (host *asset.Host) {
	return m.factory.Host().GetByHostId(ctx, hostId)

}

func (m *host) UpdateFiledStatus(ctx context.Context, hostId uint64, updateFiled string, status int8) {
	m.factory.Host().UpdateFiledStatus(ctx, hostId, updateFiled, status)

}

func (m *host) CheckHostExist(ctx context.Context, hostId uint64) bool {
	h := m.factory.Host().GetByHostId(ctx, hostId)
	if h == nil {
		return false
	}
	return true
}
