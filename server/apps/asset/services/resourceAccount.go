package services

import (
	"context"
	"github.com/lbemi/lbemi/apps/asset/entity"
	entity2 "github.com/lbemi/lbemi/pkg/common/entity"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"gorm.io/gorm"
)

type ResourceAccountGetter interface {
	ResourceBindAccount() IResourceBindAccount
}

type IResourceBindAccount interface {
	BindAccount(ctx context.Context, hostAccount *entity.HostAccount)
	UnbindAccount(ctx context.Context, haId uint64)
	UpdateHostAccount(ctx context.Context, hostAccount *entity.HostAccount)
	List(ctx context.Context, page, limit int) *entity2.PageResult
	Get(ctx context.Context, haId uint64) *entity.HostAccount
}

type resourceBindAccount struct {
	db *gorm.DB
}

func NewResourceBindAccount(db *gorm.DB) IResourceBindAccount {
	return &resourceBindAccount{
		db: db,
	}
}

func (r *resourceBindAccount) BindAccount(ctx context.Context, hostAccount *entity.HostAccount) {
	restfulx.ErrNotNilDebug(r.db.Create(&hostAccount).Error, restfulx.OperatorErr)
}

func (r *resourceBindAccount) UnbindAccount(ctx context.Context, haId uint64) {
	restfulx.ErrNotNilDebug(r.db.Where("id = ?", haId).Delete(&entity.HostAccount{}).Error, restfulx.OperatorErr)
}

func (r *resourceBindAccount) UpdateHostAccount(ctx context.Context, hostAccount *entity.HostAccount) {
	restfulx.ErrNotNilDebug(r.db.Where("id = ?", hostAccount.ID).Updates(hostAccount).Error, restfulx.OperatorErr)
}

func (r *resourceBindAccount) List(ctx context.Context, page, limit int) *entity2.PageResult {
	var (
		hostAccountList []*entity.HostAccount
		total           int64
	)

	// 全量查询
	if page == 0 && limit == 0 {
		restfulx.ErrNotNilDebug(r.db.Find(&hostAccountList).Error, restfulx.OperatorErr)
		restfulx.ErrNotNilDebug(r.db.Model(&entity.Account{}).Count(&total).Error, restfulx.OperatorErr)

		res := &entity2.PageResult{
			Data:  hostAccountList,
			Total: total,
		}
		return res
	}

	//分页数据
	restfulx.ErrNotNilDebug(r.db.Limit(limit).Offset((page-1)*limit).
		Find(&hostAccountList).Error, restfulx.OperatorErr)

	restfulx.ErrNotNilDebug(r.db.Model(&entity.HostAccount{}).Count(&total).Error, restfulx.OperatorErr)

	res := &entity2.PageResult{
		Data:  hostAccountList,
		Total: total,
	}
	return res
}

func (r *resourceBindAccount) Get(ctx context.Context, haId uint64) *entity.HostAccount {
	hostAccount := &entity.HostAccount{}
	restfulx.ErrNotNilDebug(r.db.Where("id = ?", haId).Find(&hostAccount).Error, restfulx.OperatorErr)
	return hostAccount
}
