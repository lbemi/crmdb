package services

import (
	"context"
	"github.com/lbemi/lbemi/apps/asset/entity"
	entity2 "github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"gorm.io/gorm"
)

type GroupGetter interface {
	Group() IGroup
}

type group struct {
	db *gorm.DB
}

func NewGroup(db *gorm.DB) IGroup {
	return &group{
		db: db,
	}
}

// IGroup 主机操作接口
type IGroup interface {
	Create(ctx context.Context, group *entity.Group)
	Delete(ctx context.Context, groupId uint64)
	Update(ctx context.Context, groupId uint64, group *entity.Group)
	List(ctx context.Context, page, limit int) *entity2.PageResult
	GetByGroupId(ctx context.Context, groupId uint64) (group *entity.Group)
	UpdateFiledStatus(ctx context.Context, groupId uint64, updateFiled string, status int8)
	CheckGroupExist(ctx context.Context, groupId uint64) bool
}

func (g *group) Create(ctx context.Context, group *entity.Group) {
	restfulx.ErrNotNilDebug(g.db.Create(group).Error, restfulx.OperatorErr)
}

func (g *group) Delete(ctx context.Context, groupId uint64) {
	restfulx.ErrNotNilDebug(g.db.Where("id = ?", groupId).Delete(&entity.Group{}).Error, restfulx.OperatorErr)
}

func (g *group) Update(ctx context.Context, groupId uint64, group *entity.Group) {
	restfulx.ErrNotNilDebug(g.db.Where("id = ?", groupId).Updates(group).Error, restfulx.OperatorErr)
}

func (g *group) List(ctx context.Context, page, limit int) *entity2.PageResult {
	var (
		groupList []entity.Group
		total     int64
	)

	// 全量查询
	if page == 0 && limit == 0 {
		restfulx.ErrNotNilDebug(g.db.Find(&groupList).Error, restfulx.OperatorErr)
		restfulx.ErrNotNilDebug(g.db.Model(&entity.Group{}).Count(&total).Error, restfulx.OperatorErr)

		groups := GetTree(groupList, 0)
		res := &entity2.PageResult{
			Data:  groups,
			Total: total,
		}
		return res
	}

	//分页数据
	restfulx.ErrNotNilDebug(g.db.Limit(limit).Offset((page-1)*limit).
		Find(&groupList).Error, restfulx.OperatorErr)

	restfulx.ErrNotNilDebug(g.db.Model(&entity.Group{}).Count(&total).Error, restfulx.OperatorErr)

	groups := GetTree(groupList, 0)

	res := &entity2.PageResult{
		Data:  groups,
		Total: total,
	}
	return res
}

func (g *group) GetByGroupId(ctx context.Context, groupId uint64) (group *entity.Group) {
	group = &entity.Group{}
	restfulx.ErrNotNilDebug(g.db.Where("id = ?", groupId).Find(&group).Error, restfulx.OperatorErr)
	return group
}

func (g *group) UpdateFiledStatus(ctx context.Context, groupId uint64, updateFiled string, status int8) {
	restfulx.ErrNotNilDebug(g.db.Where("id = ?", groupId).Update(updateFiled, status).Error, restfulx.OperatorErr)
}

func (g *group) CheckGroupExist(ctx context.Context, groupId uint64) bool {
	group := &entity.Group{}
	err := g.db.Where("id = ?", groupId).Find(&group).Error
	if err == nil {
		return false
	}
	return true
}
func GetTree(groups []entity.Group, parentID uint64) (treeGroups []entity.Group) {
	for _, g := range groups {
		if g.ParentId == parentID {
			child := GetTree(groups, g.ID)
			g.Children = child
			treeGroups = append(treeGroups, g)
		}
	}
	return treeGroups

}
