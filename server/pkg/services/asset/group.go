package asset

import (
	"context"
	"github.com/lbemi/lbemi/pkg/restfulx"

	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/form"
	"gorm.io/gorm"
)

type group struct {
	db *gorm.DB
}

func NewGroup(DB *gorm.DB) IGroup {
	return &group{db: DB}
}

type IGroup interface {
	Create(ctx context.Context, group *asset.Group)
	Delete(ctx context.Context, groupId uint64)
	Update(ctx context.Context, groupId uint64, group *asset.Group)
	List(page, limit int) *form.PageResult
	GetByGroupId(ctx context.Context, groupId uint64) (group *asset.Group)
	UpdateFiledStatus(ctx context.Context, groupId uint64, updateFiled string, status int8)
}

func (g *group) Create(ctx context.Context, group *asset.Group) {
	restfulx.ErrNotNilDebug(g.db.Create(group).Error, restfulx.OperatorErr)
}

func (g *group) Delete(ctx context.Context, groupId uint64) {
	restfulx.ErrNotNilDebug(g.db.Where("id = ?", groupId).Delete(&asset.Group{}).Error, restfulx.OperatorErr)
}

func (g *group) Update(ctx context.Context, groupId uint64, group *asset.Group) {
	restfulx.ErrNotNilDebug(g.db.Where("id = ?", groupId).Updates(group).Error, restfulx.OperatorErr)
}

func (g *group) List(page, limit int) *form.PageResult {
	var (
		groupList []asset.Group
		total     int64
	)

	// 全量查询
	if page == 0 && limit == 0 {
		restfulx.ErrNotNilDebug(g.db.Find(&groupList).Error, restfulx.OperatorErr)
		restfulx.ErrNotNilDebug(g.db.Model(&asset.Group{}).Count(&total).Error, restfulx.OperatorErr)

		res := &form.PageResult{
			Data:  groupList,
			Total: total,
		}
		return res
	}

	//分页数据
	restfulx.ErrNotNilDebug(g.db.Limit(limit).Offset((page-1)*limit).
		Find(&groupList).Error, restfulx.OperatorErr)

	restfulx.ErrNotNilDebug(g.db.Model(&asset.Group{}).Count(&total).Error, restfulx.OperatorErr)

	groups := GetTree(groupList, 0)

	res := &form.PageResult{
		Data:  groups,
		Total: total,
	}
	return res
}

func (g *group) GetByGroupId(ctx context.Context, groupId uint64) (group *asset.Group) {
	group = &asset.Group{}
	restfulx.ErrNotNilDebug(g.db.Where("id = ?", groupId).Find(&group).Error, restfulx.OperatorErr)
	return group
}

func (g *group) UpdateFiledStatus(ctx context.Context, groupId uint64, updateFiled string, status int8) {
	restfulx.ErrNotNilDebug(g.db.Where("id = ?", groupId).Update(updateFiled, status).Error, restfulx.OperatorErr)
}

func GetTree(groups []asset.Group, parentID uint64) (treeGroups []asset.Group) {
	for _, g := range groups {
		if g.ParentId == parentID {
			child := GetTree(groups, g.ID)
			g.Children = child
			treeGroups = append(treeGroups, g)
		}
	}
	return treeGroups
}
