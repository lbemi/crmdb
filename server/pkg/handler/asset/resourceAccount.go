package asset

import (
	"context"
	"github.com/lbemi/lbemi/pkg/restfulx"
	"gorm.io/gorm"

	"github.com/lbemi/lbemi/pkg/model/asset"
	"github.com/lbemi/lbemi/pkg/model/form"
)

type ResourceAccountGetter interface {
	ResourceBindAccount() IResourceBindAccount
}

type IResourceBindAccount interface {
	BindAccount(ctx context.Context, hostAccount *asset.HostAccount)
	UnbindAccount(ctx context.Context, haId uint64)
	UpdateHostAccount(ctx context.Context, hostAccount *asset.HostAccount)
	List(ctx context.Context, page, limit int) *form.PageResult
	Get(ctx context.Context, haId uint64) *asset.HostAccount
}

type resourceBindAccount struct {
	db *gorm.DB
}

func NewResourceBindAccount(db *gorm.DB) IResourceBindAccount {
	return &resourceBindAccount{
		db: db,
	}
}

func (r *resourceBindAccount) BindAccount(ctx context.Context, hostAccount *asset.HostAccount) {
	restfulx.ErrNotNilDebug(r.db.Create(&hostAccount).Error, restfulx.OperatorErr)
}

func (r *resourceBindAccount) UnbindAccount(ctx context.Context, haId uint64) {
	restfulx.ErrNotNilDebug(r.db.Where("id = ?", haId).Delete(&asset.HostAccount{}).Error, restfulx.OperatorErr)
}

func (r *resourceBindAccount) UpdateHostAccount(ctx context.Context, hostAccount *asset.HostAccount) {
	restfulx.ErrNotNilDebug(r.db.Where("id = ?", hostAccount.ID).Updates(hostAccount).Error, restfulx.OperatorErr)
}

func (r *resourceBindAccount) List(ctx context.Context, page, limit int) *form.PageResult {
	var (
		hostAccountList []*asset.HostAccount
		total           int64
	)

	// 全量查询
	if page == 0 && limit == 0 {
		restfulx.ErrNotNilDebug(r.db.Find(&hostAccountList).Error, restfulx.OperatorErr)
		restfulx.ErrNotNilDebug(r.db.Model(&asset.Account{}).Count(&total).Error, restfulx.OperatorErr)

		res := &form.PageResult{
			Data:  hostAccountList,
			Total: total,
		}
		return res
	}

	//分页数据
	restfulx.ErrNotNilDebug(r.db.Limit(limit).Offset((page-1)*limit).
		Find(&hostAccountList).Error, restfulx.OperatorErr)

	restfulx.ErrNotNilDebug(r.db.Model(&asset.HostAccount{}).Count(&total).Error, restfulx.OperatorErr)

	res := &form.PageResult{
		Data:  hostAccountList,
		Total: total,
	}
	return res
}

func (r *resourceBindAccount) Get(ctx context.Context, haId uint64) *asset.HostAccount {
	hostAccount := &asset.HostAccount{}
	restfulx.ErrNotNilDebug(r.db.Where("id = ?", haId).Find(&hostAccount).Error, restfulx.OperatorErr)
	return hostAccount
}
