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
	List(ctx context.Context, page, limit int, groups []uint64, ip, label, description string) *form.PageHost
	GetByHostId(ctx context.Context, hostId uint64) (host *asset.Host)
	GetHostAccounts(ctx context.Context, hostId uint64) []*asset.Account
	GetByGroup(ctx context.Context, groups []uint64, page, limit int) *form.PageResult
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

func (m *host) List(ctx context.Context, page, limit int, groups []uint64, ip, label, description string) *form.PageHost {
	return m.factory.Host().List(ctx, page, limit, groups, ip, label, description)

}

func (m *host) GetByHostId(ctx context.Context, hostId uint64) (host *asset.Host) {
	return m.factory.Host().GetByHostId(ctx, hostId)
}

func (m *host) GetByGroup(ctx context.Context, groups []uint64, page, limit int) *form.PageResult {
	return m.factory.Host().GetByGroup(ctx, groups, page, limit)
}

func (m *host) UpdateFiledStatus(ctx context.Context, hostId uint64, updateFiled string, status int8) {
	m.factory.Host().UpdateFiledStatus(ctx, hostId, updateFiled, status)

}

func (m *host) GetHostAccounts(ctx context.Context, hostId uint64) []*asset.Account {
	accountIds := make([]uint64, 0)
	list := m.factory.ResourceBindAccount().List(ctx, 0, 0)
	for _, ha := range list.Data.([]*asset.HostAccount) {
		for _, ra := range ha.ResourceId {
			if ra == hostId {
				accountIds = append(accountIds, ha.AccountId...)
			}
		}
	}
	accountIds = RemoveRepByMap(accountIds)
	return m.factory.Account().GetByIds(ctx, accountIds)
}

func (m *host) CheckHostExist(ctx context.Context, hostId uint64) bool {
	h := m.factory.Host().GetByHostId(ctx, hostId)
	if h == nil {
		return false
	}
	return true
}

func RemoveRepByMap(slc []uint64) []uint64 {
	var result []uint64
	tempMap := map[uint64]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}
