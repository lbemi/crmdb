package asset

import (
	"context"

	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/form"
	"github.com/lbemi/lbemi/pkg/services"
)

type GroupGetter interface {
	Group() IGroup
}

type group struct {
	factory services.FactoryImp
}

func NewGroup(f services.FactoryImp) IGroup {
	return &group{
		factory: f,
	}
}

// IGroup 主机操作接口
type IGroup interface {
	Create(ctx context.Context, group *asset.Group)
	Delete(ctx context.Context, groupId uint64)
	Update(ctx context.Context, groupId uint64, group *asset.Group)
	List(ctx context.Context, page, limit int) *form.PageResult
	GetByGroupId(ctx context.Context, groupId uint64) (group *asset.Group)
	UpdateFiledStatus(ctx context.Context, groupId uint64, updateFiled string, status int8)
	CheckGroupExist(ctx context.Context, groupId uint64) bool
}

func (m *group) Create(ctx context.Context, group *asset.Group) {
	m.factory.Group().Create(ctx, group)
}

func (m *group) Delete(ctx context.Context, groupId uint64) {
	m.factory.Group().Delete(ctx, groupId)

}

func (m *group) Update(ctx context.Context, groupId uint64, group *asset.Group) {
	m.factory.Group().Update(ctx, groupId, group)

}

func (m *group) List(ctx context.Context, page, limit int) *form.PageResult {
	return m.factory.Group().List(page, limit)

}

func (m *group) GetByGroupId(ctx context.Context, groupId uint64) (group *asset.Group) {
	return m.factory.Group().GetByGroupId(ctx, groupId)

}

func (m *group) UpdateFiledStatus(ctx context.Context, groupId uint64, updateFiled string, status int8) {
	m.factory.Group().UpdateFiledStatus(ctx, groupId, updateFiled, status)

}

func (m *group) CheckGroupExist(ctx context.Context, groupId uint64) bool {
	h := m.factory.Group().GetByGroupId(ctx, groupId)
	if h == nil {
		return false
	}
	return true
}
